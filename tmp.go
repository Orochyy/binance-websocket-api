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
	//var bird Bird
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	tmp, _ := json.Marshal(birdJson)
	Sf := fmt.Sprintf(string(tmp))
	return Sf
}

func main() {
	DoJson, _ := json.Marshal(marshal())
	Sf := fmt.Sprintf(string(DoJson))
	fmt.Println(Sf)
}
