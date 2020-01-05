package main

import "github.com/keitam913/accware-api/di"

func main() {
	dc := di.Container{}
	r := dc.Router()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
