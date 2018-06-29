package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}
type MyUser struct {
	Id   int
	Name string
}

type SuperUser struct {
	Id        int
	Name      string
	InnerUser User
}

func main() {
	var user User = User{Id: 1, Name: "王一"}
	userTypeOf := reflect.TypeOf(user)
	userValueOf := reflect.ValueOf(user)
	//	for i := 0; i < userTypeOf.NumField(); i++ {
	//		fmt.Println("fieldname=", userTypeOf.Field(i).Name, ",value=", userValueOf.Field(i))
	//	}
	var myuser MyUser = MyUser{}
	myuserValueOf := reflect.ValueOf(&myuser).Elem()
	fmt.Println("kind=", myuserValueOf.Kind())
	for i := 0; i < userTypeOf.NumField(); i++ {
		fieldName := userTypeOf.Field(i).Name
		fieldValue := userValueOf.Field(i)
		fmt.Println(fieldName, "=", fieldValue)
		myuserValueOf.FieldByName(fieldName).Set(fieldValue)
	}
	fmt.Println(myuser)
}
