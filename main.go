package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

//HomeHandler returns
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	fmt.Println(resp)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var ulist []user
	err = decoder.Decode(&ulist)

	if err != nil {
		fmt.Println(err)
	}

	for _, user := range ulist {
		fmt.Println(user.ID, user.UserID, user.Title+"\t", user.Completed)
	}
}

type peoples struct {
	People  []people
	Message string
	Number  int
}

type people struct {
	Craft string
	Name  string
}

func jsonTest() {
	basicJSONData := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "TAE", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "TAE", "name": "Peggy Whitson"}], "message": "success", "number": 6}`

	jsonToByte := []byte(basicJSONData)

	people := peoples{}

	err := json.Unmarshal(jsonToByte, &people)

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range people.People {
		fmt.Println(v.Craft, v.Name)
	}
}

func main() {
	jsonTest()
	mx := mux.NewRouter()
	mx.HandleFunc("/home", HomeHandler).Methods("POST", "GET")

	http.ListenAndServe(":80", mx)
}
