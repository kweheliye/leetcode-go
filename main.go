package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	//https://www.youtube.com/watch?v=2oXBOaVowYY
	// a := []int{1, 2, 3}
	// b := append(a[:1], 10)

	// fmt.Println(b)
	// fmt.Println(a)
	// message := fmt.Sprintf("Hello %v", a)
	// fmt.Println(message)

	fmt.Printf("##### Welcome to our TodoList App! ######")

	http.HandleFunc("/", showTasks)

	http.ListenAndServe(":8080", nil)

}

func showTasks(writer http.ResponseWriter, reader *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, requet *http.Request) {
	var greeting = "Hello user. Welcome to our TodoList App!"
	fmt.Fprint(writer, greeting)

}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	return append(taskItems, newTask)
}
