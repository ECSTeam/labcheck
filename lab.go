package main

import "time"

type Lab struct {

    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Status    string    `json:"status"`
    User      string    `json:"user"`
    LastUpdate  time.Time `json:"lastUpdated"`
}

type Labs []Lab
