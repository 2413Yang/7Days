package main 

import (
	"fmt"
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w,"URL.PATH=%q\n",req.URL.Path)
	})
	r.GET("/hello",func(w http.ResponseWriter,req *http.Request){
		for k,v := range req.Header{
			fmt.Fprintf(w,"Hander[%q]%q\n",k,v)
		}
	})
	r.Run(":9999")
}
