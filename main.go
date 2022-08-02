package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"http-server/users"
)

func initLog(f *os.File) {
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	logfile, err := os.OpenFile("server.log", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logfile.Close()
	initLog(logfile)
	r := mux.NewRouter()
	r.HandleFunc("/", testHandler).Methods("GET")
	r.HandleFunc("/users", listUsers).Methods("GET")
	r.HandleFunc("/user", addUser).Methods("POST")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	portNumber := os.Getenv("PORT_NUM")
	if portNumber == "" {
		portNumber = "8081"
	}
	log.Debugln("Starting Server on Port ", portNumber)
	fmt.Println("Starting Server on Port", portNumber)
	log.Fatal(http.ListenAndServe(":"+portNumber, r))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugln("Request hit to fetch users list")
	log.Debugln("Users :", users.Users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Server is Working")

}

func listUsers(w http.ResponseWriter, r *http.Request) {
	log.Debugln("Request hit to fetch users list")
	log.Debugln("Users :", users.Users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users.Users)

}

func addUser(w http.ResponseWriter, r *http.Request) {
	log.Debugln("Request hit to create a user")
	w.Header().Set("Content-Type", "application/json")
	var u users.User
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = uuid.NewString()
	users.Users = append(users.Users, u)
	log.Debugln("User :", u)
	json.NewEncoder(w).Encode(users.Users)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	log.Debugln("Request hit to delete a user")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for i, u := range users.Users {
		if u.Id == id {
			log.Debugln("User :", u)
			users.Users = append(users.Users[:i], users.Users[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("User deleted successfully")

}
