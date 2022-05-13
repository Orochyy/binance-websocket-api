package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func marshal() string {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	return birdJson
}

func main() {
	DoJson, _ := json.Marshal(marshal())
	Sf := fmt.Sprintf(string(DoJson))
	fmt.Println(Sf)
}
