package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {
	jsonParsed, err := gabs.ParseJSON([]byte(`{
	"outer":{
		"inner":{
			"value1":10,
			"value2":22
		},
		"inner2":{
			"value1":20,
			"array1":[
				30, 40
			]
		}
	}
}`))
	if err != nil {
		panic(err)
	}

	value, ok := jsonParsed.Path("outer.inner.value1").Data().(float64)
	fmt.Println("value:", value, "\tok:", ok)

	value, ok = jsonParsed.Search("outer", "inner", "value1").Data().(float64)
	fmt.Println("value:", value, "\tok:", ok)

	value, ok = jsonParsed.Search("outer", "inner2", "array1", "1").Data().(float64)
	fmt.Println("value:", value, "\tok:", ok)

	gObj, err := jsonParsed.JSONPointer("/outer/inner2/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	fmt.Println("value:", value, "\tok:", ok)

	value, ok = jsonParsed.Path("does.not.exist").Data().(float64)
	fmt.Println("value:", value, "\tok:", ok)

	exists := jsonParsed.Exists("outer", "inner", "value1")
	fmt.Println("exists:", exists)

	exists = jsonParsed.ExistsP("does.not.exist")
	fmt.Println("exists:", exists)
}
