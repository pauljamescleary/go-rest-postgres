package main

import (
	router "github.com/pauljamescleary/gomin/pkg/common"
)

func main() {
	e := router.SetupRouter()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
