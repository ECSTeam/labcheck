package db

import (
	"context"
	"testing"
	"time"

	"github.com/kamattson/labcheck/model"

	"cloud.google.com/go/datastore"
)

var env = "ecs-pcf-on-gce-2016"

func testDB(t *testing.T, db model.LabDatabase) {
	defer db.Close()

	l := &model.Lab{
		Name: "lab01test4", Available: true, Desc: "This is a test", User: "", LastUpdate: time.Now(),
	}
	_, err := db.AddLab(l)

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

/*
var testLabs Labs

func XTestLoadLab(t *testing.T) {
	initTestLabs()
	fmt.Printf("testLabs %+v", testLabs)
	fmt.Printf("\n")
	LoadLabs(testLabs)

	labs, _ := ListLabs()
	fmt.Printf("labs %+v", labs)
	fmt.Printf("\n")
	err := DeleteLabs(testLabs)
	if err != nil {
		fmt.Print("labs error:", err)
	}

	labs2, _ := ListLabs()
	fmt.Printf("labs after delete %+v", labs2)
	fmt.Printf("\n")
}

func TestListLabs(t *testing.T) {
	labs, _ := ListLabs()
	fmt.Printf("labs %+v", labs)
	//PrintLabs(os.Stdout, labs)
}

func PrintLabs(w io.Writer, labs []*Lab) {
	// Use a tab writer to help make results pretty.
	tw := tabwriter.NewWriter(w, 8, 8, 1, ' ', 0) // Min cell size of 8.
	fmt.Fprintf(tw, "Name\tAvailable\tLast Updated\n")
	for _, l := range labs {
		fmt.Fprintf(tw, "%v\t%t\tcreated %v\n", l.Name, l.Available, l.LastUpdate)
	}
	tw.Flush()
}

func initTestLabs() {

	testLabs = []Lab{
		Lab{Name: "lab01test1", Available: true, Desc: "", User: "", LastUpdate: time.Now()},
		Lab{Name: "lab02test1", Available: true, Desc: "", User: "", LastUpdate: time.Now()},
		Lab{Name: "lab03test1", Available: true, Desc: "", User: "", LastUpdate: time.Now()},
		Lab{Name: "lab04test1", Available: true, Desc: "", User: "", LastUpdate: time.Now()},
	}
}
*/
