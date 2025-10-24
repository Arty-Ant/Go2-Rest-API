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
	"fmt"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

type Response struct {
	Result string `json:"result"`
}

var (
	a, b int
	port = "1234"
	)

func GetAllInfo(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(msg)
}

func GetFirst(writer http.ResponseWriter, request *http.Request) {
	a:=rand.Intn(100)
	json.NewEncoder(writer).Encode(a)
}

func GetSecond(writer http.ResponseWriter, request *http.Request) {
	b:=rand.Intn(100)
	json.NewEncoder(writer).Encode(b)
}

func GetAdd(writer http.ResponseWriter, request *http.Request) {
	a:=rand.Intn(100)
	b:=rand.Intn(100)
	result := a+b
	json.NewEncoder(writer).Encode(result)
}

func GetMul(writer http.ResponseWriter, request *http.Request) {
	a:=rand.Intn(100)
	b:=rand.Intn(100)
	result := a*b
	json.NewEncoder(writer).Encode(result)
}

func GetDiv(writer http.ResponseWriter, request *http.Request) {
	a:=rand.Intn(100)
	b:=rand.Intn(100)
	result := a/b
	json.NewEncoder(writer).Encode(result)
}

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(10000))
	router := mux.NewRouter()
	router.HandleFunc("/info", GetAllInfo).Method("GET")
	router.HandleFunc("/first", GetFirst).Method("GET")
	router.HandleFunc("/second", GetSecond).Method("GET")
	router.HandleFunc("/add", GetAdd).Method("GET")
	router.HandleFunc("/mul", GetMul).Method("GET")
	router.HandleFunc("/div", GetDiv).Method("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
