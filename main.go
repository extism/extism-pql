package main

import (
	"github.com/extism/go-pdk"
	"github.com/runreveal/pql"
)

//export compile
func compile() int32 {
	pqlCode := pdk.InputString()
	sql, err := pql.Compile(pqlCode)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString(sql)

	return 0
}

func main() {}
