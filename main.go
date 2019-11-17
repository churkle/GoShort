package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

var letterSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var pool *redis.Pool

type urlStruct struct {
	URL string
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	url, err := getURL(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, url)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var jsonURL urlStruct
	err := decoder.Decode(&jsonURL)

	if err != nil {
		log.Fatal(err)
	}

	id, err := shortenURL(jsonURL.URL)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	setupRouter()
}

func setupRouter() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{id}", get).Methods(http.MethodGet)
	api.HandleFunc("/shorten", post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createIdentifier(length int) string {
	id := make([]rune, length)
	for i := range id {
		id[i] = letterSet[rand.Intn(len(letterSet))]
	}

	return string(id)
}

func shortenURL(url string) (string, error) {
	id := createIdentifier(5)
	err := writeShortenedURLToDb(url, id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func writeShortenedURLToDb(url string, id string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSET", id, "URL", url)

	return err
}

func getURL(id string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	url, err := redis.String(conn.Do("HGET", id, "URL"))
	if err != nil {
		return "", err
	}

	return url, err
}
