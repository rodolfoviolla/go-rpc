package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	ID int
	Title string
	Body string
}

type API int

var database []Item

func (api *API) printDB(cmd string) error {
	fmt.Println("Database after", cmd)
	fmt.Println(database)
	return nil
}

func getById(id int) (index int, item []Item) {
	for currentIndex, value := range database {
		if value.ID == id {
			item = append(item, value)
			index = currentIndex
			break
		}
	}
	return
}

func (api *API) GetItemById(id int, reply *[]Item) error {
	_, *reply = getById(id)
	return api.printDB("GetItemById")
}

func (api *API) AddItems(item []Item, reply *[]Item) error {
	database = append(database, item...)
	*reply = item
	return api.printDB("AddItems")
}

func (api *API) EditItem(edit Item, reply *[]Item) error {
	index, _ := getById(edit.ID)
	database[index] = edit
	*reply = append(*reply, edit)
	return api.printDB("EditItem")
}

func (api *API) DeleteItem(id int, reply *[]Item) error {
	var index int
	index, *reply = getById(id)
	database = append(database[:index], database[index+1:]...)
	return api.printDB("DeleteItem")
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