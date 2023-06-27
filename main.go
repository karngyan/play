package main

import (
	"fmt"
)

func main() {
	str := "🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫🌟✨💫"
	byteLen := len(str)
	fmt.Printf("string: \"%s\"\nbyteLen: %d\n", str, byteLen) // 132
}

