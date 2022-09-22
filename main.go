package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sordfish/panull"
)

func GetActiveSinks(w http.ResponseWriter, req *http.Request) {

	sinks, _ := panull.GetActiveSinks()
	for _, v := range sinks {

		u, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(u))
	}

}

func GetActiveSources(w http.ResponseWriter, req *http.Request) {

	sinks, _ := panull.GetActiveSources()
	for _, v := range sinks {

		u, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(u))
	}

}

func createNullSink(w http.ResponseWriter, req *http.Request) {
	sink := panull.Sink{Name: "Virtual Output"}
	sink.SetProperty("device.description", "Virtual Output")

	if err := sink.Create(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Done!")
}

func main() {
	http.HandleFunc("/sinks", GetActiveSinks)
	http.HandleFunc("/sources", GetActiveSources)
	http.HandleFunc("/create/nullsink", createNullSink)

	http.ListenAndServe(":7780", nil)

}
