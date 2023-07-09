package models

import "math/rand"
import "strconv"

type User struct {
	ID string `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age string `json:"age"`
}

type BodyData struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age string `json:"age"`
}

var users = []User{
	{
		ID:"1",
		Fname:"Khalid",
		Lname: "Mujahid",
		Age: "23",
	},
	{
		ID:"2",
		Fname:"Muhammad",
		Lname: "Awwal",
		Age: "43",
	},
}

// fetch all users
func GetUsers() []User {
	return users
}

// fetch single user
func GetUser(id string) User {
	for _,user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

// add new user
func AddUser(user BodyData) User {
	// generate a random id
	id := strconv.Itoa(rand.Intn(1000))
	
	users = append(users,User{ ID: id, Fname: user.Fname, Lname: user.Lname, Age: user.Age})
	return User{ ID: id, Fname: user.Fname, Lname: user.Lname, Age: user.Age}
}

// remove user
func RemoveUser(id string) string {
	for index,user := range users {
		if user.ID == id {
			users = append(users[:index],users[index+1:]...)
			break
		}
	}
	return "User with this ID " + id + " is removed"
}
