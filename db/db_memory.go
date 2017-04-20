package db

import (
	"errors"
	"sync"

	"github.com/kamattson/labcheck/model"
)

// Ensure memoryDB conforms to the LabDatabase interface.
var _ model.LabDatabase = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for labs.
type memoryDB struct {
	mu     sync.Mutex
	nextID int64                // next ID to assign to a lab.
	labs   map[int64]*model.Lab // maps from Lab ID to Lab.
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		labs:   make(map[int64]*model.Lab),
		nextID: 1,
	}
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.labs = nil
}

func (db *memoryDB) GetLabByName(name string) (*model.Lab, error) {
	// not implemented
	return nil, nil
}

// LoadLabs loads a bunch of labs
func (db *memoryDB) LoadLabs(l []model.Lab) (err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, lb := range l {
		db.labs[db.nextID] = &lb
		db.nextID++
	}
	return nil
}

// AddLab saves a given lab, assigning it a new ID.
func (db *memoryDB) AddLab(l *model.Lab) (err error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.labs[db.nextID] = l
	db.nextID++

	return nil
}

// DeleteLab removes a given lab by its ID.
func (db *memoryDB) DeleteLab(name string) error {
	if name == "" {
		return errors.New("memorydb: lab with unassigned ID passed into deleteLab")
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	//TODO: fix me
	//if _, ok := db.labs[id]; !ok {
	//	return fmt.Errorf("memorydb: could not delete lab with ID %d, does not exist", id)
	//}
	//delete(db.labs, name)
	return nil
}

// UpdateLab updates the entry for a given lab.
func (db *memoryDB) UpdateLab(l *model.Lab) error {
	if l.Name == "" {
		return errors.New("memorydb: lab with unassigned ID passed into updateLab")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, lab := range db.labs {
		if lab.Name == l.Name {
			lab = l
		}

	}

	return nil
}

// ListLabs returns a list of labs
func (db *memoryDB) ListLabs() ([]*model.Lab, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var labs []*model.Lab
	for _, l := range db.labs {
		labs = append(labs, l)
	}
	return labs, nil
}
