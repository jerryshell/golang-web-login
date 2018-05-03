package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"jerryshell.cn/login_demo/dao"
	"jerryshell.cn/login_demo/domain"
	"jerryshell.cn/login_demo/session"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userinfo", userinfo)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/register", register)
}

func main() {
	log.Println("Server is running at http://localhost:8080/. Press Ctrl+C to stop.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	user, _ := session.GetSession(w, r).GetAttr("user")

	t, err := template.ParseFiles("html/index.html")
	checkError(err)

	err = t.Execute(w, user)
	checkError(err)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginHTML, err := ioutil.ReadFile("html/login.html")
		checkError(err)
		w.Write(loginHTML)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println("login", username, password)
	if isEmpty(username, password) {
		message(w, r, "字段不能为空")
		return
	}

	user := dao.FindUserByUsernameAndPassword(username, password)
	if user == nil {
		message(w, r, "登录失败！")
		return
	}
	// 登陆成功
	sess := session.GetSession(w, r)
	sess.SetAttr("user", user)
	http.Redirect(w, r, "/userinfo", 302)
}

func message(w http.ResponseWriter, r *http.Request, message string) {
	t, err := template.ParseFiles("html/message.html")
	checkError(err)

	err = t.Execute(w, map[string]string{"Message": message})
	checkError(err)
}

func userinfo(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	user, exist := sess.GetAttr("user")
	if !exist {
		http.Redirect(w, r, "/", 302)
		return
	}
	t, err := template.ParseFiles("html/userinfo.html")
	checkError(err)
	t.Execute(w, user)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		registerHTML, err := ioutil.ReadFile("html/register.html")
		checkError(err)
		w.Write(registerHTML)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	email := r.FormValue("email")

	if isEmpty(username, password, password2, email) {
		message(w, r, "字段不能为空")
		return
	}

	if password != password2 {
		message(w, r, "两次密码不相符")
		return
	}

	user := &domain.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	dao.AddUser(user)
	message(w, r, "注册成功！")
}

func logout(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	sess.DelAttr("user")
	http.Redirect(w, r, "/", 302)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isEmpty(strs ...string) (isEmpty bool) {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" || len(str) == 0 {
			isEmpty = true
			return
		}
	}
	isEmpty = false
	return
}
