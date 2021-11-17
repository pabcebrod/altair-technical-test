package team

import (
	"github.com/altair-tecnical-test/src/data"
	"github.com/altair-tecnical-test/src/utils"
	"strings"
)

type TeamManager interface {
	SplitListByTeam()
	IsPersonWordPalindrome(id string)
	HasPersonWordFiveVowels(id string)
}

type Manager struct {
	BlueTeam []data.Person
	RedTeam  []data.Person
}

//NewManager returns a new manager instance and splits the initial list into two teams
func NewManager() *Manager {
	team := new(Manager)
	team.SplitListByTeam(data.People)

	return team
}

//SplitListByTeam divided the initial list of team in two
//depending on the team they belong to
func (m *Manager) SplitListByTeam(people []data.Person) {
	for _, person := range people {
		if person.Equipo == data.RED {
			m.RedTeam = append(m.RedTeam, person)
			continue
		}
		m.BlueTeam = append(m.BlueTeam, person)
	}
}

//IsPersonWordPalindrome iterates the RedTeam list searching for a person whose ID is equal to the provided one.
//If the word associated with that person is palindrome returns true, false otherwise.
//In case there is no person with that ID returns an error.
func (m *Manager) IsPersonWordPalindrome(id string) (res bool, err error) {
	person, err := getPersonById(id, m.RedTeam)

	if err != nil {
		return
	}

	res = isPalindrome(strings.Split(person.Palabra, ""))

	return
}

//HasPersonWordFiveVowels iterates the BlueTeam list searching for a person whose ID is equal to the provided one.
//If the word associated contains 5 vowels returns true, false otherwise.
//In case there is no person with that ID returns an error.
func (m *Manager) HasPersonWordFiveVowels(id string) (res bool, err error) {
	person, err := getPersonById(id, m.BlueTeam)

	if err != nil {
		return
	}

	numberOfVowels := 0
	for _, letter := range strings.Split(person.Palabra, "") {

		if isVowel(letter) {
			numberOfVowels++
		}
	}

	if numberOfVowels == 5 {
		res = true
	}
	return
}

//----------------------------------------------------------------------------
//-------------------------Unexported methods---------------------------------
//----------------------------------------------------------------------------

//getPersonById iterates the BlueTeam list searching for a person whose ID is equal to the provided one.
//In case there is no person with that ID returns an error.
func getPersonById(id string, team []data.Person) (person data.Person, err error) {

	for _, teamMember := range team {
		if teamMember.Id == id {
			person = teamMember
			break
		}
	}

	if person.Id == "" {
		err = utils.NO_PERSON_ID
	}
	return
}

//isVowel returns true if the received letter is a vowel, false otherwise.
func isVowel(letter string) bool {

	vowels := [5]string{"a", "e", "i", "o", "u"}

	for _, vowel := range vowels {
		if vowel == strings.ToLower(letter) {
			return true
		}
	}
	return false
}

//isPalindrome() returns true if the received word is a palindrome, false otherwise.
func isPalindrome(word []string) bool {
	wordLength := len(word)

	for i := 0; i <= wordLength/2; i++ {
		if word[i] != word[wordLength-i-1] {
			return false
		}
	}

	return true
}
