package main

import (
	"log"
	"strings"
	"time"
)

func TryAction(s Slack) Labs {

	if s.Text == "" {
		return labs

	}

	stringSlice := strings.Split(s.Text, ",")
	log.Print("strslice", stringSlice)
	if strings.Contains(strings.ToLower(stringSlice[0]), "lab") {
		labName := stringSlice[0]
		return CreateLab(s, labName)
	}

	return nil
}

//CreateLab creates a new lab
func CreateLab(s Slack, labName string) Labs {

	RepoCreateLab(Lab{labName, "in", s.User, time.Now()})

	return labs
}
