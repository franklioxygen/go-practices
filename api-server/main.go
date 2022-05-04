package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Users []User

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_Name"`
	LastName   string `json:"last_Name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

var usersInstance Users

func check(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func recoverError(err error) {
	if err := recover(); err != nil {
		fmt.Println("Error:", err)
	}
}

func readDatabase() {
	// mock data
	dataFile, err := os.Open("./database.json")
	check(err)
	defer dataFile.Close()
	byteValue, _ := ioutil.ReadAll(dataFile)
	json.Unmarshal(byteValue, &usersInstance)
}

// Example: visit http://localhost:8080/
func handler(w http.ResponseWriter, r *http.Request) {
	genJson, _ := json.Marshal(usersInstance)
	fmt.Fprintf(w, "%v", string(genJson))
}

// Example: visit http://localhost:8080/id/1
func idHandler(w http.ResponseWriter, r *http.Request) {
	dir := strings.Split(r.URL.Path, "/")
	for _, user := range usersInstance {
		if strconv.Itoa(user.Id) == dir[2] {
			response, _ := json.Marshal(user)
			fmt.Fprintf(w, "%v", string(response))
			break
		}
	}
}

// Example: visit http://localhost:8080/dept/Engineering
func departmentHandler(w http.ResponseWriter, r *http.Request) {
	dir := strings.Split(r.URL.Path, "/")
	response := []string{}
	for _, user := range usersInstance {
		if user.Department == dir[2] {
			genJson, _ := json.Marshal(user)
			response = append(response, string(genJson)+",")
		}
	}
	response[len(response)-1] = strings.TrimSuffix(response[len(response)-1], ",")
	fmt.Fprintf(w, "%v", response)
}

func main() {
	readDatabase()
	http.HandleFunc("/", handler)
	http.HandleFunc("/id/", idHandler)
	http.HandleFunc("/dept/", departmentHandler)
	err := (http.ListenAndServe(":8080", nil))
	defer recoverError(err)
	check(err)
	fmt.Println("Server started at localhost:8080\n ")
}
