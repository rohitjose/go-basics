package main

import "fmt"

func main() {
	//start basic OMIT
	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//end basic OMIT

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.
	//start basic2 OMIT
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//end basic2 OMIT

	// start maps OMIT
	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// end maps OMIT

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	// start strings OMIT
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	// end strings OMIT
}
