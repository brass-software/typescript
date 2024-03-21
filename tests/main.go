package main

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"

	"github.com/brass-software/typescript"
)

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "tests/go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	if len(pkgs) != 1 {
		panic(len(pkgs))
	}
	importer := typescript.GoImporter{}
	for _, pkg := range pkgs {
		p, err := importer.NewPackage(pkg)
		if err != nil {
			panic(err)
		}
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		// err = p.WriteToDir("tests/ts")
		// if err != nil {
		// 	panic(err)
		// }
	}
}
