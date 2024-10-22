package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ErwinSalas/go-eda/services/order-service/pkg/datastore"
	"github.com/ErwinSalas/go-eda/services/order-service/pkg/types"
)

func createOrder(orderDs datastore.IDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		order := &types.Order{}

		err := json.NewDecoder(r.Body).Decode(order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ID, err := orderDs.InsertOrder(ctx, order)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ID)
	}
}
