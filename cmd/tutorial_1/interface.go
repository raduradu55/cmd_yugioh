package main

import (
	"fmt"
	"log"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
func stop() {
	updateData(nil)
	log.Fatal("# Sei uscito del gioco")
}
func showMenu(menu []string) {
	for i, option := range menu {
		fmt.Printf("[%v] %v\n", (i + 1), option)
	}
}

var homeMenu = []string{
	"Visualizza personaggi",
	"Duello",
	"Esci",
}

var charsMenu = []string{
	"Aggiungi personaggio",
	"Rimuovi personaggio",
	"Visualizza personaggio",
	"Crea personaggio casuale",
	"Indietro",
	"Esci",
}

var charMenu = []string{
	"Modifica deck",
	"Elimina deck",
	"Cambia nome deck",
	"Indietro",
	"Esci",
}

func printChars(index bool) {
	for i, char := range state.Characters {
		if index {
			fmt.Printf("[%v] %v\n", (i + 1), char.Name)
		} else {
			fmt.Printf("%v\n", char.Name)
		}
	}
}

func removeCharMenu(message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	fmt.Println("## Rimuovendo personaggi ##")
	printChars(true)

	fmt.Println("[0] Indietro")

	var input uint8
	fmt.Printf("\n--> ")
	fmt.Scan(&input)

	switch input {
	case 0:
		viewCharsMenu("")
	default:
		if input > uint8(len(state.Characters)) {
			removeCharMenu("! Input invalido")
		} else {
			removeChar(input - 1)
			viewCharsMenu("# Personaggio rimosso")
		}
	}
}

func printDecks(character Character, showIndex bool) {
	for i, deck := range character.Decks {
		if showIndex {
			fmt.Printf("[%v] %v\n", (i + 1), deck.Name)
		} else {
			fmt.Printf("- %v [%v]\n", deck.Name, deck.Cards)
		}
	}
}

func chooseDeck(charIndex uint8) uint8 {
	fmt.Println("## Scegli deck ##")
	printDecks(state.Characters[charIndex], true)
	fmt.Println("[0] Indietro")

	var input uint8
	fmt.Printf("\n--> ")
	fmt.Scan(&input)

	if input == 0 {
		viewChar(charIndex, "")
	} else {
		return input
	}
	return 0
}

func deleteDeck(charIndex uint8, message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	deckIndex := chooseDeck(charIndex)
	if deckIndex != 0 {
		emptyDeck(charIndex, (deckIndex - 1))
		viewChar(charIndex, "# Deck eliminato")
	} else {
		viewChar(charIndex, "! Deck non trovato")
	}
}

func printMonstList() {
	fmt.Println("- Monsters")
	for _, monst := range state.MonstList {
		fmt.Printf("[%v] %v  ", string(monst.Id), monst.Name)
	}
	fmt.Printf("\n")
}

func printSpellList() {
	fmt.Println("- Spells")
	for _, spell := range state.SpellList {
		fmt.Printf("[%v] %v  ", string(spell.Id), spell.Name)
	}
	fmt.Printf("\n")
}

func printModifyInterface(deck Deck) {
	clearScreen()
	printMonstList()
	printSpellList()
	fmt.Printf("# Deck attuale (%v): %v\n", deck.Name, deck.Cards)
	fmt.Printf(
		"- inserisci nuovo deck (max %v carte, min %v carte)\n",
		MAX_CARDS,
		MIN_CARDS,
	)
	for i := 0; i <= MAX_CARDS; i++ {
		if i == (MIN_CARDS - 1) {
			fmt.Printf("m")
		} else {
			fmt.Printf("x")
		}

	}
	fmt.Printf("|\n")
}

func modifyDeckCards(charIndex uint8, message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	deckIndex := chooseDeck(charIndex)
	deck := state.Characters[charIndex].Decks[(deckIndex - 1)]

	printModifyInterface(deck)

	var newCards string
	fmt.Scan(&newCards)

	if len(newCards) > MAX_CARDS {
		viewChar(charIndex, fmt.Sprint("! Massimo ", MAX_CARDS, " carte"))
	} else if len(newCards) < MIN_CARDS {
		viewChar(charIndex, fmt.Sprint("! Minimo ", MIN_CARDS, " carte"))
	} else if isDeckValid(newCards) {
		updateDeck(
			charIndex,
			Deck{Name: deck.Name, Cards: newCards},
			(deckIndex - 1),
		)
		viewChar(charIndex, "# Deck aggiornato")
	} else {
		viewChar(charIndex, "! Deck invalido (trovata carta inesistente)")
	}
}

func modifyDeckName(charIndex uint8, message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	deckIndex := chooseDeck(charIndex)
	deck := state.Characters[charIndex].Decks[(deckIndex - 1)]
	clearScreen()

	fmt.Println("## Modificando nome deck ##")
	fmt.Printf("- Nome attuale: %v\n", deck.Name)
	fmt.Printf("Nome nuovo (senza spazi):\n--> ")

	var newName string
	fmt.Scan(&newName)

	updateDeck(
		charIndex,
		Deck{Name: newName, Cards: deck.Cards}, (deckIndex - 1),
	)
	viewChar(charIndex, "# Nome deck cambiato")
}

func viewChar(charIndex uint8, message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	fmt.Printf("## Visualizzando %v ##\n", state.Characters[charIndex].Name)
	fmt.Println("# Decks: ")

	printDecks(state.Characters[charIndex], false)
	showMenu(charMenu)

	var input uint8
	fmt.Printf("\n--> ")
	fmt.Scan(&input)

	switch input {
	case 1:
		modifyDeckCards(charIndex, "")
	case 2:
		deleteDeck(charIndex, "")
	case 3:
		modifyDeckName(charIndex, "")
	case uint8(len(charMenu) - 1):
		viewCharsMenu("")
	case uint8(len(charMenu)):
		stop()
	default:
		viewChar(charIndex, "! Input invalido")
	}
}

func viewCharMenu(message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	fmt.Println("## Scegli personaggio da visualizzare ##")
	printChars(true)
	fmt.Println("[0] Indietro")

	var input uint8
	fmt.Printf("\n--> ")
	fmt.Scan(&input)

	switch input {
	case 0:
		viewCharsMenu("")
	default:
		if input > uint8(len(state.Characters)) {
			viewCharMenu("! Input invalido")
		} else {
			viewChar(input-1, "")
		}
	}
}

func viewCharsMenu(message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	fmt.Println("## Personaggi ##")

	if state.Characters == nil {
		fmt.Println("- Non ci sono personaggi")
	} else {
		printChars(false)
	}

	showMenu(charsMenu)

	var input uint8
	fmt.Printf("\n-> ")
	fmt.Scan(&input)

	switch input {
	case 1:
		addCharMenu()
	case 2:
		removeCharMenu("")
	case 3:
		viewCharMenu("")
	case 4:
		createRngCharacter()
		viewCharsMenu("# Personaggio creato")
	case uint8(len(charsMenu) - 1):
		home("")
	case uint8(len(charsMenu)):
		stop()
	default:
		viewCharsMenu("! Input invalido")
	}
}

func addCharMenu() {
	var name string
	fmt.Printf("- Nome personaggio\n-> ")
	fmt.Scan(&name)

	addCharacter(name)
	viewCharsMenu("# Personaggio aggiunto")
}

func startGame() {
	home("")
}

func chooseChar(err_message string, message string) uint8 {
	clearScreen()
	if err_message != "" {
		fmt.Println(err_message)
	}

	fmt.Println(message)
	printChars(true)
	fmt.Println("[0] Indietro")

	var input uint8
	fmt.Printf("-> ")
	fmt.Scan(&input)

	switch input {
	case 0:
		home("")
	default:
		if input > uint8(len(state.Characters)) {
			home("! Input invalido")
		} else {
			return input - 1
		}
	}

	return 0
}

func duel() {
	clearScreen()
	var charIndex uint8 = chooseChar("", "## Duello, scegli personaggio ##")
	var enemyIndex uint8 = chooseChar("", "## Duello, scegli il avversario ##")

	state.Match = Match{
		Round: 0,
		P1LP:  LP,
		P2LP:  LP,
		P1:    state.Characters[charIndex],
		P2:    state.Characters[enemyIndex],
		P1T: Table{
			MonsterField: [4]Monster{},
			SpellField:   [4]Spell{},
			Graveyard: Graveyard{
				Monsters: []Monster{},
				Spells:   []Spell{},
			},
		},
		P2T: Table{
			MonsterField: [4]Monster{},
			SpellField:   [4]Spell{},
			Graveyard: Graveyard{
				Monsters: []Monster{},
				Spells:   []Spell{},
			},
		},
	}

	startDuel()
}

func home(message string) {
	clearScreen()
	if message != "" {
		fmt.Println(message)
	}
	fmt.Println("##! YuGiOh-dei-poveri !##")
	showMenu(homeMenu)

	var input uint8
	fmt.Printf("-> ")
	fmt.Scan(&input)

	switch input {
	case 1:
		viewCharsMenu("")
	case 2:
		duel()
	case uint8(len(homeMenu)):
		stop()
	default:
		home("! Input invalido")
	}
}
