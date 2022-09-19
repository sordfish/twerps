package main

import (
	"fmt"
	"net/http"

	"github.com/sordfish/panull"
)

func GetActiveSinks(w http.ResponseWriter, req *http.Request) {

	sinks, _ := panull.GetActiveSinks()
	for _, v := range sinks {
		fmt.Fprintf(w, "%#v\n", v)
	}

}

func main() {
	sink := panull.Sink{Name: "Virtual Output"}
	sink.SetProperty("device.description", "Virtual Output")

	if err := sink.Create(); err != nil {
		panic(err)
	}
	defer sink.Destroy()

	http.HandleFunc("/sinks", GetActiveSinks)

	http.ListenAndServe(":7780", nil)

}
