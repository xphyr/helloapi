/*
Copyright 2018 Mark DeNeve.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var (
	buf           bytes.Buffer
	myHostname, _ = os.Hostname()
	myOS          = runtime.GOOS
)

func init() {
	flag.Parse()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<html>`)) // Begin html block
	w.Write([]byte(`<head> 
		<title>Simple Hello World API</title>
		<style>
		label{
		display:inline-block;
		width:75px;
		}
		form label {
		margin: 10px;
		}
		form input {
		margin: 10px;
		}
		</style>
		</head><body>`)) // Head  Begin Body
	w.Write([]byte(fmt.Sprintf("<h3>Hello from: %s operating system</h3>", myOS)))
	w.Write([]byte(fmt.Sprintf("<h3>My hostname: %s</h3>", myHostname)))
	w.Write([]byte(`</body></html>`)) //END

}

func main() {

	fmt.Println("Starting App")
	http.HandleFunc("/", rootHandler)
	listenPort := ":8080"
	fmt.Printf("Listening on port: %v\n", listenPort)
	http.ListenAndServe(listenPort, nil)
}
