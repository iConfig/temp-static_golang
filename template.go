package main 
import (
	"log"
	"net/http"
	"html/template"
)

// person struct to display on page
type Person struct {  
	Name string
	Id int
	Refrence string
}
// defining host and port as const 
const( 
	HOST = "localhost"
	PORT = "8090"
)
// loading and parsing the first template
func templateOne(w http.ResponseWriter, r *http.Request){ 

// providing value for the person struct
	person := Person {Name: "james", Id:111, Refrence:"jm4949"} 
	// loading template from dir
	template_load1, _ := template.ParseFiles("templates/one.html")
	err := template_load1.Execute(w, person)
	if err != nil{
		log.Fatal("error parsing template ", err)
	}
}
// loading and parsing the second template
func templateTwo(w http.ResponseWriter, r *http.Request){ 

// providing value for the person struct
	person2 := Person{Name:"john", Id:222, Refrence:"jh3980"} 
	// loading template from dir
	template_load2 , _ := template.ParseFiles("templates/two.html")
	err := template_load2.Execute(w, person2)
	if err != nil{
		log.Fatal("error parsing template ", err)
	}

}
// handler function 
func handler(mux *http.ServeMux){ 
	mux.HandleFunc("/one", templateOne)
	mux.HandleFunc("/two", templateTwo)
}
// server function
func server(){ 
	mux := http.NewServeMux()
	handler(mux)

	// listen and serve on the provided host and port 
	err := http.ListenAndServe(HOST+":"+PORT, mux) 
	if err != nil{
		log.Fatal("error loading server ", err)
	}

}
func main(){
	server()
}