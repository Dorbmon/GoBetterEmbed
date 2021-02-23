package main

import (
	".."
	"fmt"
)
//go:embed t.txt
var f GoBetterEmbed.File
func main() {
	path,err := f.ToTempFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(path)
	//fmt.Println(ioutil.ReadAll(os.Open(path)))
}
