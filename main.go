package main

import (
	"database/sql"
<<<<<<< HEAD
	//"encoding/json"
=======
	"encoding/json"
>>>>>>> 662a78db9e5a7cc87e8641035d013e62eff3d79b
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var btemp *template.Template

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
	http.HandleFunc("/registered/", register)
	http.HandleFunc("/authentication", authentication)
	// 	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("../asset"))))
	// 	//	http.Handle("/resources2/", http.StripPrefix("/resources2/", http.FileServer(http.Dir("../node_modules"))))
	http.ListenAndServe(":8040", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		btemp, err = template.ParseFiles("templates/base.gohtml")
		if err != nil {
			fmt.Println(err.Error())
		}

		btemp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
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
			btemp, err = template.ParseFiles("webpage/forget_password.htm")
			if err != nil {
				fmt.Println(err.Error())
			}

			btemp.Execute(w, nil)
			fmt.Fprintln(w, `Log in succesfully`)

		} else {
			fmt.Fprintln(w, `username or password is incorrect`)
		}

	}
}

// 	qs := "select FirstName,email from `users`"
// 	rows, err := db.Query(qs)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer rows.Close()
// 	var users []user
// 	for rows.Next() {
// 		var u user
// 		if err := rows.Scan(&u.Name, &u.Email); err != nil {
// 			fmt.Println(err)
// 		}
// 		users = append(users, u)
// 		//fmt.Printf("id %d name is %s\n", id, name)
// 		//	fmt.Fprintf(w, "user name:%s email:%s \n", fname, email)
// 	}

// 	//fmt.Fprintf(w, `welcome to blood bank`)
// 	btemp, err = template.ParseFiles("templates/base.gohtml")

// 	btemp.Execute(w, users)
// 	// //excel file
// 	// f:=excelize.NewFile()

// 	// 	f.SetCellValue("Sheet1","A2",users[0].Name)
// 	// 	f.SetCellValue("Sheet1","A3",users[1].Name)
// 	// 	f.SetCellValue("Sheet1","A4",users[2].Name)
// 	// 	f.SetCellValue("Sheet1","A5",users[3].Name)
// 	// 	if err := f.SaveAs("Book1.xlsx"); err != nil {
// 	// 		fmt.Println(err)
// 	// 	}

// }
func signup(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	//fmt.Fprintf(w, `welcome to blood bank`)

	//fmt.Println("signup")
	// btemp, err = template.ParseFiles("templates/signup.gohtml")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	if r.Method == http.MethodGet {
		btemp, err = template.ParseFiles("webpage/signup.htm")
		if err != nil {
			fmt.Println(err.Error())
		}

		btemp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		email := r.FormValue("email")
		pass := r.FormValue("password")
		//display in command line
		fmt.Println(name, email, pass)
		//inser into db
		qs := "INSERT INTO `users` (`Name`, `Email`, `Password`) VALUES ('%s', '%s','%s');"
		sql := fmt.Sprintf(qs, name, email, pass)
		fmt.Println(sql)
		insert, err := db.Query(sql)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Fprintln(w, `succesfully registered`)
		// fmt.Print("form Data:", r.Form)
		// fmt.Fprintln(w, `ok`)
		//json encodede text
		// bs, err := json.Marshal(r.Form)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// w.Header().Set("Content-Type", "application/json")
		// fmt.Fprintln(w, string(bs))
=======

	if r.Method == http.MethodGet {

		//fmt.Println("signup")
		btemp, err = template.ParseFiles("webpage/signup.htm")
		if err != nil {
			fmt.Println(err.Error())
		}
		btemp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()
		//fmt.Println("form:", r.Form)

		name := r.FormValue("name")
		email := r.FormValue("email")
		pass := r.FormValue("password")
		fmt.Println(name, email, pass)

		//plain text
		//fmt.Fprintln(w, "OK")
		//json encoded text
		bs, err := json.Marshal(r.Form)
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(bs))
>>>>>>> 662a78db9e5a7cc87e8641035d013e62eff3d79b
	}

}
func signin(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	if r.Method == http.MethodGet {
		btemp, err = template.ParseFiles("webpage/signin.htm")
		if err != nil {
			fmt.Println(err.Error())
		}

		btemp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
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

			fmt.Fprintln(w, `Log in succesfully`)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Fprintln(w, `username or password is incorrect`)
		}

=======

	btemp, err = template.ParseFiles("webpage/signin.htm")
	if err != nil {
		fmt.Println(err.Error())
>>>>>>> 662a78db9e5a7cc87e8641035d013e62eff3d79b
	}
}
func forgetPassword(w http.ResponseWriter, r *http.Request) {

	btemp, err = template.ParseFiles("webpage/forget_password.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	btemp.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {

	// r.ParseForm()
<<<<<<< HEAD
	// //r.PostForm()
=======
	// r.PostForm()
>>>>>>> 662a78db9e5a7cc87e8641035d013e62eff3d79b

	// name := r.PostFormValue("name")
	// // //lname := r.FormValue("lname")
	// // email := r.FormValue("email")
	// // pass := r.FormValue("password")
	// // //display in command line
	// //fmt.Println(name, email, pass)
	//  fmt.Println(name)
	// r.ParseForm()
	// for k, v := range r.Form {

	// 	fmt.Println(k,":",v)
	// }

	// 	//display as response in browser
	// 	//fmt.Fprintf(w, `succesfully registered name:%s email:%s pass:%s`, name, email, pass)
	// 	//get value via loop
	// 	// r.ParseForm()
	// 	// for k, v := range r.Form {

	// 	// 	fmt.Println(k,":",v)
	// 	// }
	// 	// fmt.Fprintln(w, `succesfully registered`)

	//inser into db
	// qs := "INSERT INTO `users` (`Name`, `Email`, `Password`) VALUES ('%s', '%s','%s');"
	// sql := fmt.Sprintf(qs, name, email, pass)
	// //fmt.Println(sql)
	// insert, err := db.Query(sql)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Fprintln(w, `succesfully registered`)
}

func authentication(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
=======

>>>>>>> 662a78db9e5a7cc87e8641035d013e62eff3d79b
	found := false
	email := r.FormValue("email")
	pass := r.FormValue("password")
	qs := "select email,password from `users`"
	rows, err := db.Query(qs)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	//var users []user
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.Email, &u.Password); err != nil {
			fmt.Println(err)
		}
		if (u.Email == email) && (u.Password == pass) {
			found = true
			break
		}
		//users = append(users, u)
		//fmt.Printf("id %d name is %s\n", id, name)
		//	fmt.Fprintf(w, "user name:%s email:%s \n", fname, email)
	}
	//display in command line
	//fmt.Println(fname,lname,email,pass)
	//display as response in browser
	//fmt.Fprintf(w, `succesfully registered name:%s email:%s pass:%s`, name, email, pass)
	//get value via loop
	// r.ParseForm()
	// for k, v := range r.Form {

	// 	fmt.Println(k,":",v)
	// }
	// fmt.Fprintln(w, `succesfully registered`)

	if found {
		// btemp, err = template.ParseFiles("templates/base.gohtml")
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		// btemp.Execute(w, nil)
		fmt.Fprintln(w, `Log in succesfully`)
	} else {
		fmt.Fprintln(w, `username or password is incorrect`)
	}

}
