package main

import (
	"fmt"
	"math/rand"
	"time"
)

var database = map[string]string{}

var translation string
var word string
var correctTranslation string

func main() {

	//newwords.WordWriter()
	//newwords.WordWriter()
	//newwords.WordWriter()
	//newwords.WordWriter()
	//newwords.WordWriter()

	word = RandomWord(database)
	fmt.Println(word)
	fmt.Println("Enter translation: ")
	fmt.Scanln(&translation)
	correctTranslation = database[word]

	if correctTranslation == translation {
		fmt.Println("Well done")
	} else {
		fmt.Println("Not good")
	}

}

//func WordWriter(s string, e string) map[string]string {
//
//	//WordWriter registers new words in map, key in spanish, value in english
//	database[s] = e
//	return database
//
//}

func RandomWord(m map[string]string) string {

	// 1) Solution with using break, no control over randomization via range: https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/
	//	var key string
	//	for k, _ := range m {
	//		key = k
	//		break
	//	}
	//	return key

	// 2) Using randmap function: https://github.com/lukechampine/randmap Blog: https://lukechampine.com/hackmap.html

	// 3) Flattening the map:

	keys := []string{}

	for i, _ := range m {

		keys = append(keys, i)

	}

	rand.Seed(time.Now().UnixNano())
	r := keys[rand.Intn(len(keys))]
	return r
}
