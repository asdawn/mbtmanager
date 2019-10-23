package main

import (
	"fmt"
	"base"
	"os"
)

func main() {
	paramSize := len(os.Args)
	if paramSize > 1 {
		path := os.Args[1]
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				base.Err(2, path)
			}else{
				base.Err(4, path)
			}
		}else{
			info := base.GetInfoField(path, "minzoom")
			   fmt.Println(info)
		}
	}else{
		base.Err(1, "no input .mbtiles file specified")
		return
	}
}