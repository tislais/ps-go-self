###### [wait... they dont love you like i love you](https://www.youtube.com/watch?v=oIIxlgcuQRU)

## Intro to Maps
- Retrieving a map without a loop, Go's gonna return it in key-sorted order. But when iterating it with a range loop, there's no guarantee what order you will get them in!
- Maps are key-value pairs. (think dictionary, hash table...)

## Syntax and Basics
How to define maps:
```
map[<key-type>]<value-type>
```
Couple things on that,
- Key type needs to be a comparable type. (works with == and !=)
- Key entries have to be unique

As with slices, we can declare and initialize maps in a single go with the `composite literal` form

```
recentHead2HeadWins := map[string]int{
   "Sunderland": 6,
   "Newcastle": 0,
}
```

## Iterating Maps
Example of random order of elements when iterated through
```
testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}

	for mapKey, mapVal := range testMap {
		fmt.Printf("Key is: %v Value is: %v\n", mapKey, mapVal)
	}
```
Will output the following:
```
ps-go-self % go run maps/maploop.go
Key is: B Value is: 2
Key is: D Value is: 4
Key is: F Value is: 6
Key is: G Value is: 7
Key is: H Value is: 8
Key is: I Value is: 9
Key is: A Value is: 1
Key is: E Value is: 5
Key is: C Value is: 3
```

## Updating Maps
Use bracket notation
```
testMap := map[string]int{
    "A": 1,
    "B": 2,
    "C": 3,
    "D": 4,
    "E": 5,
    "F": 6,
    "G": 7,
    "H": 8,
    "I": 9,
}
testMap["A"] = 100
testMap["J"] = 1969
fmt.Println(testMap)
```
Output
```
ps-go-self % go run maps/maploop.go
map[A:100 B:2 C:3 D:4 E:5 F:6 G:7 H:8 I:9 J:1969]
```

If the key already exists, its value will be replaced with the new one. If it doesn't exist, both key and value are created. 

To delete an element, use the built in delete function and reference the key of the element you want to lose.
```
delete(myMap, <key>)
```

## Misc

Maps are reference types. When they get passed to functions, they get passed as references. That means any changes a function makes to the map are visible to the caller but also any other functions. If it's modified, it's modified. It's not acting on a copy.

Maps are cheap. It can be huge but passing it around is cheap as chips because you're just passing around pointers.

It is good practice to specify the size of the map, especially for larger maps. Go can resize when it needs to, but there is a cost to resizing large maps. 

Maps are not thread-safe. They are not safe for concurrency. Basically, it's not defined what happens to them if written to simultaneously.

## Recap

[](../mapsrecap.png)

