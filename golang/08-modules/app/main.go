package main

import (
	"fmt"
	locallib "my-local-lib" // Aliased import to be safe

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("UUID:", id)

	// Call local lib
	locallib.SecretFunction()
}
