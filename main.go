package main

import (
	"042021/26042021/nhl1/roster"
	"042021/26042021/nhl1/teams"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	fOpen, err := os.OpenFile("teams.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("cant open file: %v", err)
	}

	teamsnum, err := teams.GetAllTeams()
	if err != nil {
		log.Fatalf("cant get all the teams %v", err)
	}
	write := io.MultiWriter(os.Stdout, fOpen)
	log.SetOutput(write)

	var wg sync.WaitGroup
	wg.Add(len(teamsnum))
	results := make(chan []roster.Roster)

	for _, team := range teamsnum {
		go func(team teams.Team) {
			roster, err := roster.GetAllRosters(team.ID)
			if err != nil {
				fmt.Printf("cant get roster: %v", err)
			}
			results <- roster
			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
	}()
	display(results)
}
func display(results chan []roster.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("----------------------")
			log.Printf("ID: %d\n", ros.ID)
			log.Printf("Name: %s\n", ros.Fullname)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.Position)
			log.Println("----------------------")
		}
	}
}
