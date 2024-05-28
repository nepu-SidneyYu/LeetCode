package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func NewStringID() string {
	return bson.NewObjectId().Hex()
}

func main() {
	strid := NewStringID()
	fmt.Println(strid)
}
