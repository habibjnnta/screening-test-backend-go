package main

import (
	"context"
	"screening-test/connection"
	"screening-test/router"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
	db := connection.ConnectToDb()
	ctx := context.Background()
	rh := &router.Handlers{
		Ctx: ctx,
		DB:  db,
		R:   r,
	}
	rh.Routes()
	
	// Used Port Default
	r.Run()

	// Used Port Optional
	// r.Run(":8081")
}