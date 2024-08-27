package main

const MAX_CARDS = 30
const MIN_CARDS = 20
const LP int16 = 2000

const (
	Normal = iota
	Quick
	Equip
	Field
	Contin
	Rituals
)

type State struct {
	jsonData   map[string]interface{}
	MonstList  []Monster
	SpellList  []Spell
	Characters []Character
	Match      Match
}

type Deck struct {
	Name  string
	Cards string
}

type Character struct {
	Name  string
	Decks [3]Deck
}

type Card struct {
	Id          rune
	Name        string
	Description string
	Destroyed   bool
}

type Monster struct {
	// id: a-z
	Card
	Attack  uint16
	Defense uint16
}

type Spell struct {
	// id: A-Z
	Card
	Effect     string
	Spell_type uint8
}

type Graveyard struct {
	Monsters []Monster
	Spells   []Spell
}

type Table struct {
	MonsterField [4]Monster
	SpellField   [4]Spell
	Graveyard
}

type Match struct {
	Round uint8
	P1LP  int16
	P2LP  int16
	P1    Character
	P2    Character
	P1T   Table
	P2T   Table
}
