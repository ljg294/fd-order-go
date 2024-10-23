package main

import (
	"github.com/ljg294/fd-order/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
