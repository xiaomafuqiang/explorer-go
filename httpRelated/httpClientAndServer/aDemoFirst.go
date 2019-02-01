package httpClientAndServer

import (
	"fmt"
	"net/http"
)

func hello() {
	var Hello = func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle hello")
		fmt.Fprintf(w, "----- hello ")
	}
	var login = func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle login")
		var s = "----- login登录 "
		var bt = ([]byte)(s)
		fmt.Println(bt)
		fmt.Println(string(bt))
		fmt.Fprintf(w, string(bt))
	}

	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
