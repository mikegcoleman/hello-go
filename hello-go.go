package main // import "github.com/mikegcoleman/hello-go"

import (
	"fmt"
	"net/http"
	"os"
)

const (
	port = ":8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	output := "\nServe by Container: " + name + "\n"

	fmt.Fprintf(w, output)
}

func init() {
	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(port, nil)
}

func main() {}
