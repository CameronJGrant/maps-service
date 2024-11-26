package datastore

import "reflect"

type OptionalMap struct {
	filter map[string]interface{}
}

func Optional() OptionalMap {
	return OptionalMap{filter: map[string]interface{}{}}
}

func (q OptionalMap) Add(column string, value interface{}) OptionalMap {
	q.filter[column] = value
	return q
}

func (q OptionalMap) AddNotNil(column string, value interface{}) OptionalMap {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return q
	}
	q.Add(column, value)
	return q
}

func (q OptionalMap) Map() map[string]interface{} {
	return q.filter
}
