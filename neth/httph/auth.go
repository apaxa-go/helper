package httph

import "net/http"

const message = "401 Unauthorized access"

// AuthError replies to the request with an HTTP 401 not authorized error.
// r *http.Request does not used and may be nil.
func AuthError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "None")
	http.Error(w, message, http.StatusUnauthorized)
	return
}
