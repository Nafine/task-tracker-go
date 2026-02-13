package main

import (
	"log"

	"github.com/Nafine/task-tracker/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
