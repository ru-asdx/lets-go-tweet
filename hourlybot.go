package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/karrick/godirwalk"
)

var media_loaded int
var media []string

func loadMedia(osDirname string) ([]string, int) {

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

	return entries, len(entries)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func task(dir string) {
	var fn string

	if media_loaded == 0 || len(media) == 0 {
		media, media_loaded = loadMedia(dir)

		if media_loaded < 2 {
			log.Fatal("Not enought media files.")
		}
	}

	log.Print("F: ", fn, ", media_loaded: ", media_loaded, ", media_current: ", len(media))
	fn, media = media[0], media[1:]

}

func main() {

	mediaDir := getEnv("MEDIA_DIR", "./media")
	timezone := getEnv("TZ", "Asia/Sakhalin")

	location, _ := time.LoadLocation(timezone)
	s1 := gocron.NewScheduler(location)

	s1.Every(3).Seconds().Do(task, mediaDir)
	s1.StartBlocking()
}
