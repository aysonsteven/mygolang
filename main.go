package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Person represents a simple data model.
type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory database
var people = []Person{
	{ID: "1", Name: "John Doe", Age: 30},
	{ID: "2", Name: "Jane Doe", Age: 25},
}

// getPeople handles GET requests to /people
func getPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// getPerson handles GET requests to /people/{id}
func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Path[len("/people/"):]

	for _, person := range people {
		if person.ID == id {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	http.NotFound(w, r)
}

// createPerson handles POST requests to /people
func createPerson(w http.ResponseWriter, r *http.Request) {
	var newPerson Person
	json.NewDecoder(r.Body).Decode(&newPerson)
	people = append(people, newPerson)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPerson)
}

// deletePerson handles DELETE requests to /people/{id}
func deletePerson(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/people/"):]
	for i, person := range people {
		if person.ID == id {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/people", getPeople)        // GET all people
	http.HandleFunc("/people/", getPerson)       // GET, DELETE person by ID
	http.HandleFunc("/people/new", createPerson) // POST a new person

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
