package main

import (
	
	"fmt"
	
)

type Dimension struct{
	length int
	width int 
	height int
}

func (d *Dimension) Area()int{
	d.height=40
	return d.width*d.height
}
func (d Dimension) Area2()int{
	d.height=60
	return d.width*d.height
}


func CheckArea(){
	d:= Dimension{10,20,30}
	fmt.Println(d) // here you can see the values you can set
	fmt.Println(d.Area()) //  here we pass receiver as pointer so it will change the actual values of receiver
	fmt.Println(d) //  when you see here height will be 30 => 40
	fmt.Println(d.Area2()) // but here we just pass the values so it will update value just for function not after that
	fmt.Println(d) //  so you can't see any changes here


}

func main()  {
	x:=10
	n:=&x
	fmt.Println("n = >",n,"\n *n = >",*n )
	*n=50// *n mean value of *n and n is point to the address of n/x
	fmt.Println("n = >",n,"\n *n = >",x )
	CheckArea()
	
}