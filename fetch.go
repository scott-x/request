package request

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//parse url to html string with utf8 unicode
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d\n", resp.StatusCode)
	}

	//change io.Reader to bufio.Reader
	r := bufio.NewReader(resp.Body)
	e := determinEncoding(r)
	uf8Reader := transform.NewReader(r, e.NewDecoder())
	return ioutil.ReadAll(uf8Reader)
}

//parse url to html string with utf8 unicode
//some website don't support crawler, so you need to add header to imitate using browser to visit
func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d\n", resp.StatusCode)
	}

	//change io.Reader to bufio.Reader
	r := bufio.NewReader(resp.Body)
	e := determinEncoding(r)
	uf8Reader := transform.NewReader(r, e.NewDecoder())
	return ioutil.ReadAll(uf8Reader)
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetch error:%v\n", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
