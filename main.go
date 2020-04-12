package main

import (
	"fmt"
	"os"

	"github.com/keitam913/accware/api/di"
)

func main() {
	dc := di.Container{}
	r := dc.Router()
	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err)
	}
}
