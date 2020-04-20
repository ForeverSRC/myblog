package functions

import (
	"bytes"
	"github.com/russross/blackfriday"
	"html/template"
)

func Unescaped(x string) interface{} {
	input := []byte(x)
	input = bytes.Replace(input, []byte("\r"), nil, -1)
	unsafe := blackfriday.Run(input,blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak))
	return template.HTML(unsafe)
}