package main

import (
	"time"
)

// Здесь надо реализовать структуру User
type User struct {
	Name string
	Login string
	Password string
	LastSeenAt time.Time
}

// Здесь надо реализовать структуру Admin, включающую в себя User
type Admin struct {
	User
	Access uint8
}

// Здесь должна быть функция изменения имени пользователя
func renameUser(user interface{}, newName string) {
	switch v := user.(type){
	case *User:
		v.Name = newName
	case *Admin:
		if v.Access < 3{
			v.Name = newName
		}
	}
}

func main() {

}
