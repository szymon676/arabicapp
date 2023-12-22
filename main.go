package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Word struct {
	Arabic  string `json:"arabic"`
	English string `json:"english"`
}

type Response struct {
	Words []Word `json:"words"`
}

func readWordsFromFile() []Word {
	file, err := ioutil.ReadFile("arabicwords.json")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	var words []Word
	if err := json.Unmarshal(file, &words); err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	return words
}

func getWords(quantity int) Response {
	rand.Seed(time.Now().UnixNano())

	wordsData := readWordsFromFile()

	words := []Word{}
	for i := 0; i < quantity; i++ {
		randomNumber := rand.Intn(len(wordsData))
		randomWord := wordsData[randomNumber]
		words = append(words, randomWord)
	}

	return Response{Words: words}
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	http.HandleFunc("/api/words", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			enableCORS(w)
			w.WriteHeader(http.StatusOK)
			return
		}

		quantity := r.URL.Query().Get("quantity")

		q, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal("bad query")
		}

		randomWords := getWords(q)

		enableCORS(w)
		WriteAPI(w, 200, randomWords)
	})

	fmt.Println("starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}

func WriteAPI(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(v)
}
