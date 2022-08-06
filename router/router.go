package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	muxRouter.HandleFunc(getBlockchain).Methods("GET")
	muxRouter.HandleFunc(writeBlockchain).Methods("POST")

	return muxRouter
}

func getBlockchain(w http.ResponseWriter, r *http.Request) {
	//TODO: handle get request for blockchain
}
