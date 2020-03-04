package main

import "fmt"

type mongo map[string]string
type postg map[string]string

type accessor interface {
	save(k string, v string)
	retrieve(k string) string
}

func (m mongo) save(k string, v string) {
	m[k] = v
}

func (m mongo) retrieve(k string) string {
	return m[k]
}

func (p postg) save(k string, v string) {
	p[k] = v
}

func (p postg) retrieve(k string) string {
	return p[k]
}

func put(a accessor, k string, v string) {
	a.save(k, v)
}

func get(a accessor, k string) string {
	return a.retrieve(k)
}

func main() {
	m := mongo{}
	p := postg{}

	put(m, "name", "sw")
	fmt.Println(get(m, "name"))
	put(p, "name", "yl")
	fmt.Println(get(p, "name"))

}
