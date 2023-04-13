package internals

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Users = []User{
	{Username: "user1", Email: "test@test.c", FirstName: "John", LastName: "Doe"},
	{Username: "user2", Email: "test2@test.c", FirstName: "Jane", LastName: "Doe"},
}
