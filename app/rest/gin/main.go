package main

import (
	serviceA "ddd-hexagonal-cqrs-poc/services/serviceA"
	serviceARestGin "ddd-hexagonal-cqrs-poc/services/serviceA/inbound/command/invokers/restgin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	serviceA.Register(nil, nil)

	serviceARestGin.RegisterRoutes(r)

	r.Run(":8080")
}
