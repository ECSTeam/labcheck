package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ECSTeam/labcheck/db"
	"github.com/ECSTeam/labcheck/helpers"
	"github.com/ECSTeam/labcheck/model"
)

func index(w http.ResponseWriter, r *http.Request) {
	Render.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}

//LabsHandler handler for GET on /labs
func labsHandler(w http.ResponseWriter, r *http.Request) {
	l, err := db.DB.ListLabs()

	if err != nil {
		log.Printf("could not list labs: %v", err)
	}
	Render.JSON(w, http.StatusOK, l)
}

//LoadData handler for POST /load
func loadData(w http.ResponseWriter, r *http.Request) {
	var labs []model.Lab
	defer r.Body.Close()
	payload, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(payload, &labs); err != nil {
		Render.JSON(w, http.StatusUnprocessableEntity, err)
	}

	if err := db.DB.LoadLabs(labs); err != nil {
		log.Printf("error loading labs...%v", err)
	}
	if err := Render.JSON(w, http.StatusCreated, map[string]string{"Success": "Labs created successfully!"}); err != nil {
		panic(err)
	}
}

//LabCheck handler for POST /labs
func labCheck(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		Render.JSON(w, http.StatusUnprocessableEntity, err)
	}
	defer r.Body.Close()
	if os.Getenv("LOCAL") == "false" {
		if helpers.CheckToken(r.FormValue("token")) == false {
			//put call to token helper here
			Render.JSON(w, http.StatusUnauthorized, "invalid token")
		}

	}

	var slack = initSlackRequest(r)
	var slackResponse SlackResponse
	var comment = ""
	var rgx = regexp.MustCompile(`\{(.*?)\}`)

	//empty texts returns all labs
	if len(strings.TrimSpace(slack.Text)) == 0 {
		l, err := db.DB.ListLabs()
		if err != nil {
			log.Printf("could not list labs: %v", err)
		}
		slackResponse.ResponseType = "ephemeral"
		slackResponse.Text = "Labs"
		slackResponse.Attachments = make([]Attachments, 0)

		for _, lab := range l {

			attachment := Attachments{
				Pretext:   lab.Name,
				Color:     color(lab.Available),
				Text:      "version: " + lab.Version + " Description:" + lab.Desc,
				Footer:    lab.User,
				Timestamp: lab.LastUpdate.Unix(),
			}
			slackResponse.Attachments = append(slackResponse.Attachments, attachment)
		}

		Render.JSON(w, http.StatusOK, slackResponse)
		return
	}

	//labs help command
	if slack.Text == "help" {
		slackResponse.ResponseType = "ephemeral"
		slackResponse.Text = "Labcheck Help"
		slackResponse.Attachments = make([]Attachments, 0)
		attachment := Attachments{
			Title:     "Labcheck README.md page",
			TitleLink: "https://github.com/ECSTeam/labcheck",
			Text: `Commands:
/labs
/labs checkout labxx {"_optional comment_"}
/labs checkin labxx
/labs status labxx
/labs update labxx {"version":"x.x", "desc":"..."}
/labs help`,
		}
		slackResponse.Attachments = append(slackResponse.Attachments, attachment)
		Render.JSON(w, http.StatusOK, slackResponse)
		return

	}

	stringSlice := strings.Split(slack.Text, " ")

	//if text starts with check - do some checkin/checkout logic
	if strings.Contains(strings.ToLower(stringSlice[0]), "check") {
		checkStatus := stringSlice[0]
		//get lab by name which should be the 2nd string
		lab, err := db.DB.GetLabByName(stringSlice[1])
		//entity not found, return a 204
		if err != nil {
			Render.JSON(w, http.StatusNoContent, err)
			return
		}
		lab.LastUpdate = time.Now()
		if strings.Compare(checkStatus, "checkin") == 0 {
			lab.Available = true
			lab.User = ""
		}
		if strings.Compare(checkStatus, "checkout") == 0 {
			lab.Available = false
			lab.User = slack.User

			var cmt = rgx.FindStringSubmatch(slack.Text)
			if len(cmt) > 0 {
				comment = cmt[1]
			}
		}
		db.DB.UpdateLab(lab)

		slackResponse.ResponseType = "in_channel"
		slackResponse.Text = lab.Name

		slackResponse.Attachments = make([]Attachments, 0)
		attachment := Attachments{
			Color:      color(lab.Available),
			Title:      checkStatus,
			AuthorName: slack.User,
			Text:       comment,
		}
		slackResponse.Attachments = append(slackResponse.Attachments, attachment)
		Render.JSON(w, http.StatusOK, slackResponse)
		return
	}

	//labs status <labname>
	if strings.Contains(strings.ToLower(stringSlice[0]), "status") {
		var labName = strings.ToLower(stringSlice[1])
		lab, err := db.DB.GetLabByName(labName)
		if err != nil {
			Render.JSON(w, http.StatusNotFound, labName)
			return
		}

		slackResponse.Text = lab.Name

		slackResponse.Attachments = make([]Attachments, 0)
		attachment := Attachments{
			Color:      color(lab.Available),
			AuthorName: lab.User,
			Text:       "version: " + lab.Version + " | Description: " + lab.Desc,
			Timestamp:  lab.LastUpdate.Unix(),
		}
		slackResponse.Attachments = append(slackResponse.Attachments, attachment)

		Render.JSON(w, http.StatusOK, slackResponse)
		return
	}

	//labs update <labname>
	if strings.Contains(strings.ToLower(stringSlice[0]), "update") {
		var labName = strings.ToLower(stringSlice[1])
		var regexString = rgx.FindStringSubmatch(slack.Text)
		var l model.Lab
		b := []byte("{" + regexString[1] + "}")
		err := json.Unmarshal(b, &l)
		if err != nil {
			Render.JSON(w, http.StatusNotModified, err)
			return
		}

		lab, err := db.DB.GetLabByName(labName)
		if err != nil {
			Render.JSON(w, http.StatusNotFound, labName)
			return
		}
		lab.LastUpdate = time.Now()
		lab.Desc = l.Desc
		lab.Version = l.Version
		db.DB.UpdateLab(lab)

		slackResponse.Text = lab.Name

		slackResponse.Attachments = make([]Attachments, 0)
		attachment := Attachments{
			Color:      color(lab.Available),
			AuthorName: lab.User,
			Text:       "version: " + lab.Version + " | Description: " + lab.Desc,
			Timestamp:  lab.LastUpdate.Unix(),
		}
		slackResponse.Attachments = append(slackResponse.Attachments, attachment)

		if slack.Text != slack.OriginalText {
			attachment := Attachments{
				Text: "Original text contained special, incompatible, characters. \n*Original text*: `" + slack.OriginalText + "`\n*Cleansed text*: `" + slack.Text + "`",
			}
			slackResponse.Attachments = append(slackResponse.Attachments, attachment)
		}

		Render.JSON(w, http.StatusOK, slackResponse)

		//HTML doesn't render in slack...but if it did!!!
		//Render.HTML(w, http.StatusOK, "labs", lab)
	}

}

