package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"All The Stars", "Kendrick Lamar,SZA", "Black Panther The Album From And Inspired By", 2018, length("3m52s")},
	{"Devil Trigger", "Ali Edwards", "Devil Trigger - Single", 2018, length("6m45s")},
	{"Remember the Name(feat. Styles of Beyond)", "Fort Minor", "The Rising Tied(Deluxe Version)", 2005, length("3m50s")},
	{"Riptide", "Vance Joy", "Dream Your Life Away(Special Edition)", 2013, length("3m22s")},
}

// function type, do the real job for Less()
type By func(t1, t2 *Track) bool

// Wraps a Track slice and its order func
type TrackSorter struct {
	tracks []*Track
	by     By
}

func (ts *TrackSorter) Len() int {
	return len(ts.tracks)
}

func (ts *TrackSorter) Less(i, j int) bool {
	return ts.by(ts.tracks[i], ts.tracks[j])
}

func (ts *TrackSorter) Swap(i, j int) {
	ts.tracks[i], ts.tracks[j] = ts.tracks[j], ts.tracks[i]
}

func (by By) Sort(tracks []*Track) {
	ts := TrackSorter{
		tracks: tracks,
		by:     by,
	}
	sort.Sort(&ts)
}

func (by By) Reverse(tracks []*Track) {
	ts := TrackSorter{
		tracks: tracks,
		by:     by,
	}
	sort.Sort(sort.Reverse(&ts))
}

// same function signature as By
func Artist(t1, t2 *Track) bool {
	return t1.Artist < t2.Artist
}

func Year(t1, t2 *Track) bool {
	return t1.Year < t2.Year
}

func Title(t1, t2 *Track) bool {
	return t1.Title < t2.Title
}

func Custom(t1, t2 *Track) bool {
	if t1.Title != t2.Title {
		return t1.Title < t2.Title
	}
	if t1.Year != t2.Year {
		return t1.Year < t2.Year
	}
	if t1.Length != t2.Length {
		return t1.Length < t2.Length
	}
	return false
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	fmt.Fprint(tw, "\n")
	tw.Flush()
}

func main() {
	By(Custom).Sort(tracks)
	printTracks(tracks)

	By(Artist).Sort(tracks)
	printTracks(tracks)

	By(Title).Sort(tracks)
	printTracks(tracks)

	By(Title).Reverse(tracks)
	printTracks(tracks)
}
