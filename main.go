package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	id int
	title string
	body string
}

type API int

var database []Item

func getById(id int) (index int, item Item) {
	for currentIndex, value := range database {
		if value.id == id {
			item = value
			index = currentIndex
			break
		}
	}
	return
}

func (api *API) GetItemById(id int, reply *Item) error {
	_, *reply = getById(id)
	return nil
}

func (api *API) AddItems(item []Item, reply *[]Item) error {
	database = append(database, item...)
	*reply = item
	return nil
}

func (api *API) EditItem(edit Item, reply *Item) error {
	index, _ := getById(edit.id)
	database[index] = edit
	*reply = edit
	return nil
}

func DeleteItem(id int, reply *Item) error {
	var index int
	index, *reply = getById(id)
	database = append(database[:index], database[index+1:]...)
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registering API", err)
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error creating listener", err)
	}
	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}