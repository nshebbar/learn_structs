package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Foo struct {
	A int
	b string
}
 type Bar struct {
	 C string
	 F Foo
 }

 type Baz struct {
	 D string
	 Foo
 }

func main() {

	bob := `{"name": "Bob", "age": 30}`
	var bb Person
	json.Unmarshal([]byte(bob), &bb)
	fmt.Println(bb)
	bob2, _ := json.Marshal(bb)
	fmt.Println(string (bob2))

	e := Foo{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: e}
	fmt.Println("Bar:", b1.F.A)
	b2 := Baz{D: "Nancy", Foo: e}
	fmt.Println("Baz:", b2.A)

	var f4 Foo = b2.Foo
	fmt.Println("f4:", f4)


	f := Foo{}
	fmt.Println(f)

	f1 := Foo {
		A: 20,
	}

	var f2 Foo
	f2 = f1
	f2.A = 100
	fmt.Println("f2:", f2.A)
	fmt.Println("f1:", f1.A)

	var f3 *Foo = &f
	f3.A = 200
	fmt.Println("f3:", f3.A)
	fmt.Println("f1:", f1.A)




	g := Foo{10, "hello"}
	fmt.Println(g)

	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)
}