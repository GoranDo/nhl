package roster

import (
	"042021/26042021/nhl1/teams"
	"encoding/json"
	"fmt"
	"net/http"
)

type Roster struct {
	ID           int    `json:"id"`
	Fullname     string `json:"fullName"`
	Link         string `json:"link"`
	Jerseynumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}

type ResponseRoster struct {
	Rosters []Roster `json:"roster"`
}

func GetAllRosters(teamID int) ([]Roster, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams/%d/roster", teams.BaseURL, teamID))
	if err != nil {
		fmt.Printf("cant get response %v", err)
	}
	response := ResponseRoster{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Printf("cant decode response: %v", err)
	}
	return response.Rosters, err
}
