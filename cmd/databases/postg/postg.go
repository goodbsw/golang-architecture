package postg

type Db map[string]string

func (pg Db) Save(k string, v string) {
	pg[k] = v
}

func (pg Db) Retrieve(k string) string {
	return pg[k]
}
