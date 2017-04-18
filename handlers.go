package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kamattson/labcheck/db"
	"github.com/kamattson/labcheck/model"
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

	for key, values := range r.PostForm {
		log.Printf("key=%v, value=%v", key, values)
	}
	defer r.Body.Close()
	var slack = initSlack(r)

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
		}
		log.Printf("lab %v status %v: ", lab.Name, lab.Available)
		db.DB.UpdateLab(lab)

		Render.JSON(w, http.StatusOK, lab)

	}

	/*
		if strings.Contains(strings.ToLower(stringSlice[0]), "lab") {
			var labName = strings.ToLower(stringSlice[0])
			lab, err := RepoFindLab(labName)
			if err != nil {
				RepoCreateLab(Lab{labName, true, "", "", s.User, time.Now()})
				return labs
			} else {
				log.Printf("Lab %v already exists %v", stringSlice[0], lab)
			}

			return labs
		}
	*/

}

func initSlack(r *http.Request) model.Slack {
	var slack model.Slack
	slack.Command = r.FormValue("command")
	slack.User = r.FormValue("user_name")
	slack.Text = r.FormValue("text")
	slack.ResponseURL = r.FormValue("response_url")
	slack.TeamDomain = r.FormValue("team_domain")
	slack.ChannelID = r.FormValue("channel_id")
	slack.ChannelName = r.FormValue("channel_name")
	slack.UserID = r.FormValue("user_id")
	slack.Token = r.FormValue("token")
	slack.TeamID = r.FormValue("team_id")

	return slack

}
