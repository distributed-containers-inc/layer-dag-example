package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	resp := `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>LayerCI HTTP server example</title>
</head>
<body>
	<div style="max-width: 750px; margin: 10px auto 10px auto;">
	<h1 style="width: 100%; text-align: center">
		LayerCI HTTP server example
	</h1>
	<p>
		There are two services in this example.
	</p>
	<ol>
		<li>service-one, which has a webserver but no tests</li>
		<li>service-two, which has some unit tests to run in parallel but no webserver</li>
	</ol>
	</div>

</body>
</html>`


	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.WriteHeader(200)
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
