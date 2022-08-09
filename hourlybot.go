package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/karrick/godirwalk"
)

func dirwalk(osDirname string) []string {

	var entries []string

	godirwalk.Walk(osDirname, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				return nil
			}

			entries = append(entries, filepath.FromSlash(osPathname))
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})

	return entries
}

var media []string

func task() {

	var f string

	f, media = media[0], media[1:]

	fmt.Println("F:", f)
}

func main() {

	dir := "./media"
	media = dirwalk(dir)

	s1 = gocron.NewScheduler(time.FixedZone("Asia/Sakhalin"))

}
