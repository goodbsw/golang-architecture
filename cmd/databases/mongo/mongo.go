package mongo

type Db map[string]string

func (mg Db) Save(k string, v string) {
	mg[k] = v
}

func (mg Db) Retrieve(k string) string {
	return mg[k]
}
