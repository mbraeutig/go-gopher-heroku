package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
    <h1>Testing Heroku</h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
