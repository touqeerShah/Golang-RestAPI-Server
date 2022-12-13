package practice

import (
	
	"log"
	"fmt"
	"net/http"
	
)

type Dimension struct{
	length int
	width int 
	height int
}

func (d Dimension) Area()int{
	return d.width*d.height
}


func CheckArea(){
	d:= Dimension{10,20,30}
	fmt.Println(d.Area()) // here we call funciton direct with struct just like class function
}


func Test()  {
fmt.Println("Module Pratice !");
}
func helloworld(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello world\n");
}
func Run(addr string)  {
	http.HandleFunc("/",helloworld);
	fmt.Println("Server Started !" ,addr);

	log.Fatal(http.ListenAndServe(addr,nil))
}
