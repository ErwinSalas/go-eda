package api

import (
	"fmt"
	"net/http"

	"github.com/ErwinSalas/go-eda/services/order-service/pkg/app"
)

func NewRouter(port int, app app.App) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/order", createOrder(app.Datastore))

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
