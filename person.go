package main

// Person JSON binding 이 가능하려면 field 가 공개되어 있어야 한다.
type Person struct {
	Name string `json:"user"`
	Age  int    `json:"age"`
}
