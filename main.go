package main

import "fmt"

type Item struct {
	title string
	body string
}

var database []Item

func GetByName(title string) (index int, item Item) {
	for currentIndex, value := range database {
		if value.title == title {
			item = value
			index = currentIndex
			break
		}
	}
	return
}

func AddItems(item []Item) []Item {
	database = append(database, item...)
	return item
}

func EditItem(title string, edit Item) Item {
	index, _ := GetByName(title)
	database[index] = edit
	return edit
}

func DeleteItem(title string) (item Item) {
	index, item := GetByName(title)
	database = append(database[:index], database[index+1:]...)
	return 
}

func main() {
	fmt.Println("Initial database:", database)
	a := Item{"first", "a test item"}
	b := Item{"second", "a second test item"}
	c := Item{"third", "a third test item"}
	AddItems([]Item{a, b, c})
	fmt.Println("Second database:", database)
	DeleteItem(b.title)
	fmt.Println("Third database:", database)
	EditItem("third", Item{"second", "a new second test item"})
	fmt.Println("Fourth database:", database)
	_, x := GetByName("second")
	_, y := GetByName("first")
	fmt.Println(x, y)
}