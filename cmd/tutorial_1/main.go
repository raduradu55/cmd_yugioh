// -- yugioh like game
// data stored in json file
// can have cards
// card types: monster, spell, trap, field
// can put cards in a deck (max 30 cards, min 20, max 3 cards of the same card)
// you have to create a character to play
// can create a deck (with a name)

package main

import (
	"encoding/json"
	"log"
	"os"
	"syscall"
)

var state State

const DATA_FILE = "cmd\\tutorial_1\\data.json"

func loadData() {
	jsonData := make(map[string]interface{})
	jsonFile, err := os.ReadFile(DATA_FILE)
	if err != nil {
		if pe, ok := err.(*os.PathError); ok && pe.Err == syscall.ENOENT {
			updateData(createNewData())
			loadData()
			return
		}
	}
	err = json.Unmarshal([]byte(jsonFile), &jsonData)
	if err != nil {
		log.Fatalf("unmarshal brutto %s", err)
	}
	state.jsonData = jsonData
}

func updateData(data map[string]interface{}) {
	var toEncode map[string]interface{}
	if data == nil {
		toEncode = state.jsonData
	} else {
		toEncode = data
	}

	newJson, err := os.Create(DATA_FILE)
	if err != nil {
		log.Fatalf("create file brutto %s", err)
	}
	defer newJson.Close()

	encoder := json.NewEncoder(newJson)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(toEncode); err != nil {
		log.Fatalf("encoding bad %s", err)
	}
}

func main() {
	loadData()
	defer updateData(nil)
	loadCards()
	loadChars()

	startGame()
}
