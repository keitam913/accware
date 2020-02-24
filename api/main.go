package main

import "github.com/keitam913/accware/api/di"

func main() {
	dc := di.Container{}
	r := dc.Router()
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
