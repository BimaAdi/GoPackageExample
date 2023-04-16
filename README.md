# Go Package Example
for testing go get from github

## Instalation
`go get github.com/BimaAdi/GoPackageExample`

## How to Use
```
package main

import gopackageexample "github.com/BimaAdi/GoPackageExample"

func main() {
	gopackageexample.HelloWorld()              // Hello World
	println(gopackageexample.Greeting("Bima")) // Hello Bima
	println(gopackageexample.Add(1, 2))        // 3
	println(gopackageexample.Sub(2, 2))        // 1
	println(gopackageexample.Mul(2, 5))        // 10
	println(gopackageexample.Div(6, 3))        // 2
}

```
