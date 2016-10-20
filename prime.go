// Concurrent computation of pi.
// See https://goo.gl/la6Kli.
//
// This demonstrates Go's ability to handle
// large numbers of concurrent processes.
// It is an unreasonable way to calculate pi.
package main

import "github.com/rwcarlsen/goexif/exif" 
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	searchDir := "/Users/cieclarke/Pictures/Pictures/test"

    fileList := []string{}
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        
		//fmt.Println(fileName)
		if !f.Mode().IsDir() {
			fileList = append(fileList, path)
		}

        return nil
    })

	if err != nil {
		log.Fatal(err)
	}

    for _, file := range fileList {
        fmt.Println(file)
		//fmt.Println("test")

		fi, err2 := os.Open(file)
		
		if err2 != nil {
			//fmt.Println("err2")
			log.Fatal(err2)
		}

		x, err3 := exif.Decode(fi)

		if err3 != nil {
			//fmt.Println("err3")
			log.Fatal(err3)
		}

		camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
		fmt.Println(camModel.StringVal())
		//fmt.Println("test2")

		focal, _ := x.Get(exif.FocalLength)
		numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
		fmt.Printf("%v/%v", numer, denom)
		//fmt.Println("test3")

		// Two convenience functions exist for date/time taken and GPS coords:
		tm, _ := x.DateTime()
		fmt.Println("Taken: ", tm)
		//fmt.Println("test4")

		lat, long, _ := x.LatLong()
		fmt.Println("lat, long: ", lat, ", ", long)
    }

}
