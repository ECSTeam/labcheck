package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Index convenience
func Index(w http.ResponseWriter, r *http.Request) {
	Render.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}

//LabIndex handler for /labindex
func LabIndex(w http.ResponseWriter, r *http.Request) {
	if err := Render.JSON(w, http.StatusOK, labs); err != nil {
		panic(err)
	}
}

//LabShow handler for /labs/{labName}
func LabShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	labName := vars["labName"]
	lab, err := RepoFindLab(labName)
	if err != nil {
		Render.JSON(w, http.StatusUnprocessableEntity, map[string]string{labName: "Not Found!"})
	}
	Render.JSON(w, http.StatusOK, lab)
}

//LabCreate handler for POST /labs
func LabCreate(w http.ResponseWriter, r *http.Request) {
	var lab Lab
	defer r.Body.Close()
	payload, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err := json.Unmarshal(payload, &lab); err != nil {
		Render.JSON(w, http.StatusUnprocessableEntity, err)
	}

	t := RepoCreateLab(lab)
	if err := Render.JSON(w, http.StatusCreated, t); err != nil {
		panic(err)
	}
}

//LabCheck handler for POST /labs
func LabCheck(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Render.JSON(w, http.StatusUnprocessableEntity, err)
	}

	for key, values := range r.PostForm {
		log.Printf("key=%v, value=%v", key, values)
	}
	defer r.Body.Close()
	var slack = initSlack(r)
	Render.JSON(w, http.StatusOK, ProcessRequest(slack))

}

func ProcessRequest(s Slack) Labs {

	if s.Text == "" {
		return labs
	}

	stringSlice := strings.Split(s.Text, " ")

	if strings.Contains(strings.ToLower(stringSlice[0]), "check") {
		checkStatus := stringSlice[0]
		labName := stringSlice[1]
		lab, err := RepoFindLab(labName)
		if err != nil {
			panic(err)
		}
		lab.LastUpdate = time.Now()
		if strings.Compare(checkStatus, "checkin") == 0 {
			lab.Available = true
			lab.User = ""
		}
		log.Print("comp", strings.Compare(checkStatus, "checkout"))
		if strings.Compare(checkStatus, "checkout") == 0 {
			lab.Available = false
			lab.User = s.User
		}
		log.Printf("lab %v status %v: ", lab.Name, lab.Available)
		return labs
	}

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

	return nil
}

func initSlack(r *http.Request) Slack {
	var slack Slack
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
