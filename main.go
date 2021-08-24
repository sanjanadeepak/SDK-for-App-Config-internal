package main

import (
	"html/template"
	"log"
	"net/http"
)

type data struct {
	Done bool
}

func HomePage() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	// http.HandleFunc("/evaluate", Calculate)
	http.HandleFunc("/", First)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// func Calculate(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	exp := r.Form.Get("exp")
// 	res, err := calculator.CalculatorResult(string(exp))
// 	if err != nil {
// 		fmt.Fprintf(w, "FAILED: "+err.Error())
// 		return
// 	}
// 	t, err2 := template.ParseFiles("static/result.html")
// 	if err2 != nil {
// 		panic(err)
// 	}
// 	var d1 data
// 	if err == nil {
// 		d1 = data{Exp: string(exp), Res: fmt.Sprintf("%f", res)}
// 	} else {
// 		d1 = data{Exp: string(exp), Res: err.Error()}
// 	}

// 	err = t.Execute(w, d1)
// 	// http.ServeFile(w, r, "static/result.html")
// }
func First(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("views/index.html", "views/partials/head.html")
	if err != nil {
		panic(err)
	}
	var d1 data
	isLoggedInUser := true
	flightBookingAllowed := true
	if isLoggedInUser && flightBookingAllowed {
		d1.Done = true
	} else {
		d1.Done = false
	}

	err = t.Execute(w, d1)
	if err != nil {
		panic(err)
	}
}
func main() {
	HomePage()

}
