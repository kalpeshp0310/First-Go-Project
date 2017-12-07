package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/kalpeshp0310/hellogo/models"
)

var users []*m.User

func initData() {
	user1 := setupUserData("Kalpesh")
	user2 := setupUserData("Vagmi")
	users = append(users, user1)
	users = append(users, user2)
}

func setupUserData(name string) *m.User {
	u := m.NewUser(name)
	work := u.NewList("work")

	work.Add("Learn Go").Add("Learn React")

	personal := u.NewList("personal")
	personal.Add("Buy Milk").Add("Buy Sweets").Add("Wash Clothes")

	u.Add(*work).Add(*personal)

	return u.Save()
}
func render404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	w.Write([]byte("{\"message\":\"not found\"}"))
}
func render500(w http.ResponseWriter, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	c, _ := json.Marshal(map[string]string{"message": reason})
	w.Write(c)
}
func renderUser(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		render404(w)
		return
	}

	if id < 1 {
		render404(w)
		return
	}
	u := m.FindUser(id)
	if u == nil {
		render404(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(u)
	if err != nil {
		render500(w, err.Error())
	}
}

func StartServer() {
	// initData()
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", renderUser)

	http.ListenAndServe(":4000", router)
}
