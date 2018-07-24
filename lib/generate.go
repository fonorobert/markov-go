package lib

import (
	"fmt";
	"math/rand";
	"time";
	"strings";
)

func Generate(ngrams map[string][]string, start string, n int, length int, limitType string) (newText string) {

    rSource := rand.NewSource(time.Now().UnixNano())
    r := rand.New(rSource)

    // punctuationMarks := []string{".", "!", "?"}



    curGram := start
    var possibilities []string
    var next string
    var newTextSlice []string = strings.Fields(curGram)

    for i:= 0; i < length; {
        if poss, ok := ngrams[curGram]; ok {
            possibilities = poss
            next = possibilities[r.Intn(len(possibilities))]
            newTextSlice = append(newTextSlice, next)

            if limitType == "sentence" {
                // increment i if at end of sentence
                for _,token := range punctuationMarks {
                    if next[len(next)-1:] == token {
                        i++
                        break
                    }
                }
            } else if limitType == "word" {
                i++
            } else {
                panic(fmt.Sprintf("Invalid limitType '%s'\n", limitType))
            }

            // fmt.Printf("curGram: %s, poss: %s, next: %s\n", curGram, poss, next)
            // fmt.Printf("%s\n", newTextSlice)

            curGram = ""
            for j:=n; j>=1; j-- {
                if j<len(newTextSlice){
                    curGram += newTextSlice[len(newTextSlice)-j] + " "
                }
                continue
            }
            curGram = strings.Trim(curGram, " ")
        } else {
            return
            panic(fmt.Sprintf("No such gram as '%s'\n", curGram))
        }
    }

    for _,gram := range newTextSlice {
        newText += gram
        newText += " "
    }
    newText = strings.Trim(newText, " ")

    return newText
}

