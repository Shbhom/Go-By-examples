package main

/*

maps is a key value data structure
type->[keyType]:[valueType]

created using make()

*/

import (
	"fmt"
	"maps"
)

func Main8() {

	m := make(map[string]int)

	m["first"] = 1
	fmt.Println("map1:", m)

	//to get a value from a key
	value := m["first"]
	fmt.Println(value)

	//a second optional field is returned for the map[key] which tells whether that key is present in the map or not
	_, isPresent := m["first"]
	fmt.Println("is there any value for key \"first\":", isPresent)

	fmt.Println(len(m))

	//delete a key value from a map
	fmt.Println(m)
	delete(m, "first")
	fmt.Println("map1 :", m)

	map2 := map[string]int{"first": 1, "second": 2}
	fmt.Println(maps.Equal(m, map2))

	clear(map2)
	fmt.Println("map2:", map2)

}
