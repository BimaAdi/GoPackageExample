package gopackageexample

// print Hello World
//
//		gopackageexample.HelloWorld()
//	 	// output "Hello World"
func HelloWorld() {
	println("Hello World")
}

// print Hello {name}
//
//		gopackageexample.Hello("Bima")
//	 	// output "Hello Bima"
func Greeting(name string) string {
	return "Hello " + name
}

// Add two integer x + y
//
//		gopackageexample.Add(1, 2)
//	 	// output 3
func Add(x int, y int) int {
	return x + y
}

// Subtract two integer x - y
//
//		gopackageexample.Sub(2, 2)
//	 	// output 0
func Sub(x int, y int) int {
	return x - y
}

// Multiply two integer x * y
//
//		gopackageexample.Mul(2, 5)
//	 	// output 10
func Mul(x int, y int) int {
	return x * y
}

// Divide two integer x / y
//
//		gopackageexample.Div(6, 3)
//	 	// output 2
func Div(x int, y int) int {
	return x / y
}
