package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Users []User `json:"app_users"`
}

type User struct {
	Id         int    `json:"id"`
	First_Name string `json:"first_Name"`
	Last_Name  string `json:"last_Name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

var usersInstance Users

func check(e error) {
	if e != nil {
		panic(e)
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
	for _, user := range usersInstance.Users {
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
	for _, user := range usersInstance.Users {
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
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Server started at localhost:8080\n ")
}
