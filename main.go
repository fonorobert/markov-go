package main

import (
		"fmt";
		"io/ioutil";
		"flag";
		lib "github.com/fonorobert/markov-go/lib";
		)

func main() {

	// get command line arguments

	var n int
	var length int
	var word bool

	flag.IntVar(&n, "n", 1, "gram order")
	flag.IntVar(&length, "l", 1, "text length to generate")
	flag.BoolVar(&word, "w", false, "generate number of words instead of sentences")

	flag.Parse()

	var path string = flag.Arg(0)
	var limitType string = "sentence"

	if word {
		limitType = "word"
	}

	// check if path is present

	if path == "" {
		fmt.Println("Usage: markov-go [OPTIONS] PATH/TO/SOURCE/TEXT")
		return
	}

	var sourceText string

	if b, err := ioutil.ReadFile(path); err == nil {
		sourceText = string(b)
	} else {
		panic(err)
	}

	grams, starts := lib.GramifyWords(sourceText, n)
	var newText string = lib.Generate(grams, starts, n, length, limitType)
	fmt.Println(newText)
}
