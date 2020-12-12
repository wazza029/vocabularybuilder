package newwords

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WordWriter() (store *os.File, err error) {

	//WordWriter registers new words that one wants to learn and stores them in a text file

	//Getting input from the user and storing
	//in the map
	database := make(map[string]string)
	var spanishinput string
	var englishinput string

	//Check if text file with words already exists, if not create one:

	store, err = os.OpenFile("store.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error during initial creation or opening of a text file: ", err)
	}

	// If text file existed from before, recover words and place them in map:
	filecontent, err := ioutil.ReadFile("store.txt")
	if err != nil {
		fmt.Println("Error when reading file content", err)
	}

	if len(filecontent) > 0 {

		err = json.Unmarshal(filecontent, &database)
		if err != nil {
			fmt.Println("Error during unmarshaling", err)
		}
	}

	//For test purposes:
	//spanishinput = "dinero"
	//englishinput = "money"
	fmt.Println("Provide spanish word: ")
	fmt.Scanln(&spanishinput)
	fmt.Println("Provide english translation: ")
	fmt.Scanln(&englishinput)

	database[spanishinput] = englishinput

	//Transforming into JSON and storing in text file
	bs, err := json.Marshal(database)
	err = ioutil.WriteFile("store.txt", bs, 0666)
	if err != nil {
		fmt.Println("Error during writing text file at the end of application", err)
	}
	return store, err

}
