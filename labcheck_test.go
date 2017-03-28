package main

import "testing"
import "time"

func TestCreateLab(t *testing.T) {

	t.Log("Creating Lab")
	lab := NewLab(3, "testlab")

	t.Error("I'm in a  bad mood.", lab.Name)
}

func NewLab(id int, name string) *Lab {
	if id < 0 {
		return nil
	}
	l := Lab{id, name, "foo", "bar", time.Now()}
	return &l
}
