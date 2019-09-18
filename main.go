package main

import (
	"net/http"

	"github.com/kylebustard/lowendcode/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
