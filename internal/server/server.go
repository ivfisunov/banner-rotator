package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ivfisunov/banner-rotator/internal/app"
	"github.com/ivfisunov/banner-rotator/internal/storage/stortypes"
)

type Server struct {
	*app.App
	server *http.Server
	router *mux.Router
}

func NewServer(app *app.App, host string, port string) *Server {
	router := mux.NewRouter()

	addr := net.JoinHostPort(host, port)
	server := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return &Server{app, server, router}
}

func (s *Server) Start(ctx context.Context) error {
	s.router.Use(s.loggingMiddleware)
	s.router.HandleFunc("/hello", s.helloHandler).Methods("GET")
	s.router.HandleFunc("/add-banner", s.addBanner).Methods("POST")
	s.router.HandleFunc("/delete-banner", s.deleteBanner).Methods("POST")
	s.router.HandleFunc("/add-click", s.addClick).Methods("POST")
	s.router.HandleFunc("/display-banner", s.dispalyBanner).Methods("POST")

	s.Logger.Info(fmt.Sprintf("HTTP server started on %s", s.server.Addr))
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.Logger.Info("HTTP server is stoped")
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "hello"})
	if err != nil {
		s.Logger.Error("error sending response")
	}
}

func (s *Server) addBanner(w http.ResponseWriter, r *http.Request) {
	data := stortypes.BunnerBody{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		s.Logger.Error("body parser error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// err = s.Storage.AddBanner(data.BannerID, data.slotID)
	if err != nil {
		s.Logger.Error("error inserting banner to slot: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "banner added successfully"})
	if err != nil {
		s.Logger.Error("error sending response")
	}
}

func (s *Server) deleteBanner(w http.ResponseWriter, r *http.Request) {
	data := stortypes.BunnerBody{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		s.Logger.Error("body parser error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// err = s.Storage.DeleteBanner(data.BannerID, data.slotID)
	if err != nil {
		s.Logger.Error("error deleting banner from slot: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "banner deleted successfully"})
	if err != nil {
		s.Logger.Error("error sending response")
	}
}

func (s *Server) addClick(w http.ResponseWriter, r *http.Request) {
	data := stortypes.AddClickBody{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		s.Logger.Error("body parser error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// err = s.Storage.AddClick(data.BannerID, data.SlotID. data.GroupID)
	if err != nil {
		s.Logger.Error("click addition error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "click added successfully"})
	if err != nil {
		s.Logger.Error("error sending response")
	}
}

func (s *Server) dispalyBanner(w http.ResponseWriter, r *http.Request) {
	data := stortypes.DisplayBannerBody{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		s.Logger.Error("body parser error: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// bannerToDisplay, err =: s.Storage.DisplayBanner(data.SlotID. data.GroupID)
	if err != nil {
		s.Logger.Error("getting banner error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"id": "", "description": ""})
	if err != nil {
		s.Logger.Error("error sending response")
	}
}
