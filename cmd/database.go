package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "email@email.com",
		Id:    "user1",
		Name:  "User",
		Token: "123",
	},
}
