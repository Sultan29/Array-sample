package main

import (
	"fmt"
	"net/http"
	"io"
	"strconv"
	"net/url"
)
var x[10]string

func main(){

	x[0] = "Fries"
	x[1] = "Cola"
	x[2] =  "Angus"
	x[3] =  "SuperStar"
	x[4] = "FamousStar"
	x[5] = "Nuggets"
	x[6] = "CrispyCurls"
	x[7] = "HandSroop Cocktail"
	x[8] = "Water"
	x[9] = "Ice"
	fmt.Println(x)
	http.HandleFunc("/getall", getall)
	http.HandleFunc("/get", get)
	http.ListenAndServe(":8080", nil)
}
func getArray() (arrr []int){
	mass := []int{11,20,40,25,24,17,24,35,40,25}
	println(mass)
	return mass
}
func getall(w http.ResponseWriter, r *http.Request){
	for r,num := range getArray(){
		io.WriteString(w,"<p> "+x[r]+" price "+strconv.Itoa(num)+"</p>")
	}
}
func get(w http.ResponseWriter, r *http.Request) {


	url := url.URL{}
	println(url.String())
	attr := r.URL.Query()
	fmt.Println(attr)
	id := attr["id"][0]
	fmt.Println("id of request " + id)
	v := id
	d:= getArray()

		if a, err := strconv.Atoi(v);
			err == nil {
				io.WriteString(w, "<p> "+x[a-1]+" price "+strconv.Itoa(d[a-1])+"</p>")
		}

}

