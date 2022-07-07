A map is a data structure that provides you with an unordered collection of key/value pairs.
You store values into the map based on a key

The strength of a map is its ability to retrieve data quickly based on the key. A key works like an index, pointing to the value you associate with that key

Internals

Maps are collections, and you can iterate over them just like you do with arrays and slices. But maps are unordered collections, and there’s no way to predict the order in which the key/value pairs will be returned. Even if you store your key/value pairs in the same order, every iteration over a map could return a different order. This is because a map is implemented using a hash table

The purpose of the hash function is to generate an index that evenly distributes key/value pairs across all available buckets. 

what the internals of a bucket look like. There are two data structures that contain the data for the map. First, there’s an array with the top eight high order bits (HOB) from the same hash key that was used to select the bucket. This array distinguishes each individual key/value pair stored in the respective bucket. Second, there’s an array of bytes that stores the key/value pairs. The byte array packs all the keys and then all the values together for the respective bucket. The packing of the key/value pairs is implemented to minimize the memory required for each bucket. 


Creating and initializing 

// Create a map with a key of type string and a value of type int.
dict := make(map[string]int)

// Create a map with a key and value of type string.
// Initialize the map with 2 key/value pairs.
dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

The map key can be a value from any built-in or struct type as long as the value can be used in an expression with the == operator. Slices, functions, and struct types that contain slices can’t be used as map keys. This will produce a compiler error. 

// Create a map using a slice of strings as the key.
dict := map[[]string]int{}

Compiler Exception:
invalid map key type []string

There’s nothing stopping you from using a slice as a map value. This can come in handy when you need a single map key to be associated with a collection of data. 

// Create a map using a slice of strings as the value.
dict := map[int][]string{}

Working with maps 

// Create an empty map to store colors and their color codes.
colors := map[string]string{}

// Add the Red color code to the map.
colors["Red"] = "#da1337"

You can create a nil map by declaring a map without any initialization. A nil map can’t be used to store key/value pairs. Trying will produce a runtime error. 

// Create a nil map by just declaring the map.
var colors map[string]string

// Add the Red color code to the map.
colors["Red"] = "#da1337"

Runtime Error:
panic: runtime error: assignment to entry in nil map

When retrieving a value from a map, you have two choices. You can retrieve the value and a flag that explicitly lets you know if the key exists.

// Retrieve the value for the key "Blue".
value, exists := colors["Blue"]

// Did this key exist?
if exists {
    fmt.Println(value)
}

The other option is to just return the value and test for the zero value to determine if the key exists. This will only work if the zero value is not a valid value for the map. 

// Retrieve the value for the key "Blue".
value := colors["Blue"]

// Did this key exist?
if value != "" {
    fmt.Println(value)
}

Iterating over a map is identical to iterating over an array or slice. You use the keyword range; but when it comes to maps, you don’t get back the index/value, you get back the key/value pairs. 

// Create a map of colors and color hex codes.
colors := map[string]string{
	"AliceBlue":   "#f0f8ff",
	"Coral":       "#ff7F50",
	"DarkGray":    "#a9a9a9",
	"ForestGreen": "#228b22",
}

// Display all the colors in the map.
for key, value := range colors {
	fmt.Printf("Key: %s  Value: %s\n", key, value)
}

If you want to remove a key/value pair from the map, you use the built-in function delete. 

// Remove the key/value pair for the key "Coral".
delete(colors, "Coral")

// Display all the colors in the map.

for key, value := range colors {
    fmt.Printf("Key: %s  Value: %s\n", key, value)
}

Passing maps between functions 

Passing a map between two functions doesn’t make a copy of the map. In fact, you can pass a map to a function and make changes to the map, and the changes will be reflected by all references to the map. 

func main() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		 "AliceBlue":   "#f0f8ff",
		 "Coral":       "#ff7F50",
		 "DarkGray":    "#a9a9a9",
		 "ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
			fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

	// Call the function to remove the specified key.
	removeColor(colors, "Coral")

	// Display all the colors in the map.
	for key, value := range colors {
			fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}

// removeColor removes keys from the specified map.
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

Key: AliceBlue Value: #F0F8FF
Key: Coral Value: #FF7F50
Key: DarkGray Value: #A9A9A9
Key: ForestGreen Value: #228B22

Key: AliceBlue Value: #F0F8FF
Key: DarkGray Value: #A9A9A9
Key: ForestGreen Value: #228B22