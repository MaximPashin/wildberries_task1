package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"wb_task1/db"
	"wb_task1/entity"
	"wb_task1/storage"

	stan "github.com/nats-io/stan.go"
)

var s *storage.Storage

func main() {
	connURL := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	conn, err := db.New(connURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	s, err = storage.New(conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.Recovery()

	natsChan := make(chan []byte, 10)
	sub(natsChan)

	go func() {
		for {
			data := <-natsChan
			entity := entity.Order{}
			err = json.Unmarshal([]byte(data), &entity)
			if err != nil {
				log.Print(err.Error())
			}
			err = s.Put(entity)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

	http.HandleFunc("/get/", get)
	http.ListenAndServe(":8080", nil)
}

func sub(eventsChan chan<- []byte) {
	sc, err := stan.Connect("nats-streaming", "my_sub")
	if err != nil {
		log.Println(err.Error())
		return
	}
	sc.Subscribe("foo", func(m *stan.Msg) {
		eventsChan <- m.Data
	})
}

func get(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/get/")
	log.Println(id)
	order, ok := s.Get(id)
	if ok {
		data, err := json.Marshal(order)
		if err != nil {
			log.Println(err.Error())
		}
		_, err = w.Write(data)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
	}
}
