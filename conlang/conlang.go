package main

import (
	"fmt"
	"math/rand"
	"slices"
)

type Letter struct {
	Name       string
	Values     []string
	IsOptional bool
}

type Syllable struct {
	Letters []Letter
}

type Word struct {
	Syllable Syllable
	MinSylls int
	MaxSylls int
}

func GenSyllable(syllable Syllable) string {
	syll := ""
	value := ""
	chosenLetters := map[string]string{}

	for _, letter := range syllable.Letters {
		// Because it's optional, it may be omitted.
		if letter.IsOptional && rand.Intn(2) == 0 {
			continue
		}
		// BEGIN Conditions
		// SKIP G IF C1 IS J OR W
		if letter.Name == "G" {
			if val, ok := chosenLetters["C1"]; ok {
				vals := []string{"j", "w"}
				if slices.Contains(vals, val) {
					continue
				}
			}
		}
		if letter.Name == "C3" {
			if val, ok := chosenLetters["C2"]; ok {
				vals := []string{"r", "s"}
				if slices.Contains(vals, val) {
					continue
				}
			}
		}
		// END

		value = letter.Values[rand.Intn(len(letter.Values))]
		chosenLetters[letter.Name] = value
		syll += value
	}

	return syll
}

func GenWord(word Word) string {
	genWord := ""
	numSyllables := rand.Intn((word.MaxSylls-word.MinSylls)+1) + word.MinSylls

	for syll := 0; syll < numSyllables; syll++ {
		genWord += GenSyllable(word.Syllable) + "|"
	}

	return genWord
}

func main() {
	// BEGIN Declarations + definition
	c1 := Letter{
		Name:       "C1",
		Values:     []string{"p", "t", "k", "m", "n", "b", "f", "th", "d", "s", "z", "sh", "c", "x", "g", "l", "r"},
		IsOptional: true,
	}
	c2 := Letter{
		Name:       "C2",
		Values:     []string{"p", "t", "k", "m", "n", "b", "f", "th", "d", "s", "z", "sh", "c", "x", "g", "l", "r"},
		IsOptional: true,
	}
	v := Letter{
		Name:       "V",
		Values:     []string{"a", "e", "i", "o", "u"},
		IsOptional: false,
	}
	g := Letter{
		Name:       "G",
		Values:     []string{"j", "w"},
		IsOptional: true,
	}
	c3 := Letter{
		Name:       "C3",
		Values:     []string{"r", "s"},
		IsOptional: true,
	}
	syllable := Syllable{
		Letters: []Letter{c1, g, v, c2, c3},
	}
	word := Word{
		Syllable: syllable,
		MinSylls: 1,
		MaxSylls: 4,
	}
	// END

	// BEGIN Printing
	for wrd := 0; wrd < 10; wrd++ {
		fmt.Println(GenWord(word))
	}
	// END
}
