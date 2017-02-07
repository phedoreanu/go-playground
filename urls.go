package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	l := "https://www.google.co.uk/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	var r io.ReadCloser

	ok, err := regexp.MatchString("^https?://", l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok)
	if ok {
		resp, err := http.Get(l)
		if err != nil {
			fmt.Println(err)
		}
		r = resp.Body
	} else {
		if r, err = os.Open(l); err != nil {
			fmt.Println(err)
		}
	}

	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	}
	defer r.Close()
}
