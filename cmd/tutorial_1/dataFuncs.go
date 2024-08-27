package main

import (
	"math/rand"
	"time"
)

func createMonsters() []Monster {
	cards := make([]Monster, 0, 2)
	cards = append(cards, Monster{
		Card: Card{
			Id:          'a',
			Name:        "monster1",
			Description: "",
			Destroyed:   false,
		},
		Attack: 1000, Defense: 1000,
	})
	cards = append(cards, Monster{
		Card: Card{
			Id:          'b',
			Name:        "monster2",
			Description: "",
			Destroyed:   false,
		},
		Attack: 1500, Defense: 900,
	})
	cards = append(cards, Monster{
		Card: Card{
			Id:          'c',
			Name:        "monster3",
			Description: "",
			Destroyed:   false,
		},
		Attack: 500, Defense: 2000,
	})
	return cards
}

func createSpells() []Spell {
	cards := make([]Spell, 0, 2)
	cards = append(cards, Spell{
		Card: Card{
			Id:          'A',
			Name:        "unicorn horn",
			Description: "",
			Destroyed:   false,
		},
		Effect: "", Spell_type: Equip,
	})
	cards = append(cards, Spell{
		Card: Card{
			Id:          'B',
			Name:        "sougen",
			Description: "",
			Destroyed:   false,
		},
		Effect: "", Spell_type: Field,
	})
	return cards
}

func createRngCharacter() {

	var decks [3]Deck
	decks[0] = createRngDeck()
	decks[1] = createRngDeck()
	decks[2] = createRngDeck()
	newChar := Character{Name: "rngChar", Decks: decks}

	var characters []interface{}
	if state.jsonData["characters"] != nil {
		characters = state.jsonData["characters"].([]interface{})
	} else {
		characters = []interface{}{}
	}

	characters = append(characters, newChar)
	state.Characters = append(state.Characters, newChar)
	state.jsonData["characters"] = characters
}

func updateDeck(charIndex uint8, newDeck Deck, deckIndex uint8) {
	state.Characters[charIndex].Decks[deckIndex].Cards = newDeck.Cards
	state.Characters[charIndex].Decks[deckIndex].Name = newDeck.Name

	charList := state.jsonData["characters"].([]interface{})
	for _, char := range charList {
		char := char.(map[string]interface{})
		if char["Name"] == state.Characters[charIndex].Name {
			var deck = char["Decks"].([]interface{})[deckIndex].(map[string]interface{})
			deck["Cards"] = newDeck.Cards
			deck["Name"] = newDeck.Name
		}
	}
	state.jsonData["characters"] = charList

}

func isDeckValid(deck string) bool {

	valid := false

	var monstIds []rune
	for _, monst := range state.MonstList {
		monstIds = append(monstIds, monst.Id)
	}

	for _, card_id := range deck {
		for _, id := range monstIds {
			if id == card_id {
				valid = true
				break
			}
			valid = false
		}
	}
	return valid
}

func createRngDeck() Deck {
	// seed
	rand.NewSource(time.Now().UnixNano())
	deck := Deck{Name: "prova"}

	var cards string
	// monsters
	for i := 0; i < MAX_CARDS/2; i++ {
		cards += string(state.MonstList[rand.Intn(len(state.MonstList))].Id)
	}
	// spells
	for i := MAX_CARDS / 2; i < MAX_CARDS; i++ {
		cards += string(state.SpellList[rand.Intn(len(state.SpellList))].Id)
	}
	deck.Cards = cards
	return deck
}

func createNewData() map[string]interface{} {
	return map[string]interface{}{
		"cards": map[string]interface{}{
			"monsters": createMonsters(),
			"spells":   createSpells(),
			"traps":    nil,
		},
		// "characters": [...]Character{createRngCharacter()},
		"characters": nil,
	}
}

func addCharacter(name string) {
	var characters []interface{}
	if state.jsonData["characters"] != nil {
		characters = state.jsonData["characters"].([]interface{})
	} else {
		characters = []interface{}{}
	}
	var decks [3]Deck
	newChar := Character{Name: name, Decks: [3]Deck(decks)}

	characters = append(characters, newChar)
	state.Characters = append(state.Characters, newChar)
	state.jsonData["characters"] = characters
}

func emptyDeck(charIndex uint8, deckIndex uint8) {
	state.Characters[charIndex].Decks[deckIndex].Cards = ""

	charList := state.jsonData["characters"].([]interface{})
	for _, char := range charList {
		char := char.(map[string]interface{})
		if char["Name"] == state.Characters[charIndex].Name {
			var deck = char["Decks"].([]interface{})[deckIndex].(map[string]interface{})
			deck["Cards"] = ""
		}
	}
	state.jsonData["characters"] = charList
}

func removeChar(index uint8) {

	state.Characters = append(
		state.Characters[:index],
		state.Characters[index+1:]...,
	)

	charList := state.jsonData["characters"].([]interface{})
	charList = append(charList[:index], charList[index+1:]...)
	state.jsonData["characters"] = charList
}

func loadSpellList(cards map[string]interface{}) []Spell {
	spells := cards["spells"].([]interface{})
	var spellList []Spell

	for _, spell := range spells {
		spell := spell.(map[string]interface{})
		spellList = append(spellList, Spell{
			Card: Card{
				Id:          rune(spell["Id"].(float64)),
				Name:        spell["Name"].(string),
				Description: spell["Description"].(string),
				Destroyed:   spell["Destroyed"].(bool),
			},
			Effect:     spell["Effect"].(string),
			Spell_type: uint8(spell["Spell_type"].(float64)),
		})
	}

	return spellList
}

func loadMonstList(cards map[string]interface{}) []Monster {

	monsters := cards["monsters"].([]interface{})
	var monstList []Monster

	for _, monster := range monsters {
		monster := monster.(map[string]interface{})
		monstList = append(monstList, Monster{
			Card: Card{
				Id:          rune(monster["Id"].(float64)),
				Name:        monster["Name"].(string),
				Description: monster["Description"].(string),
				Destroyed:   monster["Destroyed"].(bool),
			},
			Attack:  uint16(monster["Attack"].(float64)),
			Defense: uint16(monster["Defense"].(float64)),
		})
	}
	return monstList
}

func loadCards() {
	cards := state.jsonData["cards"].(map[string]interface{})
	state.MonstList = loadMonstList(cards)
	state.SpellList = loadSpellList(cards)
}

func loadChars() {

	if state.jsonData["characters"] == nil {
		return
	}

	for _, character := range state.jsonData["characters"].([]interface{}) {
		character := character.(map[string]interface{})
		var decks = character["Decks"].([]interface{})

		var deckList [3]Deck
		for i, deck := range decks {
			deck := deck.(map[string]interface{})
			deckList[i] = Deck{
				Name:  deck["Name"].(string),
				Cards: deck["Cards"].(string),
			}
		}

		state.Characters = append(state.Characters, Character{
			Name:  character["Name"].(string),
			Decks: deckList,
		})
	}
}
