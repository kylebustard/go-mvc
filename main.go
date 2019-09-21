package main

import (
	"net/http"

	"github.com/kylebustard/go-mvc/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
