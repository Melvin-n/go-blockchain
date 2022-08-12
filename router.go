package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
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
		w.WriteHeader(400)
		err = fmt.Errorf("Unable to decode post body: %s", r.Body)
		fmt.Printf(err.Error())

	}

	newBlock := generateNewBlock(m.BPM, Blockchain[len(Blockchain)-1])

	ok, err := validateBlock(Blockchain[len(Blockchain)-1], newBlock)
	if err != nil {
		w.WriteHeader(400)
		err = fmt.Errorf("Unable to validate block: %s", err.Error())
		fmt.Printf(err.Error())
	}

	if ok {
		Blockchain = append(Blockchain, newBlock)
		refreshChain(Blockchain)
		w.WriteHeader(200)
		spew.Dump(Blockchain)
		json.NewEncoder(w).Encode(Blockchain)
	}
}
