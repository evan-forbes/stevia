package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers stevia-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/stevia/sweetners/{id}", getSweetnerHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/stevia/sweetners", listSweetnerHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/stevia/sweetners", createSweetnerHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/stevia/sweetners/{id}", updateSweetnerHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/stevia/sweetners/{id}", deleteSweetnerHandler(clientCtx)).Methods("POST")

}
