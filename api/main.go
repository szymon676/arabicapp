package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Word struct {
	Arabic  string
	English string
}

func getWords(quantity int) []string {
	rand.Seed(time.Now().UnixNano())

	wordsData := []Word{
		{"سلام", "Peace"},
		{"شمس", "Sun"},
		{"قلب", "Heart"},
		{"ماء", "Water"},
		{"زهرة", "Flower"},
		{"سماء", "Sky"},
		{"كتاب", "Book"},
		{"عائلة", "Family"},
		{"مدينة", "City"},
		{"طعام", "Food"},
		{"جبل", "Mountain"},
		{"بحر", "Sea"},
		{"شجرة", "Tree"},
		{"سمك", "Fish"},
		{"سمير", "Companion in evening talk"},
		{"علم", "Knowledge"},
		{"حلم", "Dream"},
		{"رياحين", "Basil"},
		{"نور", "Light"},
		{"حياة", "Life"},
		{"سفر", "Travel"},
		{"فن", "Art"},
		{"صحة", "Health"},
		{"عمل", "Work"},
		{"نجمة", "Star"},
		{"لغة", "Language"},
		{"فرح", "Joy"},
		{"قهوة", "Coffee"},
		{"موسيقى", "Music"},
		{"قمر", "Moon"},
		{"حضن", "Embrace"},
		{"فضاء", "Space"},
		{"سعادة", "Happiness"},
		{"أمل", "Hope"},
		{"ليل", "Night"},
		{"نهر", "River"},
		{"رمل", "Sand"},
		{"سرور", "Delight"},
		{"قوة", "Strength"},
		{"ضوء", "Light"},
		{"سكون", "Tranquility"},
		{"ضحك", "Laughter"},
		{"فهم", "Understanding"},
		{"سحر", "Magic"},
		{"مدينة", "City"},
		{"غابة", "Forest"},
		{"جدار", "Wall"},
		{"ضياء", "Radiance"},
		{"جسر", "Bridge"},
		{"سمفونية", "Symphony"},
		{"قلادة", "Necklace"},
		{"سفرة", "Table"},
		{"نجاح", "Success"},
		{"زمان", "Time"},
		{"أفق", "Horizon"},
		{"مهندس", "Engineer"},
		{"صداقة", "Friendship"},
		{"حرية", "Freedom"},
		{"ملك", "King"},
		{"ملكة", "Queen"},
		{"مطار", "Airport"},
		{"سوق", "Market"},
		{"جامعة", "University"},
		{"طيارة", "Airplane"},
		{"قصر", "Palace"},
		{"حاسوب", "Computer"},
		{"فستان", "Dress"},
		{"صحراء", "Desert"},
		{"غروب", "Sunset"},
		{"بناء", "Building"},
		{"تصميم", "Design"},
		{"رحلة", "Journey"},
		{"سفير", "Ambassador"},
		{"مشروع", "Project"},
		{"سينما", "Cinema"},
		{"تفاؤل", "Optimism"},
		{"شاعر", "Poet"},
		{"تاريخ", "History"},
		{"مهارة", "Skill"},
		{"فكرة", "Idea"},
		{"حب", "Love"},
	}

	words := []string{}
	for i := 0; i < quantity; i++ {
		randomNumber := rand.Intn(len(wordsData))
		randomWord := wordsData[randomNumber]
		words = append(words, randomWord.Arabic, randomWord.English)
	}

	return words
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(randomWords)
	})

	http.ListenAndServe(":3000", nil)
}
