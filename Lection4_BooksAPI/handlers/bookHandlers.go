package handlers

import (
	"BooksAPI/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use parament ID as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
	} else {
		writer.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(writer).Encode(book)
	}
}