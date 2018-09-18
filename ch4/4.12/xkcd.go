package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

const (
	baseURL string = "https://xkcd.com/"
	cache   string = "/tmp/xkcd.com"
	latest  int    = 2047
)

// Comic binds the json resp of a query for
// comic json representation
type Comic struct {
	Title      string
	IMG        string
	Transcirpt string `json:"alt"`
}

// download JSON representation
func download(id int) error {
	p := path.Join(strconv.Itoa(id), "info.0.json")
	resp, err := http.Get(baseURL + p)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s failed: %s", p, resp.Status)
	}

	localDir := path.Join(cache, strconv.Itoa(id))
	mkErr := os.MkdirAll(localDir, 0755)
	if mkErr != nil {
		return mkErr
	}
	f, err := os.Create(path.Join(cache, p))
	if err != nil {
		return err
	}

	// Prefer copy error
	_, copyErr := io.Copy(f, resp.Body)
	if closeErr := f.Close(); copyErr == nil {
		err = closeErr
	}

	return err
}

func transcript(id int) (*Comic, error) {
	var c Comic
	local := path.Join(cache, strconv.Itoa(id), "info.0.json")

	f, err := os.Open(local)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		if err := download(id); err != nil {
			return nil, err
		}
		f, err = os.Open(local)
		if err != nil {
			return nil, err
		}
	}
	defer f.Close()

	if decodeErr := json.NewDecoder(f).Decode(&c); decodeErr != nil {
		return nil, decodeErr
	}

	return &c, nil
}

func main() {
	term := flag.Int("num", 0, "number of the comic to search")
	flag.Parse()

	c, err := transcript(*term)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%d\t%s\t%s\t%s\n", *term, c.Title, c.IMG, c.Transcirpt)
}
