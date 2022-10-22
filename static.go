package main 

import (
	// "fmt"
	"log"
	"html/template"
	"net/http"
)
// defining host and port as const 
const(
	HOST = "localhost"
	PORT = "8090"
)
// person struct to display on page
type Person struct{
	Name string
	Id int
	Refrence string
}
// loading and parsing the first template
func page1(w http.ResponseWriter, r *http.Request){
	Person1 := Person{Name:"michael", Id:40, Refrence:"304mich24"}
	template_one, _ := template.ParseFiles("templates/one.html")
	err:= template_one.Execute(w, Person1)
	if err != nil{
		log.Fatal("error loading template ", err.Error())
	} 
}
// loading and parsing the second template
func page2(w http.ResponseWriter, r *http.Request){
	Person2 := Person{Name:"lincoln", Id:101, Refrence:"linc2345"}
	template_two, _ := template.ParseFiles("templates/two.html")
	err:= template_two.Execute(w, Person2)
	if err != nil{
		log.Fatal("error loading template ", err.Error())
	}
}
// handler function with static file dir 
func Handler(mux *http.ServeMux){
	staticfile := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticfile))
	mux.HandleFunc("/one", page1)
	mux.HandleFunc("/two", page2)

}
// server function
func server(){
	mux := http.NewServeMux()
	Handler(mux)
	// listen and serve on the provided host and port 
	err := http.ListenAndServe(HOST+":"+PORT, mux)

	if err != nil{
		log.Fatal("error serving resources ", err.Error())
	}

}
func main(){
	server()
}