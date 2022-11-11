package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// // Todo struct
// type Todo struct {
// 	UserID    int    `json:"userId"`
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// // global var
// var todoStruct Todo

func main() {

	var data map[string]any

	req, err := http.NewRequest("GET", "https://quranme-api.vercel.app/quran/surah", nil)
	req.Header.Set("auth-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MmE3NTU5MTAzYzc5YzAyM2I1OTVlYjIiLCJpYXQiOjE2NTUxMzQwOTd9.eF0HAthc6ZNasRXmGcps6oF6neH5jHAfZXGRON51kAg")

	client := &http.Client{}
	res, err := client.Do(req)

	bodyBytes, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyBytes, &data)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	fmt.Println(data["code"])

	// // get
	// fmt.Println("1. Performing Http Get...")
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // post
	// fmt.Println("2. Performing Http Post...")
	// todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	// jsonReq, err := json.Marshal(todo)
	// resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // put
	// fmt.Println("3. Performing Http Put...")
	// todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	// jsonReq, err := json.Marshal(todo)
	// req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 		log.Fatalln(err)
	// }

	// // delete
	// fmt.Println("4. Performing Http Delete...")
	// id := "3"
	// req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/"+id, bytes.NewBuffer(nil))
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer resp.Body.Close()
	// bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// // Convert response body to string
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	// // Convert response body to Todo struct
	// json.Unmarshal(bodyBytes, &todoStruct)
	// fmt.Printf("%+v\n", todoStruct)
}
