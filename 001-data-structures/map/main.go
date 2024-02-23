package main

import (
	"fmt"
	"maps"
)

func main() {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	map1 := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	map1["key1"] = 7
	map1["key2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("map:", map1)

	// Get a value for a key with name[key].
	value1 := map1["key1"]
	fmt.Println("value1:", value1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	value3 := map1["key3"]
	fmt.Println("value3:", value3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(map1))

	// The builtin delete removes key/value pairs from a map.
	delete(map1, "key2")
	fmt.Println("map:", map1)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(map1)
	fmt.Println("map:", map1)

	// The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, present := map1["key2"]
	fmt.Println("present:", present)

	// You can also declare and initialize a new map in the same line with this syntax.
	firstMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("firstMap:", firstMap)

	// The maps package contains a number of useful utility functions for maps.
	secondMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("secondMap:", secondMap)
	if maps.Equal(firstMap, secondMap) {
		fmt.Println("firstMap == secondMap")
	}

	map1["key1"] = 2

	if _, ok := map1["key1"]; ok {
		// 'key1' exists within the map
		fmt.Println("'key1' exists")
	} else {
		fmt.Println("'key1' doesn't exist")
	}

	if _, ok := map1["key5"]; ok {
		// 'key5' exists within the map
		fmt.Println("'key5' exists")
	} else {
		fmt.Println("'key5' doesn't exist")
	}

}
