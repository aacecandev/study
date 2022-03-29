// unordered collection of key/value pairs
// implemented by hash tables
// provide efficient add, get and delete operations

package main

import (
	"fmt"
)

func main() {

	// <map_name> := map[<key_data_type]<value_data_type>]{<key>:<value>, <key>:<value>, ...}
	codes := map[string]string{"en": "English", "fr": "French", "de": "German"}
	fmt.Println(codes)

	fmt.Println()

	// <map_name> := make(map[<key_type>]<value_type>, <initial_capacity>)
	codes2 := make(map[string]int)
	fmt.Println(codes2)

	fmt.Println()

	fmt.Println(len(codes))
	fmt.Println(len(codes2))

	fmt.Println()

	fmt.Println(codes["en"])

	fmt.Println()

	// value, found := <map_name>[<key>]
	value, found := codes["en"]
	fmt.Println(value, found)

	fmt.Println()

	// adding key-value pair
	codes["ru"] = "Russian"
	fmt.Println(codes)

	fmt.Println()

	// updating value
	codes["ru"] = "Ukraine"
	fmt.Println(codes)

	fmt.Println()

	// deleting key-value pair
	delete(codes, "ru")
	fmt.Println(codes)

	fmt.Println()

	// truncate a map
	// a
	for key := range codes {
		delete(codes, key)
	}
	fmt.Println(codes)

	// b
	// codes3 := map[string]string{"en": "English", "fr": "French", "es": "Spanish"}
	// codes3 = make(map[string]string)
}
