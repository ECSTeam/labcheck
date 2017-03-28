package main

import "fmt"

var currentId int

var labs Labs

// Give us some seed data
func init() {
    RepoCreateLab(Lab{Name: "Lab9", Status: "Checked Out"})
    RepoCreateLab(Lab{Name: "Lab10", Status: "Checked In"})
}

func RepoFindLab(id int) Lab {
    for _, t := range labs {
        if t.Id == id {
            return t
        }
    }
    // return empty lab if not found
    return Lab{}
}

func RepoCreateLab(t Lab) Lab {
    currentId += 1
    t.Id = currentId
    labs = append(labs, t)
    return t
}

func RepoDestroyLab(id int) error {
    for i, t := range labs {
        if t.Id == id {
            labs = append(labs[:i], labs[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Lab with id of %d to delete", id)
}
