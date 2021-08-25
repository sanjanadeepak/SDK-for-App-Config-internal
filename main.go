package main

import (
	"html/template"
	"log"
	"net/http"
	// AppConfiguration "github.com/IBM/appconfiguration-go-sdk/lib"
)

type data struct {
	Done          bool
	LeftNavMenu   bool
	DiscountValue string
	// UserEmail      string
	// IsLoggedInUser bool
}

func main() {
	// appConfiguration := AppConfiguration.GetInstance()
	// appConfiguration.Init("region", "guid", "apikey")
	http.HandleFunc("/", First)
	//http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/flights", Second)
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", http.FileServer(http.Dir("stylesheets"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func First(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("views/index.html", "views/partials/head.html", "views/partials/header_default.html", "views/partials/header_new.html", "views/partials/footer.html")
	if err != nil {
		panic(err)
	}
	var d1 data
	IsLoggedInUser := true
	FlightBookingAllowed := true
	if IsLoggedInUser && FlightBookingAllowed {
		d1.Done = true
	} else {
		d1.Done = false
	}
	d1.LeftNavMenu = true
	// d1.UserEmail = "alice@bluecharge.com"
	// d1.IsLoggedInUser = true
	err = t.Execute(w, d1)
	if err != nil {
		panic(err)
	}
}
func Second(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("views/flights.html", "views/partials/head.html", "views/partials/header_default.html", "views/partials/header_new.html", "views/partials/footer.html")
	if err != nil {
		panic(err)
	}
	d2 := data{DiscountValue: "30", LeftNavMenu: true}
	err = t.Execute(w, d2)
	if err != nil {
		panic(err)
	}
}
