
package main

import (
	"fmt"
	"base"
)

func main() {
	info := base.GetInfo("/mnt/c/tmp/blsub.mbtiles")
   	fmt.Println(info)
}
