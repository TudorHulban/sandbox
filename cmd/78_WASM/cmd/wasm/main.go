package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("exposedFunction", exposedFunction())

	fmt.Println("Go Web Assembly")

	<-make(chan struct{})
}

func exposedFunction() js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) any {
			if len(args) != 1 {
				return map[string]interface{}{
					"error": "invalid no of arguments passed",
				}
			}

			jsDOC := js.Global().Get("document")
			if !jsDOC.Truthy() {
				return map[string]interface{}{
					"error": "unable to get document object",
				}
			}

			jsonOuputTextArea := jsDOC.Call("getElementById", "output")
			if !jsonOuputTextArea.Truthy() {
				return map[string]interface{}{
					"error": "unable to get output text area",
				}
			}

			stringInput := args[0].String()

			if len(stringInput) == 0 {
				return map[string]interface{}{
					"error": "nil input",
				}
			}

			jsonInputTextArea := jsDOC.Call("getElementById", "input")
			if !jsonInputTextArea.Truthy() {
				return map[string]interface{}{
					"error": "unable to get input text area",
				}
			}

			jsonInputTextArea.Set(
				"value",
				"",
			)

			jsonOuputTextArea.Set(
				"value",
				fmt.Sprintf("JS input direct %s\n", stringInput),
			)

			return nil
		},
	)
}
