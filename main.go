package main

import (
	"net/http"

	"github.com/ljg294/fd-order/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
