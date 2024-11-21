package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

//Domain..............

type Counter struct {
	Count int
}

//Service.................

type Service interface {
	AddCounter()
	GetCounter(id string) int
}

type service struct {
	Counter Counter
}

func NewService(c Counter) *service {
	return &service{Counter: c}
}

func (s *service) AddCounter() {
	s.Counter.Count++
}

func (s *service) GetCounter(Id string) int {
	fmt.Println(Id)
	return s.Counter.Count
}

//Transport....................

type httpTransport struct {
	Service
}

func NewHttpTransport(s Service) *httpTransport {
	return &httpTransport{Service: s}
}

type AddReq struct{}
type AddRes struct {
	Success bool `json:"success"`
}

type GetReq struct {
	Id string `json:"_id"`
}

type GetRes struct {
	Count   int  `json:"count"`
	Success bool `json:"success"`
}

type GetUrl struct {
	Id string `url:"_id"`
}

func (t *httpTransport) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t.AddCounter()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AddRes{Success: true})
}

func (t *httpTransport) Get(w http.ResponseWriter, r *http.Request) {
	var req GetReq
	json.NewDecoder(r.Body).Decode(&req)
	count := t.GetCounter(req.Id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}

func (t *httpTransport) GetUrl(w http.ResponseWriter, r *http.Request) {
	u, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		json.NewEncoder(w).Encode(GetRes{Count: -1, Success: false})
	}
	id := u["_id"][0]
	count := t.GetCounter(id)
	json.NewEncoder(w).Encode(GetRes{Count: count, Success: true})
}

func main() {
	s := NewService(Counter{})

	t := NewHttpTransport(s)
	//Routing
	r := mux.NewRouter()
	r.HandleFunc("/api/add", t.Add).Methods("POST")
	r.HandleFunc("/api/get", t.Get).Methods("GET")
	r.HandleFunc("/api/get/url", t.GetUrl).Methods("GET")
	fmt.Println("i am router")
	//Server
	http.ListenAndServe(":3000", r)
}
