package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

// searchMovie makes a get request to obtain info of a movie
func searchMovie(title string) (img string, err error) {
	term := url.QueryEscape(title)
	resp, err := http.Get("https://www.omdbapi.com/?apikey=a6afe0aa&t=" + term)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("search failed: %s", resp.Status)
	}

	var movie struct{ Poster string }
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return "", err
	}

	return movie.Poster, nil
}

// Download download whatever url
func Download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("search failed: %s", resp.Status)
	}

	local := path.Base(resp.Request.URL.Path)
	f, createErr := os.Create(local)
	if createErr != nil {
		return createErr
	}

	_, copyErr := io.Copy(f, resp.Body)
	if closeErr := f.Close(); copyErr == nil {
		copyErr = closeErr
	}

	return copyErr
}

func main() {
	titlePtr := flag.String("title", "", "title of movie to search")
	flag.Parse()

	if url, err := searchMovie(*titlePtr); err != nil {
		log.Fatal(err)
	} else {
		err = Download(url)
		if err != nil {
			log.Fatal(err)
		}
	}

}
