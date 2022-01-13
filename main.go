package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

func main() {
	parser := syntax.NewParser()
	file, err := parser.Parse(os.Stdin, "env")
	if err != nil {
		log.Fatal(err)
	}
	var value bytes.Buffer
	printer := syntax.NewPrinter()
	syntax.Walk(file, func(node syntax.Node) bool {
		if decl, ok := node.(*syntax.DeclClause); ok && decl.Variant.Value == "export" {
			for _, assign := range decl.Args {
				if len(decl.Args) != 1 {
					log.Fatalf("expected export to have 1 argument, got %d", len(decl.Args))
				}
				value.Reset()
				if err := printer.Print(&value, decl.Args[0].Value); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s=%s\n", assign.Name.Value, value.String())
			}
		}
		return true
	})
}
