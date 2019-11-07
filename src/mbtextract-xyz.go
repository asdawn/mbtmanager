package main

import (
	"fmt"
	"strconv"
	"base"
	"os"
)
//writeTileFile
//mbtextract-xyz [mbtpathIN] [zoom] [x] [y] [tileOUT]
func main() {
	paramSize := len(os.Args)
	if paramSize >= 6 {
		path := os.Args[1]
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				base.Err(2, path)
			}else{
				base.Err(4, path)
			}
		}else{
			var z, x, y int
			z, err = strconv.Atoi(os.Args[2])
			if err != nil {
				base.Err(1, "z="+os.Args[2])
				return
			}
			x, err = strconv.Atoi(os.Args[3])
			if err != nil {
				base.Err(1, "x="+os.Args[3])
				return
			}
			y, err = strconv.Atoi(os.Args[4])
			if err != nil {
				base.Err(1, "y="+os.Args[4])
				return
			}

			tilepath := os.Args[5]
			tiledata := base.GetTile(path, z, x, y)
			if(tiledata != nil){
				base.WriteTileFile(tilepath, tiledata)
				fmt.Printf("Extract tile (z=%d, x=%d, y=%d) to %s\n",z, x, y, tilepath)
			}
		}
	}else{
		base.Err(1, "\nUsage: mbtextract-xyz [mbtpathIN] [zoom] [x] [y] [tileOUT]")
		return
	}
}