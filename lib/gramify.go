package lib

import (
	"strings";
)

func isNewSentence(gram string) bool{

    // punctuationMarks := []string {".", "!", "?"}

    for _,token := range punctuationMarks {
        if gram[len(gram)-1:] == token {
            return true
        }
    }
    return false
}

func GramifyWords(text string, n int) (ngrams map[string][]string, starts []string) {

    ngrams = make(map[string][]string)

    // split source text into []string
    var words []string = strings.Fields(text)

    var wordCount int = len(words)
    var gram string

    var nextStartIndex int

    for i:= 0; i <= wordCount-n; i++ {
        gram = ""
        for j := 0; j < n; j++ {
            gram += words[i+j]
            gram += " "
        }

        // strip trailing whitespace
        gram = strings.Trim(gram, " ")


        // add gram to starts slice
        if  i == nextStartIndex {
            starts = append(starts, gram)
        }

        //
        if isNewSentence(gram) {
            nextStartIndex = i+n
        }

        // if ngram isn't in ngrams map, add it, then add next word to it's slice
        if i != wordCount-n {
            if _,ok := ngrams[gram]; !ok {
                ngrams[gram] = make([]string,0)
            }
            ngrams[gram] = append(ngrams[gram], words[i+n])
        }
    }

    return ngrams, starts
}

