package main

import (
	"bufio"
	"log"
	"os"

	"mvdan.cc/sh/v3/syntax"
)

func main() {
	input := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		input = f
	}
	parser := syntax.NewParser()
	file, err := parser.Parse(input, "")
	if err != nil {
		log.Fatal(err)
	}
	printer := syntax.NewPrinter()
	output := bufio.NewWriter(os.Stdout)
	syntax.Walk(file, func(node syntax.Node) bool {
		if decl, ok := node.(*syntax.DeclClause); ok && decl.Variant.Value == "export" {
			if len(decl.Args) != 1 {
				log.Fatalf("expected export to have 1 argument, got %d", len(decl.Args))
			}
			for _, assign := range decl.Args {
				output.WriteString(assign.Name.Value)
				output.WriteByte('=')
				printer.Print(output, decl.Args[0].Value)
				output.WriteByte('\n')
			}
		}
		return true
	})
	if err := output.Flush(); err != nil {
		log.Fatal(err)
	}
}
