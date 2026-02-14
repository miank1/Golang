package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("Hello, ")
	buf.WriteString("World!")

	// Read the complete buffer content
	fmt.Println("Buffer content:", buf.String())

	buf.Write([]byte("Hello "))
	fmt.Fprintf(&buf, "Hello friends how are you")
	buf.WriteTo(os.Stdout)

}
