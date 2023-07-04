package main

import "syscall/js"

func increment(x int) int {
	return x + 1
}
func decrement(x int) int {
	return x - 1
}

func main() {
	c := make(chan struct{}, 0)
	// setup values and variables
	value := 0
	countElement := js.Global().Get("document").Call("getElementById", "count")
	addElement := js.Global().Get("document").Call("getElementById", "add")
	subElement := js.Global().Get("document").Call("getElementById", "subtract")

	// on click add
	addElement.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		value = increment(value)
		countElement.Set("innerText", value)
		return nil
	}))
	// on click subtract
	subElement.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		value = decrement(value)
		countElement.Set("innerText", value)
		return nil
	}))
	<-c
}
