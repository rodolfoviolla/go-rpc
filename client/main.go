package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	ID int
	Title string
	Body string
}

func main() {
	var reply []Item
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	a := Item{1, "First", "A first item"}
	b := Item{2, "Second", "A second item"}
	c := Item{3, "Third", "A third item"}
	items := []Item{a, b, c}
	client.Call("API.AddItems", items, &reply)
	fmt.Println("Items added:")
	fmt.Println(&reply)
	client.Call("API.EditItem", Item{2, "None", "A none item"}, &reply)
	fmt.Println("Item edited:")
	fmt.Println(&reply[0])
	client.Call("API.GetItemById", 1, &reply)
	fmt.Println("Item got:")
	fmt.Println(&reply[0])
	client.Call("API.DeleteItem", 3, &reply)
	fmt.Println("Item deleted:")
	fmt.Println(&reply[0])
}