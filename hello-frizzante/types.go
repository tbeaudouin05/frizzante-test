//go:build types

package main

import (
	"main/lib/core/types"
	"main/lib/routes/handlers/todos"
)

func main() {
	types.Generate[todos.Props]()
}
