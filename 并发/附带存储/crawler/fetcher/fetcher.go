package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	// as browser access ，prevent 403
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print("NewRequest is err ", err)
		return nil, fmt.Errorf("NewRequest is err %v\n", err)
	}
	// set request header
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0")
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Error: http Get, err is %v\n", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("error: status code ", response.StatusCode)
		return nil, errors.New("error: status code " + string(response.StatusCode))
	}
	// get the target website html -- all
	bodyReader := bufio.NewReader(response.Body)
	encodingType := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, encodingType.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

// check html encoding type
func determineEncoding(read *bufio.Reader) encoding.Encoding {
	bytes, err := read.Peek(1024)
	if err != nil {
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}
	encodingType, _, _ := charset.DetermineEncoding(bytes, "")
	return encodingType
}
