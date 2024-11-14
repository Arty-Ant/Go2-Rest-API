package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type Response struct {
	Result string `json:"result"`
}

var (
	port         = "1234"
	firstNumber  int
	secondNumber int
)

func main() {
	// Маршруты
	mux := http.NewServeMux()

	mux.HandleFunc("GET /info", infoHandler)
	mux.HandleFunc("GET /first", firstHandler)
	mux.HandleFunc("GET /second", secondHandler)
	mux.HandleFunc("GET /add", addHandler)
	mux.HandleFunc("GET /sub", subHandler)
	mux.HandleFunc("GET /mul", mulHandler)
	mux.HandleFunc("GET /div", divHandler)

	log.Println("Start server")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// info
func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Result: "Калькулятор API",
	}
	json.NewEncoder(w).Encode(response)
}

// Генерация случайного числа
func generateRandomNumber() int {
	return rand.Intn(100)
}

// first
func firstHandler(w http.ResponseWriter, r *http.Request) {
	firstNumber = generateRandomNumber()
	response := Response{
		Result: fmt.Sprintf("%d", firstNumber),
	}
	json.NewEncoder(w).Encode(response)
}

// second
func secondHandler(w http.ResponseWriter, r *http.Request) {
	secondNumber = generateRandomNumber()
	response := Response{
		Result: fmt.Sprintf("%d", secondNumber),
	}
	json.NewEncoder(w).Encode(response)
}

// add
func addHandler(w http.ResponseWriter, r *http.Request) {
	result := firstNumber + secondNumber
	response := Response{
		Result: fmt.Sprintf("%d", result),
	}
	json.NewEncoder(w).Encode(response)
}

// sub
func subHandler(w http.ResponseWriter, r *http.Request) {
	result := firstNumber - secondNumber
	response := Response{
		Result: fmt.Sprintf("%d", result),
	}
	json.NewEncoder(w).Encode(response)
}

// mul
func mulHandler(w http.ResponseWriter, r *http.Request) {
	result := firstNumber * secondNumber
	response := Response{
		Result: fmt.Sprintf("%d", result),
	}
	json.NewEncoder(w).Encode(response)
}

// div
func divHandler(w http.ResponseWriter, r *http.Request) {
	if secondNumber == 0 {
		response := Response{
			Result: "Ошибка деления на 0",
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	result := float64(firstNumber) / float64(secondNumber)
	response := Response{
		Result: fmt.Sprintf("%.2f", result),
	}
	json.NewEncoder(w).Encode(response)
}
