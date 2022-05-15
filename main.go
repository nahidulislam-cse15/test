package main

import (
	"database/sql"
	//"encoding/json"

	"fmt"
	"html/template"

	"net/http"
	"net/smtp"
	"math/rand"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var btemp *template.Template
var otp string

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer db.Close()
	//fmt.Println("db connection succesful")
}

type user struct {
	ID                    int
	Name, Email, Password string
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/forget", forgetPassword)
	http.HandleFunc("/registered", register)
	http.HandleFunc("/authentication", authentication)
	http.HandleFunc("/recover", recoverPassword)
	// 	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../asset"))))

	http.ListenAndServe(":8040", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		btemp, err = template.ParseGlob("templates/*.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}

		btemp.ExecuteTemplate(w,"base.gohtml","Home")
	// }else if r.Method == http.MethodPost {
	// 	r.ParseForm()
	// 	found := false
	// 	email := r.FormValue("email")
	// 	pass := r.FormValue("password")
	// 	fmt.Println(email, pass)
	// 	qs := "select email,password from `users`"
	// 	rows, err := db.Query(qs)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	defer rows.Close()
	// 	for rows.Next() {
	// 		var u user
	// 		if err := rows.Scan(&u.Email, &u.Password); err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		if (u.Email == email) && (u.Password == pass) {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	if found {
	// 		fmt.Fprintln(w, `Log in succesfully`)
	// 		btemp, err = template.ParseFiles("templates/base.gohtml")
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}

	// 		btemp.Execute(w, nil)
	// 		fmt.Fprintln(w, `Log in succesfully`)

	// 	} else {
	// 		fmt.Fprintln(w, `username or password is incorrect`)
	// 	}
		

	  }
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//btemp, err = template.ParseFiles("templates/signup.gohtml")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		
		btemp.ExecuteTemplate(w,"signup.gohtml", "Sign Up")
	} else if r.Method == http.MethodPost {
		// btemp, err = template.ParseFiles("webpage/signin.html")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }

		// btemp.Execute(w, nil)
		
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		pass := r.FormValue("password")
		//display in command line
		fmt.Println(name, email, pass)
		// //inser into db
		qs := "INSERT INTO `users` (`Name`, `Email`, `Password`) VALUES ('%s', '%s','%s');"
		sql := fmt.Sprintf(qs, name, email, pass)
		fmt.Println(sql)
		insert, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Fprintln(w, `succesfully registered`)
	}
}

// fmt.Print("form Data:", r.Form)
// fmt.Fprintln(w, `ok`)
//json encodede text
// bs, err := json.Marshal(r.Form)
// if err != nil {
// 	fmt.Println(err)
// }
// w.Header().Set("Content-Type", "application/json")
// fmt.Fprintln(w, string(bs))

// if r.Method == http.MethodGet {

// 	//fmt.Println("signup")
// 	btemp, err = template.ParseFiles("webpage/signup.html")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	btemp.Execute(w, nil)

// } else if r.Method == http.MethodPost {

// 	r.ParseForm()
// 	//fmt.Println("form:", r.Form)

// 	name := r.FormValue("name")
// 	email := r.FormValue("email")
// 	pass := r.FormValue("password")
// 	fmt.Println(name, email, pass)

// 	//plain text
// 	//fmt.Fprintln(w, "OK")
// 	//json encoded text
// 	bs, err := json.Marshal(r.Form)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	fmt.Fprintln(w, string(bs))
// }

func signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// btemp, err = template.ParseFiles("templates/signin.gohtml")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }

		btemp.ExecuteTemplate(w,"signin.gohtml", "SignIn")
	}else if r.Method == http.MethodPost {
		r.ParseForm()
		found := false
		email := r.FormValue("email")
		pass := r.FormValue("password")
		qs := "select email,password from `users`"
		rows, err := db.Query(qs)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var u user
			if err := rows.Scan(&u.Email, &u.Password); err != nil {
				fmt.Println(err)
			}
			if (u.Email == email) && (u.Password == pass) {
				found = true
				break
			}
		}
		if found {
			fmt.Fprintln(w, "Log in succesfully")
			// btemp, err = template.ParseFiles("templates/base.gohtml")
			// if err != nil {
			// 	fmt.Println(err.Error())
			// }

			// btemp.Execute(w, nil)
			// fmt.Fprintln(w, `Log in succesfully`)

		} else {
			fmt.Fprintln(w, `username or password is incorrect`)
		}
	}
}
func forgetPassword(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
	// 	fmt.Println("get")
	// 	btemp, err = template.ParseFiles("webpage/forget_password.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	btemp.ExecuteTemplate(w,"forget_password.gohtml","Forget Password")
}else if r.Method == http.MethodPost {
	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println("post")
	otp = strconv.Itoa(rand.Intn(10000))
	emailSend(email,otp)
}
}
func emailSend(toEmail ,otp string){
	from := "nahidulss@gmail.com" //ex: "John.Doe@gmail.com"
	password :="umxsaeekiabfbhgj"  //app specific password yourAppPassword

	// receiver address
	
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: OTP\n\n"
	body := "Your OTP is:\n" +otp
	fmt.Println(body)
	message := []byte(subject + body)
	fmt.Println(message)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("email successfully sent")
}
func register(w http.ResponseWriter, r *http.Request) {

}

func authentication(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		
	// 	btemp, err = template.ParseFiles("webpage/authentication.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	btemp.ExecuteTemplate(w,"authentication.gohtml","Authentication")
}else if r.Method == http.MethodPost {
	r.ParseForm()
	otp2:= r.FormValue("otp")
	if(otp2 ==otp){
		fmt.Fprintln(w,"verified")
	}else{
		fmt.Fprintln(w, "otp doese not match ")
	}
}
}
func recoverPassword(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		
	// 	btemp, err = template.ParseFiles("webpage/recoverpassword.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	btemp.ExecuteTemplate(w,"recover_password.gohtml","Recover Password")
}else if r.Method == http.MethodPost {
	r.ParseForm()
	email := r.FormValue("email")
	newpassword := r.FormValue("newpassword")
	confirmpassword := r.FormValue("confirmpassword")
	fmt.Println(email,newpassword,confirmpassword)
	upStmt := "UPDATE `demo`.`users` SET `Password` = ? WHERE (`Email` = ?);"
	// stmt:=fmt.Sprintf(upStmt,newpassword,email)
	//fmt.Println("db.Prepare stmt:", stmt)
	//func (db *DB) Prepare(query string) (*Stmt, error)
	stmt, err := db.Prepare(upStmt)
	if err != nil {
		fmt.Println("error preparing stmt")
		panic(err)
	}
	fmt.Println("db.Prepare err:", err)
	fmt.Println("db.Prepare stmt:", stmt)
	defer stmt.Close()
	var res sql.Result
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err = stmt.Exec(newpassword,email)
	//rowsAff, _ := res.RowsAffected()
	if err != nil {
		fmt.Println("errror in rows affected")
		//temp.ExecuteTemplate(w, "result.gohtml", "There was a problem updating the product")
		return
	}else{
	fmt.Fprintln(w,  "Password was Successfully Updated")
	fmt.Println(res.RowsAffected())
}
}
}
