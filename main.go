package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Consuming API Endpoints ->
var urls = []string{
	"https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b",
	"https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c",
	"https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148",
}

//Item Struct ->
type Item struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
	Quantity	int	`json:"quantity"`
	Price	int	`json:"price"`
}

//Init Items var as a slice Item struct ->
var Items []Item

//Home Page ->
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><strong>Food Aggregator</strong></h1>"))
}

//Get Items by Name ->
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	 flag := 0
	 for _, web := range urls {
		 res, err := http.Get(web)

		 if err != nil {
			 panic(err)
		 }

		 dataBytes, err := ioutil.ReadAll(res.Body)
		 if err != nil {
			 panic(err)
		 }

		 data := dataBytes
		 json.Unmarshal(data, &Items)

		 params := mux.Vars(r)

		 for _, item := range Items {

			if item.Name == params["name"] {
				flag = 1
				json.NewEncoder(w).Encode(item)
				return
			}
		 }
	 }

	 if flag == 0 {
		 json.NewEncoder(w).Encode("NOT FOUND")
	 }
}


//Get Items by Name & Quantity ->
func getItemsByQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	 flag := 0
	 for _, web := range urls {
		 res, err := http.Get(web)

		 if err != nil {
			 panic(err)
		 }

		 dataBytes, err := ioutil.ReadAll(res.Body)
		 if err != nil {
			 panic(err)
		 }

		 data := dataBytes
		 json.Unmarshal(data, &Items)

		 params := mux.Vars(r)
		 qty, _ := strconv.Atoi(params["quantity"])

		 for _, item := range Items {
			 fmt.Println(item)
			 if item.Name == params["name"] && item.Quantity >= qty {
				 flag = 1
				 json.NewEncoder(w).Encode(item)
			 }
		 }
	}

	if flag == 0 {
		json.NewEncoder(w).Encode("NOT FOUND")
	}
}

//Get Items by Name , Quantity & Price ->
func getItemsByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	 flag := 0
	 for _, web := range urls {
		 res, err := http.Get(web)

		 if err != nil {
			 panic(err)
		 }

		 dataBytes, err := ioutil.ReadAll(res.Body)
		 if err != nil {
			 panic(err)
		 }

		 data := dataBytes
		 json.Unmarshal(data, &Items)

		 params := mux.Vars(r)
		 price, _ := strconv.Atoi(params["price"])

		 for _, item := range Items {
			 fmt.Println(item)
			 if item.Name == params["name"] && item.Price >= price {
				 flag = 1
				 json.NewEncoder(w).Encode(item)
			 }
		 }
	}

	if flag == 0 {
		json.NewEncoder(w).Encode("NOT FOUND")
	}
}

func main() {
	//Init Router ->
	r := mux.NewRouter()

	//Route Handlers - Endpoints ->
	r.HandleFunc("/api", Home).Methods("GET")
	r.HandleFunc("/api/buy-item/{name}", getItems).Methods("GET")
	r.HandleFunc("/api/buy-item-qty/{name}&{quantity}", getItemsByQty).Methods("GET")
	r.HandleFunc("/api/buy-item-qty-price/{name}&{quantity}&{price}", getItemsByPrice).Methods("GET")

	fmt.Println("Server started on port:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}