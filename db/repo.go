package db

import "github.com/kamattson/labcheck/model"

var labs model.Lab

// Give us some seed data
/*
func init() {
	log.Print("init labs...")
	RepoCreateLab(Lab{Name: "lab01", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab02", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab03", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab04", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab05", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab06", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab07", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab08", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab09", Available: true, Desc: "", User: "", LastUpdate: time.Now()})
	RepoCreateLab(Lab{Name: "lab10", Available: true, Desc: "", User: "", LastUpdate: time.Now()})

	//InitDB(env)
}*/
/*
func LoadLabs(labs model.Labs) error {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, env)
	for _, l := range labs {
		log.Printf("Loading %+v", l)
		key, err := AddLab(ctx, client, l)
		if err != nil {
			log.Printf("ERROR %+v", err)
			return err
		}
		log.Printf("Key %+v", key)

	}
	return nil
}
*/
/*
//ListLabs lists all labs
func ListLabs() ([]Lab, error) {
	var labs []Lab
	var err error
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, env)

	query := datastore.NewQuery("Lab").Order("Name")

	fmt.Printf("query %+v", query)
	fmt.Printf("\n")

	keys, err := client.GetAll(ctx, query, &labs)
	fmt.Printf("keys %+v", keys)
	fmt.Printf("\n")
	if err != nil {
		return nil, err
	}

	// Set the id field on each Lab from the corresponding key.
	for i, key := range keys {
		labs[i].ID = key.String()
	}

	return labs, err
}

// AddLab adds a lab to the datastore,
// returning the key of the newly created entity.
func AddLab(ctx context.Context, client *datastore.Client, lab model.Lab) (*datastore.Key, error) {
	key := datastore.IncompleteKey("Lab", nil)
	return client.Put(ctx, key, &lab)
}

//RepoFindLab find lab in repo
func RepoFindLab(labName string) (model.Lab, error) {
	for _, t := range labs {
		if strings.Compare(t.Name, labName) == 0 {
			return t, nil
		}
	}
	// return empty lab if not found
	return Lab{}, errors.New("No Lab Found")
}

//RepoCreateLab create a lab and append to slice
func RepoCreateLab(t model.Lab) Lab {
	labs = append(labs, t)
	return t
}
func DeleteLabs(labs model.Labs) error {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, env)
	for _, l := range labs {
		err := DeleteLab(ctx, client, l.Name)
		if err != nil {
			return err
		}
	}
	return nil

}

// [START delete_entity]
// DeleteLab deletes the lab with the given ID.
func DeleteLab(ctx context.Context, client *datastore.Client, labName string) error {
	log.Printf("deleting..%v", labName)
	return client.Delete(ctx, datastore.NameKey("Lab", labName, nil))
}
*/
