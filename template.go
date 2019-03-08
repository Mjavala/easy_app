package main

import (
	"html/template"
	"net/http"
	"log"
)

type Button struct {
	Name	string
	Value	string
	IsDisabled	bool
	IsChecked	bool
	Text	string
}

type PageVariables struct {
	PageTitle		string
	PageButtons		[]Button
	Answer			string
}


func main() {
	http.HandleFunc("/", DisplayButtuons)
	http.HandleFunc("/selected", UserSelected)
	log.Fatal(http.ListenAndServe(":8000", nil)) 
}


func DisplayButtuons(w http.ResponseWriter, r *http.Request) {
// Display radio button to user

	Title := "Do you prefer cats or dogs?"
	MyButtons := []Button{
		Button{"animalselect", "cats", false, false, "Cats"},
		Button{"animalselect", "dogs", false, false, "Dogs"},
	}
	
	MyPageVariables := PageVariables{
		PageTitle:	Title,
		PageButtons:	MyButtons,
		}
		
		t, err := template.ParseFiles("select.html")
		if err != nil {
			log.Print("template parsing error: ", err)
		}
		
		err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil { // if there is an error
			log.Print("template executing error: ", err) //log it
		}

	}
	
func UserSelected(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	youranimal := r.Form.Get("animalselect")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer : youranimal,
		}

 // generate page by passing page variables into template
		t, err := template.ParseFiles("select.html") //parse the html file homepage.html
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it
		}

		err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
		}
}