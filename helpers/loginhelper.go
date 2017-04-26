package helpers

import (
	"encoding/json"
	"log"
	"os"
)

func CheckToken(token string) (valid bool) {

	services, err := parseVcapServices(os.Getenv("VCAP_SERVICES"))
	if err != nil {
		log.Printf("Error parsing information: %v\n", err.Error())
		return
	}
	info := services["user-provided"]
	if len(info) == 0 {
		log.Printf("No slack-services are bound to this application.\n")
		return
	}
	if token != info[0].Credentials.Token {
		return false
	}
	return true

}

func parseVcapServices(vcapStr string) (VcapServices, error) {
	var vcapServices VcapServices

	if err := json.Unmarshal([]byte(vcapStr), &vcapServices); err != nil {
		return vcapServices, err
	}
	return vcapServices, nil
}

type VcapServices map[string][]VcapService

type VcapService struct {
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	// Genenal field
	Token string `json:"token,omitempty"`
}
