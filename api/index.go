package api

import (
	"fmt"
	"io"
	"net/http"
)

const index = "" +
	`
<head>	
<html>
	<link rel="shortcut icon" type="image/x-icon" href="./images/favicon.ico">
	<h1>Want to see a gopher, try this!</h1>
	<h1><a href="gopher">Gopher</a></h1>
</html>
</head>
`

func Index(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, index); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
}
