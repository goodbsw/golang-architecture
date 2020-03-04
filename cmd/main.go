package main

import (
	"fmt"

	architecture "github.com/goodbsw/golang-architecture"
	"github.com/goodbsw/golang-architecture/cmd/databases/mongo"
	"github.com/goodbsw/golang-architecture/cmd/databases/postg"
)

func main() {
	m := mongo.Db{}
	p := postg.Db{}

	architecture.Put(m, "name", "sw")
	architecture.Put(p, "name", "yl")
	fmt.Println(architecture.Get(m, "name"))
	fmt.Println(architecture.Get(p, "name"))
}
