package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := os.Args[1]
	in, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("package gopcua\n\n")
	fmt.Printf("// This file was auto-generated from %v. Do not change it manually! \n", file)

	fmt.Println(`
/*
#cgo CFLAGS: -std=c99 -I .
#include <stdlib.h>
#include "gopcua_client.h"
 
*/
import "C"
`)

	lines := strings.Split(string(in), "\n")
	pos := 0

	for pos < len(lines) {
		words := strings.Split(lines[pos], " ")
		if words[0] == "typedef" {
			switch {
			case words[1] == "struct" && strings.Contains(lines[pos], "{"):
				// find closing brackets
				for {
					pos++
					if len(lines[pos]) > 0 && lines[pos][0] == '}' {
						w := strings.Split(lines[pos], " ")
						fmt.Printf("type %v C.%v // line: %v \n", strings.Replace(w[1], ";", "", -1), strings.Replace(w[1], ";", "", -1), pos+1)
						break
					}
				}
			case words[1] == "enum":
				orig := pos
				t := ""
				// find next name of the enum typ and map it, like a struct.
				for {
					pos++
					if len(lines[pos]) > 0 && lines[pos][0] == '}' {
						w := strings.Split(lines[pos], " ")
						t = strings.Replace(w[1], ";", "", -1)
						fmt.Printf("type %v C.%v // line: %v \n", t, t, pos+1)
						break
					}
				}
				// declare the const values
				fmt.Printf("const (\n")
				for i := orig + 1; i < pos; i++ {
					l := strings.Split(lines[i], " //")
					l[0] = strings.Replace(l[0], ",", "", -1)
					if strings.Contains(l[0], "=") {
						n := "= " + t + "("
						l[0] = strings.Replace(l[0], "= ", n, -1)
						l[0] = l[0] + ")"
					} else {
						if i == orig+1 {
							l[0] = l[0] + " = iota"
						}
					}
					if len(l) > 1 {
						l[0] = l[0] + " //" + l[1]
					}
					fmt.Printf("%v\n", l[0])
				}
				fmt.Printf(")\n")
			case words[1] == "void":
				//do nothing
			default:
				if !strings.Contains(lines[pos], "(") {
					fmt.Printf("type %v C.%v // line: %v \n", strings.Replace(words[2], ";", "", -1), strings.Replace(words[2], ";", "", -1), pos+1)
				}
			}
		}
		pos++
	}
}
