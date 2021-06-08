package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type DisposalEvent struct {
	Type     string    `json:"type"`
	Nextdate time.Time `json:"nextDate"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/v1/get-disposal-events", GetDisposalEvents)
	r.Get("/v1/get-disposal-event/{id}", GetDisposalEvent)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Serving http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func GetDisposalEvents(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://him.ffail.win/")
	if err != nil {
		http.Error(res, "Error on HTTP/GET/https://him.ffail.win/: "+err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var disposalEvents []DisposalEvent
	err = json.NewDecoder(resp.Body).Decode(&disposalEvents)
	if err != nil {
		http.Error(res, "Error on JSON decoding: "+err.Error(), http.StatusInternalServerError)
	}

	for i := range disposalEvents {
		disposalEvents[i].Type = disposalEvents[i].Type + "Marius"
	}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(disposalEvents)
	if err != nil {
		http.Error(res, "Error on JSON encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetDisposalEvent(res http.ResponseWriter, req *http.Request) {
	disposalEventId, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(res, "Error on converting URL param 'id' from string to int: "+err.Error(), http.StatusInternalServerError)
	}

	resp, err := http.Get("https://him.ffail.win/")
	if err != nil {
		http.Error(res, "Error on HTTP/GET/https://him.ffail.win/: "+err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var disposalEvents []DisposalEvent
	err = json.NewDecoder(resp.Body).Decode(&disposalEvents)
	if err != nil {
		http.Error(res, "Error on JSON decoding: "+err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(disposalEvents[disposalEventId])
	if err != nil {
		http.Error(res, "Error on JSON encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
