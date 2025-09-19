package main

func main() {
	Add(2, 3)
	Sub(2, 3)
	Mul(2, 3)
	Div(2, 3)
	Mod(2, 3)

}

func Add(a, b int) {
	var c = a + b
	println(c)
}

func Sub(a, b int) {
	var c = a - b
	println(c)
}

func Mul(a, b int) {
	var c = a * b
	println(c)
}

func Div(a, b int) {
	var c = a / b
	println(c)
}

func Mod(a, b int) {
	var c = a % b
	println(c)
}
