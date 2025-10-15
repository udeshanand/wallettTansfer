package helper

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type DashboardData struct {
	Name    string
	Balance float64
}

type Transaction struct {
	Type      string
	OtherUser string
	Amount    float64
	Date      string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/login.html")
	if err != nil {
		http.Error(w, "Error loading login page: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	userid := r.FormValue("userid")
	password := r.FormValue("password")

	if err := ValidateUser(userid, password); err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Invalid UserID or Password"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_userId",
		Value:    userid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   300,
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/register.html")
	if err != nil {
		http.Error(w, "Error loading register page: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		userid := r.FormValue("userid")
		username := r.FormValue("username")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm_password")

		if err := RegisterUser(userid, username, email, mobile, password, confirm); err != nil {
			tmpl.Execute(w, map[string]string{"Error": err.Error()})
			return
		}

		tmpl.Execute(w, map[string]string{"Success": "User registered successfully! You can login now."})
		return
	}

	tmpl.Execute(w, nil)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_userId")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_userId")
	userid := cookie.Value

	var username string
	var balance float64
	err := DB.QueryRow("SELECT username, balance FROM users WHERE userId=?", userid).Scan(&username, &balance)
	if err != nil {
		http.Error(w, "User not found or DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/dashboard.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, DashboardData{Name: username, Balance: balance})
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/transection.html")
	if err != nil {
		fmt.Println("Template error:", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("session_userId")
	if err != nil || cookie.Value == "" {
		fmt.Println("No cookie found, redirecting to /login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	senderID := cookie.Value

	if r.Method == http.MethodGet {

		tmpl.Execute(w, nil)
		return
	}

	// POST request
	receiverID := r.FormValue("to_userid")
	amountStr := r.FormValue("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount input")
		tmpl.Execute(w, map[string]string{"Error": "Invalid amount"})
		return
	}

	// transaction
	err = Transection_process(senderID, receiverID, amount)
	if err != nil {
		fmt.Println("Transaction failed:", err)
		tmpl.Execute(w, map[string]string{"Error": err.Error()})
		return
	}

	fmt.Println(" Transaction successful")
	tmpl.Execute(w, map[string]string{"Success": fmt.Sprintf("₹%.2f sent to %s successfully!", amount, receiverID)})
}

// UPDATE BALANCE
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_userId")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userid := cookie.Value

	tmpl, err := template.ParseFiles("static/update.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {

		var username, email, mobile string
		err := DB.QueryRow("SELECT username, email, Mobile_no FROM users WHERE userId=?", userid).
			Scan(&username, &email, &mobile)
		if err != nil {
			tmpl.Execute(w, map[string]string{"Error": "User not found"})
			return
		}

		tmpl.Execute(w, map[string]string{
			"Username": username,
			"Email":    email,
			"Mobile":   mobile,
		})
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm_password")

		// validate passwords
		if password != "" && password != confirm {
			tmpl.Execute(w, map[string]string{
				"Error":    "Passwords do not match",
				"Username": username,
				"Email":    email,
				"Mobile":   mobile,
			})
			return
		}

		// update user
		err := UpdateUser(userid, username, email, mobile, password)
		if err != nil {
			tmpl.Execute(w, map[string]string{
				"Error":    err.Error(),
				"Username": username,
				"Email":    email,
				"Mobile":   mobile,
			})
			return
		}

		tmpl.Execute(w, map[string]string{
			"Success":  "Profile updated successfully!",
			"Username": username,
			"Email":    email,
			"Mobile":   mobile,
		})
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_userId",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_userId")
	userid := cookie.Value

	rows, err := DB.Query("SELECT from_userid, to_userid, amount, transaction_time FROM history WHERE from_userid=? OR to_userid=? ORDER BY transaction_time DESC", userid, userid)
	if err != nil {
		http.Error(w, "Error fetching history: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var from, to string
		var amount float64
		var date string
		rows.Scan(&from, &to, &amount, &date)

		t := Transaction{Amount: amount, Date: date}
		if from == userid {
			t.Type = "Sent"
			t.OtherUser = to
		} else {
			t.Type = "Received"
			t.OtherUser = from
		}
		transactions = append(transactions, t)
	}

	tmpl, err := template.ParseFiles("static/history.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, transactions)
}
func UpdateBalanceHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_userId")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userid := cookie.Value

	tmpl, err := template.ParseFiles("static/balence.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	amountStr := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		tmpl.Execute(w, map[string]string{"Error": "Invalid amount"})
		return
	}

	// Add balance safely
	res, err := DB.Exec("UPDATE users SET balance = balance + ? WHERE userId = ?", amount, userid)
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Failed to add balance"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		tmpl.Execute(w, map[string]string{"Error": "User not found"})
		return
	}

	tmpl.Execute(w, map[string]string{"Success": fmt.Sprintf("₹%.2f added to your wallet!", amount)})
}
