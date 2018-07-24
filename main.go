package main

import (
		"fmt";
		"os";
		// "strings";
		"math/rand";
		"time";
		"io/ioutil";
		"strconv";
		lib "github.com/fonorobert/markov-go/lib";
		)

func main() {

	rSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rSource)

	args := os.Args[1:]

	if len(args)!=4 {
		fmt.Println("Usage: markov-go SOURCE_PATH GRAM_ORDER RESULT_LENGTH LIMIT_TYPE")
		return
	}

	var e error
	var n int
	var length int

	if n, e = strconv.Atoi(args[1]); e != nil {
		panic(e)
	}
	if length, e = strconv.Atoi(args[2]); e != nil {
		panic(e)
	}

	var limitType string = args[3]

	if limitType != "sentence" && limitType != "word" {
		fmt.Println("Limit type can only be \"sentence\" or \"word\".")
		return
	}

	var path string = args[0]

	var sourceText string

	if b, err := ioutil.ReadFile(path); err == nil {
		sourceText = string(b)
	} else {
		panic(err)
	}

	grams, starts:= lib.GramifyWords(sourceText, n)
	var start string = starts[r.Intn(len(starts))]
	var newText string = lib.Generate(grams, start, n, length, limitType)
	fmt.Println(newText)
}
