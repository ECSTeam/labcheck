package main

import (
	"fmt"
	"log"
	"time"
)

var labs Labs

// Give us some seed data
func init() {
	RepoCreateLab(Lab{Name: "Lab01", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab02", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab03", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab04", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab05", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab06", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab07", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab08", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab09", Status: "Available", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "Lab10", Status: "Available", User: "", LastUpdate: time.Now()})
}

//RepoFindLab find lab in repo
func RepoFindLab(labName string) Lab {
	for _, t := range labs {
		if t.Name == labName {
			return t
		}
	}
	// return empty lab if not found
	return Lab{}
}

//RepoCreateLab create a lab and append to slice
func RepoCreateLab(t Lab) Lab {
	//TODO check for duplicates
	log.Print("creating a lab...", t.Name)
	labs = append(labs, t)
	return t
}

func RepoDestroyLab(name string) error {
	for i, t := range labs {
		if t.Name == name {
			labs = append(labs[:i], labs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Lab with name of %v to delete", name)
}
