// / xkcd catalog
// / build index on June 22, 2024 up to #2949 (4m real, 3s user, 2s sys)
// / xkcd #404 is missing (HTTP Status Code 404?!)
// / xkcd at around #1650 (2016 March), no more transcript

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const XkcdURL = "https://xkcd.com/%d/info.0.json"

type Comic struct {
	Num   int
	Title string
	Alt   string
	Img   string
}

var flagBuildIndex bool

func init() {
	flag.BoolVar(&flagBuildIndex, "buildindex", false, "Build Index")
}
func main() {
	flag.Parse()
	if flagBuildIndex {
		fmt.Println("buildindex mode")
		buildIndex()
	} else {
		fmt.Println("search mode")
		search()
	}
}

func search() {
}

func buildIndex() {
	for i := range 2 { // 2949 {
		comic, ok := getInfo(i)
		if ok != false {
			fmt.Printf("%d\t%s\t%s\n", comic.Num, comic.Title, comic.Alt)
		}
	}
}

// get json info for one xkcd comic
func getInfo(id int) (Comic, bool) {
	var comic Comic

	// http get
	s := fmt.Sprintf(XkcdURL, id+1)
	resp, err := http.Get(s)
	if err != nil {
		log.Fatalf("HTTP Get failed: %s", err)
	}

	// http status ok
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Printf("JSON query failed: %s", resp.Status)
		return comic, false
	}

	// json decode into go struct
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	resp.Body.Close()

	return comic, true
}
