package db

import (
	"context"
	"log"

	"github.com/ECSTeam/labcheck/model"

	"cloud.google.com/go/datastore"
)

var (
	DB model.LabDatabase
)

func init() {
	var err error

	//DB = newMemoryDB()

	// [START datastore]
	// To use Cloud Datastore, uncomment the following lines and update the
	// project ID.
	// More options can be set, see the google package docs for details:
	// http://godoc.org/golang.org/x/oauth2/google
	//
	DB, err = configureDatastoreDB("ecs-pcf-on-gce-2016")
	// [END datastore]

	if err != nil {
		log.Fatal(err)
	}
}

func configureDatastoreDB(projectID string) (model.LabDatabase, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return newDatastoreDB(client)
}
