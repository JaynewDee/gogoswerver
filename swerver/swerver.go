package swerver

import "net/http"

func Launch() {
	http.ListenAndServe(":8080", nil)
}
