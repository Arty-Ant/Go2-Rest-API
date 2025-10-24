/*
## Задача № 1 (Calculator API)
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление

*результат вернуть в виде JSON

import "math/rand"

number := rand.Intn(100)

// Queries
GET http://127.0.0.1:1234/info

GET http://127.0.0.1:1234/first

GET http://127.0.0.1:1234/second

GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
*/
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Info struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	Port      string   `json:"port"`
	Endpoints []string `json:"endpoints"`
}

type Response struct {
	Result int `json:"result"`
}

var (
	a, b int
	port = "1234"
)

func GetAllInfo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	info := Info{
		Name:    "Calculator API",
		Version: "1.0",
		Port:    port,
		Endpoints: []string{
			"/info",
			"/first",
			"/second",
			"/add",
			"/sub",
			"/mul",
			"/div",
		},
	}

	json.NewEncoder(writer).Encode(info)
}

func GetFirst(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	a = rand.Intn(100)
	json.NewEncoder(writer).Encode(Response{Result: a})
}

func GetSecond(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	b = rand.Intn(100)
	json.NewEncoder(writer).Encode(Response{Result: b})
}

func GetAdd(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	a = rand.Intn(100)
	b = rand.Intn(100)
	result := a + b
	json.NewEncoder(writer).Encode(Response{Result: result})
}

func GetSub(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	a = rand.Intn(100)
	b = rand.Intn(100)
	result := a - b
	json.NewEncoder(writer).Encode(Response{Result: result})
}

func GetMul(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	a = rand.Intn(100)
	b = rand.Intn(100)
	result := a * b
	json.NewEncoder(writer).Encode(Response{Result: result})
}

func GetDiv(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	a = rand.Intn(100)
	b = rand.Intn(100)
	if b == 0 {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result := a / b
	json.NewEncoder(writer).Encode(Response{Result: result})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info", GetAllInfo).Methods("GET")
	router.HandleFunc("/first", GetFirst).Methods("GET")
	router.HandleFunc("/second", GetSecond).Methods("GET")
	router.HandleFunc("/add", GetAdd).Methods("GET")
	router.HandleFunc("/sub", GetSub).Methods("GET")
	router.HandleFunc("/mul", GetMul).Methods("GET")
	router.HandleFunc("/div", GetDiv).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
