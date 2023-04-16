package gopackageexample

import "testing"

func TestHelloWorld(t *testing.T) {
	HelloWorld()
}

func TestGreeting(t *testing.T) {
	result := Greeting("Bima")
	if result != "Hello Bima" {
		t.Errorf("test greeting failed")
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 1)
	if result != 2 {
		t.Errorf("test add failed")
	}
}

func TestSub(t *testing.T) {
	result := Sub(1, 1)
	if result != 0 {
		t.Errorf("test sub failed")
	}
}

func TestMul(t *testing.T) {
	result := Mul(1, 1)
	if result != 1 {
		t.Errorf("test mul failed")
	}
}
func TestDiv(t *testing.T) {
	result := Div(1, 1)
	if result != 1 {
		t.Errorf("test div failed")
	}
}
