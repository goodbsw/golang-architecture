package main

import (
	"fmt"
	architecture "github.com/goodbsw/golang-architecture"
	"github.com/goodbsw/golang-architecture/storage/mongo"
	"github.com/goodbsw/golang-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := architecture.Person{
		First: "Jimmy",
	}
	p2 := architecture.Person{
		First: "James",
	}

	ps := architecture.NewPersonService(dbm)

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))

	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}
