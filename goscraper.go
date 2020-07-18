// http_request_with_cookie.go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.pentesteracademy.com/course?id="
	id := os.Args[1]
	cookiearg := os.Args[2]
	url = url + id

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// set your cookie (Value)
	myCookie := &http.Cookie{
		Name:  "SACSID",
		Value: cookiearg,
	}

	request.AddCookie(myCookie)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// get the response for goquery
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	title := document.Find("h2").First().Text()
	fmt.Println(title)

	//make a folder with the title name and move into it, so files are stored there.
	os.Mkdir(title, 0777)
	os.Chdir(title)

	//video url slice
	type videodeets struct {
		videoidlist   []string
		videonamelist []string
	}

	var vd videodeets

	//get all the video URL's
	document.Find("h4").Each(func(index int, element *goquery.Selection) {

		a := element.Find("a")
		href, exists := a.Attr("href")
		if exists {
			vidid := strings.Contains(href, "video?id=")
			if vidid == true {
				href = strings.ReplaceAll(href, "video", "accounting")
				name := strings.TrimSpace(element.Text())
				videoidurl := href
				vd.videoidlist = append(vd.videoidlist, videoidurl)
				vd.videonamelist = append(vd.videonamelist, name)

			}

		}
	})

	for i := range vd.videoidlist {

		//get video url
		fmt.Println("[-] Contacting pentesteracademy.com" + vd.videoidlist[i])
		newurl := "https://www.pentesteracademy.com"
		newurl = newurl + vd.videoidlist[i]
		request2, err := http.NewRequest("GET", newurl, nil)
		if err != nil {
			log.Fatal(err)
		}
		request2.AddCookie(myCookie)
		resp2, err := client.Do(request2)
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()

		//create a file for the video
		vidname := strconv.Itoa(i) + "_"
		vidname = vidname + vd.videonamelist[i]
		vidname = vidname + ".mp4"
		vidname = strings.ReplaceAll(vidname, " ", "")
		fmt.Printf("[/] Downloading %s \n", vidname)
		outFile, err := os.Create(vidname)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		// Copy data from HTTP response to file
		_, err = io.Copy(outFile, resp2.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[+] Download of %s finished.\n\n", vidname)

	}

}
