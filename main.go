package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var input arrayFlags
var output string

func main() {
	flag.Var(&input, "input", "Input files. You can add this option multiple times: --input file1.json --input file2.json")
	flag.Var(&input, "i", "shorthand for --input")
	flag.StringVar(&output, "output", "notes.md", "Output file: --output out.md")
	flag.StringVar(&output, "o", "notes.md", "Shorthand for --output")
	flag.Parse()

	if len(input) == 0 {
		log.Fatalln("No input files specified.")
	}
	if output == "" {
		log.Fatalln("No output file specified.")
	}

	outFile, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalln("Error creating/opening output file:", err.Error())
	}
	defer func() {
		err = outFile.Close()
		if err != nil {
			log.Println("Error closing output file:", err.Error())
		}
	}()

	var writeError error
	write := func(s string) {
		_, err := outFile.WriteString(s)
		if err != nil && writeError == nil {
			writeError = err
		}
	}

	for _, in := range input {
		inFile, err := ioutil.ReadFile(in)
		if err != nil {
			log.Fatalln("Error opening input file", in, "\n", err.Error())
		}

		var doc Document

		err = json.Unmarshal(inFile, &doc)
		if err != nil {
			log.Fatalln("Error parsing content of file", in, "\n", err.Error())
		}

		for _, n := range doc.Notes {
			if n != nil {
				write("## ")
				write(n.Title)
				write("\n")

				if len(input) > 1 {
					write("\nSource File: ")
					write(in)
				}

				write("\nColour: ")
				write(n.Colour)
				write("  \nOriginal Font Size: ")
				write(fmt.Sprintf("%d", n.FontSize))
				if n.Favoured {
					write("  \nFavoured: ")
					write("Yes")
				}
				if n.HideBody {
					write("  \nBody Hidden: ")
					write("Yes")
				}
				write("\n\n")
				lines := strings.Split(n.Body, "\n")
				first := true
				for _, l := range lines {
					if first {
						first = false
					} else {
						if l != "" && !isList(l) {
							// Add double space on previous line end to make line break
							write("  ")
						}
						write("\n")
					}
					write(l)
				}
				write("\n\n")
			}
		}

		if writeError != nil {
			log.Fatalln("Error writing to output file while processing", in, "\n", err.Error())
		}
	}
}

func isList(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.ContainsAny(trimmed[0:1], "-*+0123456789")
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ";")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Document struct {
	Notes []*Note `json:"notes,omitempty"`
}

type Note struct {
	Colour   string `json:"colour,omitempty"`
	Body     string `json:"body,omitempty"`
	Title    string `json:"title,omitempty"`
	FontSize int    `json:"fontSize,omitempty"`
	HideBody bool   `json:"hideBody,omitempty"`
	Favoured bool   `json:"favoured,omitempty"`
}
