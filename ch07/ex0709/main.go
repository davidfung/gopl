// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {
	// sort.Sort(byArtist(tracks))
	http.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func list(w http.ResponseWriter, req *http.Request) {
	// endpoint /list?by=title|artist|album|year|length
	by := req.URL.Query().Get("by")
	fmt.Printf("%v by=%s\n", time.Now(), by)
	switch {
	case by == "title":
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Title < j.Title }})
	case by == "artist":
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Artist < j.Artist }})
	case by == "album":
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Album < j.Album }})
	case by == "year":
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Year < j.Year }})
	case by == "length":
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Length < j.Length }})
	default:
		sort.Sort(byCustom{tracks, func(i, j Track) bool { return i.Title < j.Title }})
	}
	printTracks(w, tracks)
}

var trackList = template.Must(template.New("tracklist").Parse(`
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
  <th><a href='{{"list?by=title"}}'>Title</a></th>
  <th><a href='{{"list?by=artist"}}'>Artist</a></th>
  <th><a href='{{"list?by=album"}}'>Album</a></th>
  <th><a href='{{"list?by=year"}}'>Year</a></th>
  <th><a href='{{"list?by=length"}}'>Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</a></td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(wr http.ResponseWriter, tracks []*Track) {
	if err := trackList.Execute(wr, tracks); err != nil {
		log.Fatal(err)
	}
}

type byCustom struct {
	tracks []*Track
	less   func(i, j Track) bool
}

func (x byCustom) Len() int           { return len(x.tracks) }
func (x byCustom) Less(i, j int) bool { return x.less(*x.tracks[i], *x.tracks[j]) }
func (x byCustom) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

/*
//!+artistoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Go          Delilah         From the Roots Up  2012  3m38s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Moby            Moby               1992  3m37s
//!-artistoutput

//!+artistrevoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
//!-artistrevoutput

//!+yearoutput
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
Go          Delilah         From the Roots Up  2012  3m38s
//!-yearoutput

//!+customout
Title       Artist          Album              Year  Length
-----       ------          -----              ----  ------
Go          Moby            Moby               1992  3m37s
Go          Delilah         From the Roots Up  2012  3m38s
Go Ahead    Alicia Keys     As I Am            2007  4m36s
Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//!-customout
*/
