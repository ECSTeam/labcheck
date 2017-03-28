package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index convenience
func Index(w http.ResponseWriter, r *http.Request) {
	Render.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}

//LabIndex handler for /labs
func LabIndex(w http.ResponseWriter, r *http.Request) {
	if err := Render.JSON(w, http.StatusOK, labs); err != nil {
		panic(err)
	}
}

//LabShow handler for /labs/{labName}
func LabShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	labName := vars["labName"]
	lab := RepoFindLab(labName)
	Render.JSON(w, http.StatusOK, lab)
	//fmt.Fprintln(w, "Lab show:", labName)
}

//LabCreate handler for POST /labs/
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
