package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"os"
)

func main() {
	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create a new function with two integer parameters.
	addFunc := ir.NewFunc("addFn", types.I32, ir.NewParam("a", types.I32), ir.NewParam("b", types.I32))
	m.Funcs = append(m.Funcs, addFunc)

	// Create a new basic block for the add function.
	addBB := addFunc.NewBlock("entry")

	// Add the two parameters.
	sum := addBB.NewAdd(addFunc.Params[0], addFunc.Params[1])

	// Return the result.
	addBB.NewRet(sum)

	// Write the LLVM IR to a file.
	irFile := "./llvm/add/add_fn.ll"
	err := os.WriteFile(irFile, []byte(m.String()), 0644)
	if err != nil {
		fmt.Printf("Failed to write LLVM IR to file: %v\n", err)
		return
	}
}
