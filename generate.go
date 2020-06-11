package main

import (
	"github.com/solo-io/external-apis/codegen"
	"log"
)

func main() {
	log.Println("starting skv2 code generation")
	skv2Cmd := codegen.Command()
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Finished generating skv2 code\n")
}
