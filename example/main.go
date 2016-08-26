package main

import (
	"fmt"

	"github.com/nanoeru/termcol"
)

func main() {
	fmt.Println(termcol.Blue, "sample")
	fmt.Printf("Hello world\n")
	fmt.Println(termcol.Red, "sample", termcol.Reset)
	fmt.Printf("Hello world\n")

	fmt.Println(termcol.Bright.Red, "orange", termcol.Default)
	fmt.Println(termcol.Green.Wrap("green", termcol.Magenta, "magenta"), "default")
	fmt.Println(termcol.Green.WrapSpace("green", termcol.Magenta, "magenta"), "default")
	fmt.Print(termcol.Green.Wrap("green", termcol.Magenta, "magenta"), "default")
	fmt.Println()
	fmt.Print(termcol.Green.WrapF("%s%s%s", "green", termcol.Magenta, "magenta"), "default")
}
