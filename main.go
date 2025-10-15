package main

import (
	"log"
	"net/http"
	"wallet/helper"
)

func main() {
	defer helper.DB.Close()
	http.HandleFunc("/", helper.HomeHandler)
	http.HandleFunc("/login", helper.LoginHandler)
	http.HandleFunc("/register", helper.RegisterHandler)
	http.HandleFunc("/dashboard", helper.AuthMiddleware(helper.DashboardHandler))
	http.HandleFunc("/payment", helper.AuthMiddleware(helper.PaymentHandler))
	http.HandleFunc("/update-balance", helper.AuthMiddleware(helper.UpdateBalanceHandler))
	http.HandleFunc("/logout", helper.AuthMiddleware(helper.LoginHandler))
	http.HandleFunc("/update", helper.AuthMiddleware(helper.UpdateHandler))
	http.HandleFunc("/history", helper.AuthMiddleware(helper.HistoryHandler))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server running at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
