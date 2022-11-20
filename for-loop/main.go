package main

import "log"

func main() {
	// animals := []string{"cat", "dog", "fish"}

	// for idx, val := range animals {
	// 	log.Println("idx =", idx, ",val=", val)
	// }

	// animals := make(map[string]string)
	// animals["cat"] = "Mary"
	// animals["dog"] = "Bob"

	// for key, val := range animals {
	// 	log.Println("key =", key, ",val=", val)
	// }

	// var firstLine = "Long long time ago..."
	// // a string in go is a slice of bytes

	// for i, v := range firstLine {
	// 	log.Println(i, string(v))
	// }

	type User struct {
		Name string
		Age  int
	}

	var users []User
	users = append(users, User{"John", 18})
	users = append(users, User{"Alice", 36})
	users = append(users, User{"Bob", 25})

	for i, v := range users {
		// log.Println(i, ":", v)
		log.Println(i, ":", v.Name)
	}
}
