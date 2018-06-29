package main

import (
	"errors"
	"fmt"
	"reflect"
)

type User struct {
	Id        int
	Name      string
	Addresses []Address
}
type Address struct {
	Id       int
	Location string
}

type UserInfo struct {
	User
	Desc string
}

// https://gocn.io/question/99
func Copy(dst interface{}, src interface{}) (err error) {
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		err = errors.New("dst isn't a pointer to struct")
		return
	}
	dstElem := dstValue.Elem()
	if dstElem.Kind() != reflect.Struct {
		err = errors.New("pointer doesn't point to struct")
		return
	}

	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		err = errors.New("src isn't struct")
		return
	}

	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		sv := srcValue.FieldByName(sf.Name)
		// make sure the value which in dst is valid and can set
		if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
			dv.Set(sv)
		}
	}
	return
}

func main() {
	addresses := []Address{Address{Id: 1, Location: "山沟子"}}
	user := User{Id: 1, Name: "王一", Addresses: addresses}
	fmt.Println(user)
	userinfo := UserInfo{}
	Copy(&userinfo, user)
	userinfo.Desc = "我是咸鱼"
	fmt.Println(userinfo)
	user.Addresses[0].Location = "北戴河"
	fmt.Println(userinfo)

}
