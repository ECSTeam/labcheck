package model

import "time"

type Lab struct {
	Name       string    `json:"name"`
	Available  bool      `json:"available"`
	Version    string    `json:"version"`
	Desc       string    `json:"desc"`
	User       string    `json:"user"`
	LastUpdate time.Time `json:"lastUpdated"`
}

// LabDatabase provides thread-safe access to a database of Labs.
type LabDatabase interface {

	// ListLabs returns a list of Labs, ordered by name.
	ListLabs() ([]*Lab, error)

	// GetLabByName returns a lab by name
	GetLabByName(name string) (*Lab, error)

	// LoadLabs saves a given Lab, assigning it a new name.
	LoadLabs(l []Lab) (err error)

	// AddLab saves a given Lab, assigning it a new name.
	AddLab(l *Lab) (id int64, err error)

	// DeleteLab removes a given Lab by its name.
	DeleteLab(name string) error

	// UpdateLab updates the entry for a given Lab.
	UpdateLab(l *Lab) error

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}
