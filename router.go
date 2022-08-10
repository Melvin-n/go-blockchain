package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/melvin-n/go-blockchain/models"
)

func Run() error {
	mux := makeMuxRouter()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading env: %s", err.Error())
		return err
	}

	httpAddr := os.Getenv("PORT")
	fmt.Printf("Server listening on port %s...\n", httpAddr)

	s := &http.Server{
		Addr:    "localhost:" + httpAddr,
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("Error running server: %s", err.Error())
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", getBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", writeBlockchain).Methods("POST")

	return muxRouter
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {

	bytes, err := json.MarshalIndent(&Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	io.WriteString(w, string(bytes))
}

func writeBlockchain(w http.ResponseWriter, r *http.Request) {

	var m models.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		//TODO return error with error request body
		return
	}
	newBlock := models.Block{
		Index:     Blockchain[len(Blockchain)-1].Index + 1,
		Timestamp: time.Now(),
		BPM:       m.BPM,
		Hash:      generateHash(newBlock),
		PrevHash:  Blockchain[len(Blockchain)-1].Hash,
	}

}
