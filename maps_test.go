package maps

func ExamplePrintln() {
	var m = map[string]int{
		"b": 2,
		"a": 1,
		"c": 3,
	}
	Println(m)
	// Output: map[a:1 b:2 c:3]
}
