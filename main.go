package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcmahonliam837/microcache/config"
	"github.com/mcmahonliam837/microcache/storage"
)

var conf = config.Config{}

var store = storage.Storage{Data: make(map[string][]byte)}

func parseArgs() {
	portArg := flag.Uint("p", 3001, "The port number")
	backupArg := flag.Bool("backup", false, "Set to true inorder to keep a copy of the data to disk. This is not guaranteed to be perfectly insync with the in-memory store.")
	diskModeArg := flag.Bool("diskMode", false, "Data is stored on the disk, not in-memory. This is slow.")
	flag.Parse()

	conf.Port = *portArg
	conf.Backup = *backupArg
	conf.DiskMode = *diskModeArg
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		handlePut(w, r)
	case "GET":
		handleGet(w, r)
	case "DELETE":
		handleDelete(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	value, err := store.GetValue(key, &conf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		w.Write(value)
	}

}

func handlePut(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

    value, err := ioutil.ReadAll(r.Body)
    if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    }
	if err := store.Store(key, value, &conf); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} 
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

    if err := store.Delete(key, &conf); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
    }
}

func main() {

	parseArgs()

	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(":3001", nil); err != nil {
		fmt.Println(err)
	}
}
