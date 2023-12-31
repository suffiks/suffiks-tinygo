package main

import (
	"encoding/json"
	"os"

	"github.com/dave/jennifer/jen"
)

type arg struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type"`
}

type decl struct {
	Name   string `json:"name"`
	Doc    string `json:"doc"`
	Args   []arg  `json:"args"`
	Return []arg  `json:"return"`
}

func main() {
	decls, err := readFile()
	if err != nil {
		panic(err)
	}

	f := jen.NewFile("env")
	f.HeaderComment("Code generated by gen-env/main.go; DO NOT EDIT.")

	for _, decl := range decls {
		if decl.Doc != "" {
			f.Comment(decl.Doc)
		}
		f.Comment("//go:wasmimport suffiks " + decl.Name)
		f.Func().Id(decl.Name).ParamsFunc(func(g *jen.Group) {
			for _, arg := range decl.Args {
				g.Id(arg.Name).Id(arg.Type)
			}
		}).ParamsFunc(func(g *jen.Group) {
			for _, arg := range decl.Return {
				g.Id(arg.Name).Id(arg.Type)
			}
		})
	}

	if err := f.Save("internal/env/env.go"); err != nil {
		panic(err)
	}
}

func readFile() ([]decl, error) {
	f, err := os.Open("../suffiks/extension/wasi/wasi_env.json")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var decls []decl
	return decls, json.NewDecoder(f).Decode(&decls)
}
