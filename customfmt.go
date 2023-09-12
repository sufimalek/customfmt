package customfmt

import (
	"github.com/fatih/color"
)

// PrintColorful prints text in the specified color.
func PrintColorful(text string, textColor color.Attribute) {
	c := color.New(textColor)
	c.Println(text)
}

// ColorPrint prints text in the specified color.
func Print(text string, textColor color.Attribute) {
	c := color.New(textColor)
	c.Print(text)
}

// ColorPrintln prints text in the specified color and adds a newline.
func Println(text string, textColor color.Attribute) {
	c := color.New(textColor)
	c.Println(text)
}

// ColorPrintf formats and prints text in the specified color.
func Printf(text string, textColor color.Attribute, args ...interface{}) {
	c := color.New(textColor)
	c.Printf(text, args...)
}
