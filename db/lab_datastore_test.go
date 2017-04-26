package db

import (
	"context"
	"testing"
	"time"

	"github.com/ECSTeam/labcheck/model"

	"cloud.google.com/go/datastore"
)

var env = "ecs-pcf-on-gce-2016"

func testDB(t *testing.T, db model.LabDatabase) {
	defer db.Close()

	l := &model.Lab{
		Name: "lab01test4", Available: true, Desc: "This is a test", User: "", LastUpdate: time.Now(),
	}

	if err := db.AddLab(l); err != nil {
		t.Error(err)
	}

	l.Desc = "Test desc changed"
	if err := db.UpdateLab(l); err != nil {
		t.Error(err)
	}

	gotLab, err := db.GetLabByName(l.Name)
	if err != nil {
		t.Error(err)
	}

	if got, want := gotLab.Desc, l.Desc; got != want {
		t.Errorf("Update description: got %q, want %q", got, want)
	}

	if err := db.DeleteLab(l.Name); err != nil {
		t.Error(err)
	}

	if _, err := db.GetLabByName(l.Name); err == nil {
		//not sure why this keeps (sometimes) returning a value after delete...doesn't show in db
		t.Errorf("want non-nil err")
	}

}

func TestDatastoreDB(t *testing.T) {
	//tc := testutil.SystemTest(t)
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, env)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	db, err := newDatastoreDB(client)
	if err != nil {
		t.Fatal(err)
	}
	testDB(t, db)
}
