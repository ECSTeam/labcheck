package main

import (
	"fmt"
	"io"
	"os"
	"testing"
	"text/tabwriter"
)

func TestLoadLab(t *testing.T) {
	t.Logf("TestLoadLab")

	LoadLabs()
	labs, _ := ListLabs()

	t.Logf("labs %+v", labs)

	DeleteLabs()
	labs2, _ := ListLabs()
	t.Logf("labs after delete %+v", labs2)
}

func TestListLabs(t *testing.T) {
	t.Logf("TestListLabs")
	labs, _ := ListLabs()
	PrintLabs(os.Stdout, labs)
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