func color(c bool) string {
	if c {
		return "#000000"
	}
	return "#ff0000"

}

func initSlackRequest(r *http.Request) model.Slack {
	var slack model.Slack
	slack.Command = r.FormValue("command")
	slack.User = r.FormValue("user_name")
	slack.OriginalText = r.FormValue("text")
	slack.Text = cleanupSmartQuotes(slack.OriginalText)
	slack.ResponseURL = r.FormValue("response_url")
	slack.TeamDomain = r.FormValue("team_domain")
	slack.ChannelID = r.FormValue("channel_id")
	slack.ChannelName = r.FormValue("channel_name")
	slack.UserID = r.FormValue("user_id")
	slack.Token = r.FormValue("token")
	slack.TeamID = r.FormValue("team_id")

	return slack

}

func cleanupSmartQuotes(myText string) string {
	var returnText string
	returnText = strings.Replace(myText, "\uFFFD", `\\'`, -1)
	returnText = strings.Replace(myText, "\u201A", `\\'`, -1)
	returnText = strings.Replace(myText, "\u2018", `\\'`, -1)
	returnText = strings.Replace(myText, "\u2019", `\\'`, -1)
	returnText = strings.Replace(myText, "\u201c", `\\"`, -1)
	returnText = strings.Replace(myText, "\u201d", `\\"`, -1)
	returnText = strings.Replace(myText, "\u201e", `\\"`, -1)
	returnText = strings.Replace(myText, "\u02C6", `^`, -1)
	returnText = strings.Replace(myText, "\u2039", `<`, -1)
	returnText = strings.Replace(myText, "\u203A", `>`, -1)
	returnText = strings.Replace(myText, "\u2013", `-`, -1)
	returnText = strings.Replace(myText, "\u2014", `--`, -1)
	returnText = strings.Replace(myText, "\u2026", `...`, -1)
	returnText = strings.Replace(myText, "\u00A9", `(c)`, -1)
	returnText = strings.Replace(myText, "\u00AE", `(r)`, -1)
	returnText = strings.Replace(myText, "\u2122", `TM`, -1)
	returnText = strings.Replace(myText, "\u00BC", `1/4`, -1)
	returnText = strings.Replace(myText, "\u00BD", `1/2`, -1)
	returnText = strings.Replace(myText, "\u00BE", `3/4`, -1)
	returnText = strings.Replace(myText, "\u02DC", ` `, -1)
	returnText = strings.Replace(myText, "\u00A0", ` `, -1)
	return returnText
}

//SlackResponse represents a slack response
type SlackResponse struct {
	ResponseType string        `json:"response_type"`
	Text         string        `json:"text"`
	Attachments  []Attachments `json:"attachments"`
}

type Attachments struct {
	Fallback   string `json:"fallback"`
	Color      string `json:"color"`
	Pretext    string `json:"pretext"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
	ImageURL   string `json:"image_url"`
	ThumbURL   string `json:"thumb_url"`
	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	Timestamp  int64  `json:"ts"`
}

type Fields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short string `json:"short"`
}
