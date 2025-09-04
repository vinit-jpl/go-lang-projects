package main

import "net/http"

func hanndlerErr(w http.ResponseWriter, r *http.Request) {
	responWithError(w, 400, "Something went wrong")
}
