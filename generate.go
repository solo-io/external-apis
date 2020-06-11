package main

import (
	"log"

	"github.com/solo-io/external-apis/codegen"
)

func main() {
	log.Println("starting skv2 code generation")
	skv2Cmd := codegen.Command()
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating skv2 code\n")
}
