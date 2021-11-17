package server

import (
	"encoding/json"
	"github.com/altair-tecnical-test/src/team"
	"github.com/gorilla/mux"
	"net/http"
)

type Server interface {
	Router() http.Handler
}

type api struct {
	router      http.Handler
	teamManager *team.Manager
}

//New returns a new api instance and defines the API's handler
func New() Server {
	api := &api{}
	api.teamManager = team.NewManager()

	r := mux.NewRouter()
	r.HandleFunc("/teams/palindrome-filter/{ID}", api.isPalindrome).Methods(http.MethodGet)
	r.HandleFunc("/teams/vowel-filter/{ID}", api.hasFiveVowels).Methods(http.MethodGet)
	api.router = r
	return api
}

//Router returns API's handler
func (a *api) Router() http.Handler {
	return a.router
}

//----------------------------------------------------------------------------
//-------------------------Unexported methods---------------------------------
//----------------------------------------------------------------------------

//isPalindrome collects the id of the url and passes it to the team manager class
//to see if the person's word associated with that id is palindrome.
func (a *api) isPalindrome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	isPalindrome, err := a.teamManager.IsPersonWordPalindrome(vars["ID"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(isPalindrome)
}

//isPalindrome collects the id of the url and passes it to the team manager class
//to see if the person's word associated with that id contains 5 vowels.
func (a *api) hasFiveVowels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	isPalindrome, err := a.teamManager.HasPersonWordFiveVowels(vars["ID"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(isPalindrome)
}
