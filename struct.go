package main

import (
	"encoding/json"
	"fmt"
)

type rec struct {
	Int1 int
	Int2 int
	Str  string
}

const payload = "[{\"Int1\": 34, \"Str\": \"abba\"}]"

func main() {
	var r []rec
	err := json.Unmarshal([]byte(payload), &r)
	fmt.Println(err)
	fmt.Printf("%#v\n", r)
}
