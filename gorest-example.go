package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contact", GetPeople).Methods("GET")
    router.HandleFunc("/contact/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contact/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/contact/{id}", DeletePerson).Methods("DELETE")
    
    //Feed Person Array
    people = append(people, Person{ID: "1", DESCRIPTION: "Pessoa1"})
    people = append(people, Person{ID: "2", DESCRIPTION: "Pessoa2"})
    people = append(people, Person{ID: "3", DESCRIPTION: "Pessoa3"})
    people = append(people, Person{ID: "4", DESCRIPTION: "Pessoa4"})
    people = append(people, Person{ID: "5", DESCRIPTION: "Pessoa5"})
    people = append(people, Person{ID: "6", DESCRIPTION: "Pessoa6"})
	
    log.Fatal(http.ListenAndServe(":8000", router))
}

type Person struct{
    ID string `json:"id,omitempty"`
    DESCRIPTION string `json:"description,omitempty"`
}
var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people{
        if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
            return
        }        
    }
    json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    json.NewEncoder(w).Encode(people)
    
    }
}