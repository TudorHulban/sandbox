package main

import (
	"fmt"
	"syscall/js"
)

func exposedFunction() js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) any {
			if len(args) != 1 {
				errorWASM("invalid no of arguments passed")

				return nil
			}

			jsDOC := js.Global().Get("document")
			if !jsDOC.Truthy() {
				errorWASM("unable to get document object")

				return nil
			}

			jsonOuputTextArea := jsDOC.Call("getElementById", "output")
			if !jsonOuputTextArea.Truthy() {
				errorWASM("unable to get output text area")

				return nil
			}

			stringInput := args[0].String()

			if len(stringInput) == 0 {
				errorWASM("nil input")

				return nil
			}

			jsonOuputTextArea.Set(
				"value",
				fmt.Sprintf("JS input direct %s\n", stringInput),
			)

			return nil
		},
	)
}
