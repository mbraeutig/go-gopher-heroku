package api

import (
	"fmt"
	"io"
	"net/http"
)

const index = "" +
	`
<html>
    <h1>Testing heroku.com</h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	if _, err := io.WriteString(w, index); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
}
