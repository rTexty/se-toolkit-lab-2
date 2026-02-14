package main

import (
	"encoding/json"
	"os"

	
)

func loadOutcomes() ([]Outcome, error) {
	var root Root
	
	path := os.Getenv("OUTCOMES_PATH")
	if path == "" {
		path = "/app/data/outcomes.json"
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	
	err = json.Unmarshal(bytes, &root)
	if err != nil {
		return nil, err
	}
	
	return root.Outcomes, nil
}