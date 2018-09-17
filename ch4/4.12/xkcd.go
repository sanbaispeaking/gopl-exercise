package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://xkcd.com/"
const cache = "/tmp/_xkcd_"

// Comic binds the json resp of a query for
// comic json representation
type Comic struct {
	Title      string
	IMG        string
	Transcirpt string `json:"alt"`
}

func download(id int) error {
	path := fmt.Sprintf("%d/info.0.json", id)
	resp, err := http.Get(baseURL + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s failed: %s", path, resp.Status)
	}

	f, err := os.Open(fmt.Sprintf("%s/%d.json", cache, id))
	if err != nil {
		return err
	}
	defer f.Close()

	_, e := io.Copy(f, resp.Body)
	if e != nil {
		return e
	}

	return nil
}

func main() {

}

func init() {
	err := os.Mkdir(cache, 0755)
	if err != nil {
		log.Fatalln(err)
	}
}
