package user

var users = []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

func GetAll() []User {
    return users
}
