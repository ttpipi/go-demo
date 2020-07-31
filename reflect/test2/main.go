package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `name`
	Age  int    `age`
}

type User2 struct {
	Name string `json:"jname" bson:"bname"`
	Age  int    `json:"jage" bson:"bage"`
}

func main() {
	//testJson()
	//testTag()
	testDifferentTag()
}

func testJson() {
	var u User
	h := `{"name": "张三", "age": 20}`
	json.Unmarshal([]byte(h), &u)
	fmt.Println(u)

	newJson, _ := json.Marshal(&u)
	fmt.Println(string(newJson))
}

func testTag() {
	var u User
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}
}

func testDifferentTag() {
	var u User2
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag.Get("json"))
		fmt.Println(sf.Tag.Get("bson"))
	}
}
