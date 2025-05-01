package main

import "fmt"

func main() {

	type User struct {
		Name string
		Age  int
		Tech string
	}

	var users []User

	for i := 0; i < 3; i++ {
		var name string
		fmt.Print("What is your name? ")
		fmt.Scan(&name)
		var age int
		fmt.Print("How old are you? ")
		fmt.Scan(&age)
		var tech string
		fmt.Print("What is your favorite technologie? ")
		fmt.Scan(&tech)

		users = append(users, User{Name: name, Age: age, Tech: tech})
	}

	for _, user := range users {
		fmt.Println(user.Name, "is", user.Age, "years old and their favorite technology is", user.Tech)
	}

}
