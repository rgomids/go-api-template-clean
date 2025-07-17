package main

import (
	"log"
	"os"

	"github.com/rgomids/go-api-template-clean/cli/internal/scaffold"
)

func main() {
	if len(os.Args) < 2 {
		log.Printf("usage: go-api scaffold <Entity> [fields...]")
		return
	}
	if os.Args[1] != "scaffold" {
		log.Printf("unknown command: %s", os.Args[1])
		return
	}
	spec, err := scaffold.Parse(os.Args[2:])
	if err != nil {
		log.Printf("parse error: %v", err)
		return
	}
	if err := scaffold.Generate(spec); err != nil {
		log.Printf("generate error: %v", err)
		return
	}
	if err := scaffold.Update(spec); err != nil {
		log.Printf("update error: %v", err)
		return
	}
}
