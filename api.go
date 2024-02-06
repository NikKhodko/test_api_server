package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

		}
	}
}

type APIserver struct {
	listenAdd string
}

func NewAPIServer(listenAdd string) *APIserver {
	return &APIserver{
		listenAdd: listenAdd,
	}
}

func (s *APIserver) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount())
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleTransaction(w http.ResponseWriter, r *http.Request) error {
	return nil
}
