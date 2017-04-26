package main

import (
	"net/http"
	"net/http/httptest"

	"github.com/ECSTeam/labcheck/model"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

var (
	request   *http.Request
	recorder  *httptest.ResponseRecorder
	lab       model.Lab
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

/*
func TestLabsHandler(t *testing.T) {
	server := MakeTestServer()
	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/labs", nil)
	server.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &lab)
	if err != nil {
		t.Errorf("Error unmarshaling  response to Lab: %v", err)
	}
	fmt.Printf("/labs/%+v", lab)

}
*/
/*
func TestCreateLab(t *testing.T) {


	server := MakeTestServer()

	labName := "lab10"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/lab/"+labName, nil)
	server.ServeHTTP(recorder, request)

	var lab Lab

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &lab)
	if err != nil {
		t.Errorf("Error unmarshaling  response to Lab: %v", err)
	}

	if lab.Name != labName {
		t.Errorf("Expected Lab10, got %v", lab.Name)
	}

}
func TestCheckout(t *testing.T) {
	log.Print("/// TestCheckout ///")
	var s Slack
	var labName = "lab01"

	s.Text = "checkout " + labName
	TryAction(s)
	labx, _ := RepoFindLab(labName)
	log.Print(labName, " status:", labx.Status)
	if strings.Compare(labx.Status, "InUse") != 0 {
		t.Errorf("Error Returning labs  //TestCheckout//")
	}
}

func TestNoTextAction(t *testing.T) {
	log.Print("/// TestNoTextAction ///")
	var s Slack
	s.Text = ""

	TryAction(s)

	log.Print("Labs :", labs)
	if labs == nil {
		t.Errorf("Error Returning labs //TestNoTextAction//")
	}
}

func TestAddLabAction(t *testing.T) {
	log.Print("/// TestAddLabAction ///")

	//todo fix this
	var labname = "lab10"
	var s Slack
	s.Text = labname

	count1 := len(labs)
	TryAction(s)
	count2 := len(labs)

	log.Print("counts:", count1, count2)
	if count1 >= count2 {
		t.Errorf("Error creating Lab //TestAddLabAction//")
	}
}

func NewLab(name string) *Lab {
	if name != "" {
		return nil
	}
	l := Lab{"lab10", "out", "kmattson", time.Now()}
	return &l
}
*/

func MakeTestServer() *negroni.Negroni {
	return NewServer()
}
