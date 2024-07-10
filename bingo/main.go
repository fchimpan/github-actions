package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type BingoCard struct {
	Columns map[string][]int
	Marks   map[string]map[int]bool
}

var (
	card  BingoCard
	mutex sync.Mutex
)

func generateBingoCard() BingoCard {
	rand.Seed(time.Now().UnixNano())
	columns := map[string][]int{
		"B": generateColumn(1, 15),
		"I": generateColumn(16, 30),
		"N": generateColumn(31, 45),
		"G": generateColumn(46, 60),
		"O": generateColumn(61, 75),
	}
	columns["N"][2] = 0 // Free Space

	marks := map[string]map[int]bool{
		"B": {},
		"I": {},
		"N": {},
		"G": {},
		"O": {},
	}
	return BingoCard{Columns: columns, Marks: marks}
}

func generateColumn(start, end int) []int {
	nums := rand.Perm(end - start + 1)[:5]
	for i := range nums {
		nums[i] += start
	}
	return nums
}

func main() {
	card = generateBingoCard()
	tmpl := template.Must(template.ParseFiles("bingo.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		tmpl.Execute(w, card)
	})

	http.HandleFunc("/mark", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		column := r.URL.Query().Get("column")
		indexStr := r.URL.Query().Get("index")

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		if card.Marks[column] == nil {
			card.Marks[column] = map[int]bool{}
		}
		card.Marks[column][index] = !card.Marks[column][index]

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
