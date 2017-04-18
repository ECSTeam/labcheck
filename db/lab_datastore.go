package db

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/api/iterator"

	"github.com/kamattson/labcheck/model"

	"cloud.google.com/go/datastore"
)

// datastoreDB persists labs to Cloud Datastore.
type datastoreDB struct {
	client *datastore.Client
}

func (db *datastoreDB) datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey("Lab", id, nil)
}

var _ model.LabDatabase = &datastoreDB{}

func newDatastoreDB(client *datastore.Client) (model.LabDatabase, error) {
	ctx := context.Background()
	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return nil, fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	return &datastoreDB{
		client: client,
	}, nil
}

func (db *datastoreDB) LoadLabs(l []model.Lab) (err error) {
	ctx := context.Background()
	var keys []*datastore.Key

	for _, lab := range l {
		keys = append(keys, datastore.NameKey("Lab", lab.Name, nil))
		//keys = append(keys, datastore.IncompleteKey("Lab", nil))
	}
	_, err = db.client.PutMulti(ctx, keys, l)

	if err != nil {
		return err
	}
	return nil
}

func (db *datastoreDB) AddLab(l *model.Lab) (id int64, err error) {
	ctx := context.Background()
	k := datastore.NameKey("Lab", l.Name, nil)
	k, err = db.client.Put(ctx, k, l)
	if err != nil {
		return 0, fmt.Errorf("datastoredb: could not put Lab: %v", err)
	}
	log.Printf("k%v", k)
	return 0, nil
}

func (db *datastoreDB) DeleteLab(name string) error {
	ctx := context.Background()
	k := datastore.NameKey("Lab", name, nil)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("datastoredb: could not delete Lab: %v", err)
	}
	return nil
}

// UpdateLab updates the entry for a given Lab.
func (db *datastoreDB) UpdateLab(l *model.Lab) error {
	ctx := context.Background()
	labKey := datastore.NameKey("Lab", l.Name, nil)
	log.Printf("labkey%+v", labKey)

	if _, err := db.client.Put(ctx, labKey, l); err != nil {
		log.Fatalf("tx.Put: %v", err)
	}

	/*
		k := db.datastoreKey()
		if _, err := db.client.Put(ctx, k, l); err != nil {
			return fmt.Errorf("datastoredb: could not update Lab: %v", err)
		}
	*/
	return nil
}

// ListLabs returns a list of Labs, ordered by title.
func (db *datastoreDB) ListLabs() ([]*model.Lab, error) {
	ctx := context.Background()
	labs := make([]*model.Lab, 0)
	q := datastore.NewQuery("Lab").
		Order("Name")

	_, err := db.client.GetAll(ctx, q, &labs)

	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not list labs: %v", err)
	}

	return labs, nil
}

// GetLabByName returns a lab by name
func (db *datastoreDB) GetLabByName(name string) (*model.Lab, error) {
	ctx := context.Background()

	log.Printf("Name = %v", name)
	q := datastore.NewQuery("Lab").
		Filter("Name =", name)

	log.Printf("query%+v", q)

	it := db.client.Run(ctx, q)
	for {
		var l model.Lab
		_, err := it.Next(&l)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Error fetching next lab: %v", err)
		}

		fmt.Printf("Lab %q,  %q\n", l.Desc, l.Name)
		return &l, nil
	}
	return nil, errors.New("no lab found")
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.
}
