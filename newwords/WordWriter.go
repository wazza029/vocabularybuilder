package newwords

import (
	"encoding/json"
	"fmt"
	"os"
)

func WordWriter() (store *os.File, err error) {

	//WordWriter registers new words that one wants to learn

	//Creating a new file to store words. This file will be later used in main program
	store, err = os.Create("store.txt")

	if err != nil {
		fmt.Println("Error during creating a file: ", err)
	}

	//Getting input from the user and storing
	//in the map
	database := make(map[string]string)
	var spanishinput string
	var englishinput string
	spanishinput = "amigo"
	englishinput = "friend"
	//fmt.Println("Provide spanish word: ")
	//fmt.Scanln(&spanishinput)
	//fmt.Println("Provide english translation: ")
	//fmt.Scanln(&englishinput)
	database[spanishinput] = englishinput

	//Transforming into JSON
	bs, err := json.Marshal(database)
	store.Write(bs)
	return store, err

}
