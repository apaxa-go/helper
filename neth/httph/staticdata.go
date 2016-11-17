package httph

import (
	"net/http"
	"time"
)

// StaticDataElement used to store arbitrary file in HTTP server application.
type StaticDataElement struct {
	UnixTime int64  // Store original file modification time as UNIX time (warning: http.TimeFormat doesn't use anything less then second in "If-Modified-Since" HTTP requests)
	Mime     string // Store mime type
	Data     []byte // Binary data (file content)
}

// Time return original file modification time as time.Time
func (s *StaticDataElement) Time() time.Time {
	return time.Unix(s.UnixTime, 0)
}

// ServeHTTP implement http.Handler interface for *StaticDataElement.
// ServeHTTP implement GET & HEAD HTTP methods.
// On POST method it answers with 405 - MethodNotAllowed.
// On other methods it answers with 501 - NotImplemented.
// ServeHTTP set "Last-Modified" header on GET & HEAD methods and use "If-Modified-Since" header on GET method to send payload only if needed.
// Happened errors just ignored.
// So usually it is enough to just register StaticDataElement using http.ServeMux.Handle.
func (s *StaticDataElement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fallthrough
	case "HEAD":
		w.Header().Set("Last-Modified", s.Time().UTC().Format(http.TimeFormat))
		if ims := r.Header.Get("If-Modified-Since"); ims != "" {
			if t, e := time.Parse(http.TimeFormat, ims); e == nil && !t.Before(s.Time()) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
		if r.Method == "GET" {
			w.Header().Set("Content-type", s.Mime)
			_, _ = w.Write(s.Data)
		}
		return
	case "POST":
		w.Header().Set("Allow", "Get")
		w.Header().Add("Allow", "Head")
		w.WriteHeader(http.StatusMethodNotAllowed)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	return
}
