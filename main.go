package main

import (
	"ekinoksgcc/api"
	"ekinoksgcc/repository"
	"net/http"
	"strings"
)

func main() {
	repository.InitDB()
	defer repository.DisposeDB()
	mux := http.NewServeMux()
	handleEndpoints(mux)
	http.ListenAndServe(":8080", slashStripper(mux))
}

func handleEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/events", api.HandlerGetAllEvents)
	mux.HandleFunc("/login", api.HandlerUserLogin)
	mux.HandleFunc("/register", api.HandlerUserRegister)
	mux.HandleFunc("/logout", api.HandlerUserLogout)
	mux.HandleFunc("/user/orders", api.HandlerGetUserOrders)
	mux.HandleFunc("/user/orders/new", api.HandlerPlaceOrder)
	mux.HandleFunc("/product", api.HandlerProductInfo)
	mux.HandleFunc("/product/add", api.HandlerAddProduct)
	mux.HandleFunc("/product/update", api.HandlerUpdateProduct)
	mux.HandleFunc("/user/add", api.HandlerAddUser)
}

func slashStripper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
