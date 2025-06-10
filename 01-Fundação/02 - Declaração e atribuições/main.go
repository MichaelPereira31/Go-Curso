package main

const a = "Hello, World!"
type ID int

var (
	b bool   = true
	c int    = 10
	d string = "Michael"
	e float64
)

func main() {
	a := "Hello, World!" // na primeira vez é atribuição
	a = "Hello, World"  // atribuição
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
}
