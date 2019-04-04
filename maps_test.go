package maps

import "fmt"

func ExamplePrintln_stringKeys() {
	var m = map[string]int{
		"b": 2,
		"a": 1,
		"c": 3,
	}
	Println(m)
	// Output: map[a:1 b:2 c:3]
}

func ExamplePrintln_intKeys() {
	var m = map[int]string{
		2: "b",
		1: "a",
		3: "c",
	}
	Println(m)
	// Output: map[1:a 2:b 3:c]
}

func ExamplePrintln_stringIntKeys() {
	var m = map[string]map[int]int{
		"b": {3: 3, 2: 2},
		"a": {2: 1},
		"c": {2: 1, 3: 2},
	}
	Println(m)
	// Output: map[a:map[2:1] b:map[2:2 3:3] c:map[2:1 3:2]]
}

func ExamplePrintln_uintKeys() {
	var m = map[uint]string{
		2: "b",
		1: "a",
		3: "c",
	}
	Println(m)
	// Output: map[1:a 2:b 3:c]
}

func ExamplePrintln_floatKeys() {
	var m = map[float32]string{
		2: "b",
		1: "a",
		3: "c",
	}
	Println(m)
	// Output: map[1:a 2:b 3:c]
}

func ExamplePrintln_unsupportedKeys() {
	defer func() {
		fmt.Println(recover())
	}()
	var m = map[bool]int{true: 1, false: 0}
	Println(m)
	// Output: unsupported key type: bool
}

func ExamplePrintln_nonMap() {
	var p *int
	Println(1, nil, p)
	// Output: 1 <nil> <nil>
}
