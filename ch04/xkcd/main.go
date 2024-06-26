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

const CSV = `.\xkcd.idx`

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
		fmt.Println("lookup mode")
		lookup()
	}
}

func lookup() {
	keyterm := flag.Arg(0)
	fmt.Printf("Looking up: %s\n", keyterm)
	comics := load(CSV)
	matches := search(comics, keyterm)
	results := report(matches)
	display(results)
}

// Load the csv into an in-memory array of Comic.
func load(csv string) []Comic {
	return nil
}

// search an array of Comic by a key term.
func search(comics []Comic, keyterm string) []Comic {
	return nil
}

// Compose a report from the matches.
func report(comics []Comic) string {
	return ""
}

// Display the search report.
func display(results string) {

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
