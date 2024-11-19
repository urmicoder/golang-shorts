package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

/*
StatusOK                     = 200
StatusCreated                = 201
StatusAccepted               = 202
StatusFound                  = 302
StatusBadRequest             = 400
StatusNotFound               = 404
StatusNotAcceptable          = 406
StatusInternalServerError    = 500
StatusBadGateway             = 502
*/
// Domain Layer
type RoutingTable struct {
	Id   int    `json:"id" bson:"_id"` // These are the Tags, it basically used for Decoding or Encoding of Json/ Bson/ XML FROM or TO Struct Variables/Object.
	Name string `json:"name" bson:"name"`
	IP   string `json:"ip" bson:"ip"`
}

var rts []RoutingTable // Creating a slice, its basically act as a in-memory Database.
// Transport
type AddRes struct {
	Success bool `json:"success"`
}
type ListRes struct {
	Success bool           `json:"success"`
	Data    []RoutingTable `json:"data"`
}
type GetRes struct {
	Success bool         `json:"success"`
	Data    RoutingTable `json:"data"`
}
type UpdateRes struct {
	Success bool `json:"success"`
}
type DeleteRes struct {
	Success bool `json:"success"`
}

// -------------------------------------------------------------

func Add(w http.ResponseWriter, r *http.Request) {

	var req RoutingTable
	// reads from the request's body (r.Body)
	// decode the JSON data into the variable/object req
	// See here NewDecoder needs io.Reader.
	// r.Body implements io.Reader and io.Closer both by Composing Interfaces together.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(AddRes{Success: false})
		return
	}
	if len(req.Name) < 1 || len(req.IP) < 1 { // checks whether the length of the Name or IP fields in the req are empty or missing.
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(AddRes{Success: false})
		return
	}
	req.Id = int(time.Now().UnixNano()) // Auto Gen ID by Time representaion of INT64
	rts = append(rts, req)
	w.WriteHeader(http.StatusCreated) // sets the HTTP status code to 201,
	//  creates a new JSON encoder that writes to the HTTP response writer w
	json.NewEncoder(w).Encode(AddRes{Success: true})
}

// -------------------------------------------------------------

func List(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // sets the HTTP status code to 200,
	// tells the client (like a web browser or PostMan) that the response body contains JSON(can be also binary-file/XML) data. This is important for proper handling and parsing of the response by the client.
	w.Header().Set("Content-Type", "application/json")
	//  creates a new JSON encoder that writes to the HTTP response writer w
	json.NewEncoder(w).Encode(ListRes{Success: true, Data: rts})
}

// -------------------------------------------------------------

func Get(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id") // retrieves the value of the URL query parameter named "id" from the incoming HTTP request.
	idint, _ := strconv.Atoi(id)  // converts the string value stored in id to an integer using the strconv.Atoi function, which returns the integer and an error (ignored with _ here).
	for _, rt := range rts {
		if idint == rt.Id {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(GetRes{Success: true, Data: rt})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetRes{Success: false, Data: RoutingTable{}})
}

// -------------------------------------------------------------

func Update(w http.ResponseWriter, r *http.Request) {
	var req RoutingTable
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.Name) < 1 || len(req.IP) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(UpdateRes{Success: false})
		return
	}

	for i, rt := range rts {
		if rt.Id == req.Id { // Match the ID
			rts[i] = req // Assign Updated Values
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(UpdateRes{Success: true})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(UpdateRes{Success: false})
}

// -------------------------------------------------------------

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)

	for i, rt := range rts {
		if rt.Id == idint {
			rts = append(rts[:i], rts[i+1:]...) // Remove the entry
			/*
			   rts[:i]: This gives you a new slice that includes all elements from the start of the slice up to (but not including) the element at index i. It effectively represents the first part of the original slice.

			   rts[i+1:]: This gives you a new slice that starts from the element at index i+1 to the end of the original slice. It effectively represents the remaining part of the original slice after the element at index i.
			*/
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(DeleteRes{Success: true})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(DeleteRes{Success: false})
}

// -------------------------------------------------------------

func main() {
	// Routing refers to the process of defining how an application responds to client requests for specific endpoints (URLs).

	// There are 2 types of Routing, one is Static Routing (fixed URL paths) and another is Dynamic Routing (variable URL paths).

	// Static Routing Examples: /rtable/api/add , /rtable/api/get

	// The route /rtable/{ID} captures dynamic parts of the URL
	// Dynamic Routing Examples: /rtable/123 , /rtable/124

	// For Dynamic Routing we can use the Split function which split the URL by Slash and then we get the ID part from the Slice. You can extract parameters from the URL using r.URL.Path
	// pathParts := strings.Split(r.URL.Path, "/")
	// rTableId := pathParts[2]

	// Any Function which has the Signature of AnyFunc(http.ResponseWrite, *http.Request) - then it can be passes as arguments of HandleFunc() function
	http.HandleFunc("/rtable/api/add", MethodCheck(http.MethodPost, Add))
	http.HandleFunc("/rtable/api/get", MethodCheck(http.MethodGet, Get))
	http.HandleFunc("/rtable/api/list", MethodCheck(http.MethodGet, List))
	http.HandleFunc("/rtable/api/update", MethodCheck(http.MethodPut, Update))
	http.HandleFunc("/rtable/api/delete", MethodCheck(http.MethodDelete, Delete))

	if err := http.ListenAndServe(":8080", nil); err != nil { // This sets up a simple HTTP server listening on port 8080.
		panic(err)
	}
}

// -------------------------------------------------------------

// Middleware can inspect, modify, or reject incoming requests before they reach the final handler
func MethodCheck(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(AddRes{Success: false})
			return
		}
		next(w, r) // this will call the function passed as arguments.
	}
}
