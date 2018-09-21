package main

import (
	"fmt"
	"os/exec"
)

type pkg struct {
	ImportPath string `json:"ImportPath"`
	Deps       string `json:"Deps"`
}

func list(terms ...string) ([]byte, error) {
	if len(terms) == 0 {
		terms = []string{"..."}
	}
	cmd := exec.Command("go", append([]string{"list", "-e", "-f", `{"ImportPath": {{.ImportPath|printf "%q"}}, "Deps": {{.Deps|printf "%q"}}},`}, terms...)...)
	out, err := cmd.Output()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("%s", err.Stderr)
		}
		return nil, err
	}
	return out, nil
}

func main() {
	// allPackages, err := list("net")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(allPackages))
	fmt.Printf("%q", []string{"shit", "fuck"})
	// var p pkg
	// decodeErr := json.Unmarshal(append(append([]byte("["), allPackages...), []byte("]")...), &p)
	// if decodeErr != nil {
	// 	log.Fatalln(decodeErr)
	// }
	// fmt.Println(p)
}
