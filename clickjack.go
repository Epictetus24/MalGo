package main

import (
	"html/template"
	"net/http"
	"fmt"
	"os"
)

var x = `
<!doctype html>
<html>
 <title>Clickjacking Demo: {{.}}</title>
 <h1>Clickjacking test for site {{.}}</h1>
  <iframe src="https://{{.}}" width="1280" height="720"> </iframe>
  <br>
  <body> If webcontent is displayed above, the site is vulnerable to clickjacking </body>
</html>
`

func main() {
	if os.Args[1] == ""{
		fmt.Printf("Please pass the domain as %s domain.com\n",os.Args[0])
		os.Exit(3)
	}
	t, err := template.New("hello").Parse(x)
	if err != nil {
		fmt.Printf("Please pass the domain as %s domain.com\n",os.Args[0])
		panic(err)
	}

	url := os.Args[1]
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        	t.Execute(w, url)
    	})
    	http.ListenAndServe(":6969", nil)
}
