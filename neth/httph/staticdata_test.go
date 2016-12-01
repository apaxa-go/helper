package httph

import (
	"github.com/apaxa-go/helper/mimeh"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStaticDataElement_Time(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	sde := StaticDataElement{UnixTime: now.Unix()}
	if r := sde.Time(); !r.Equal(now) {
		t.Errorf("expect %v, got %v", now, r)
	}
}

type sdeServeTest struct {
	// input
	method   string
	ifModSin time.Time

	// output
	status    int
	emptyBody bool
	emptyType bool
	emptyModT bool
}

var noTime = time.Unix(-1, -1)
var testTimeStamp = time.Date(2016, 1, 1, 12, 00, 00, 00, time.UTC)
var sdeServeTests = []sdeServeTest{
	{method: "GET", ifModSin: noTime, status: http.StatusOK, emptyBody: false, emptyType: false, emptyModT: false},
	{method: "HEAD", ifModSin: noTime, status: http.StatusOK, emptyBody: true, emptyType: true, emptyModT: false},
	{method: "POST", ifModSin: noTime, status: http.StatusMethodNotAllowed, emptyBody: true, emptyType: true, emptyModT: true},
	{method: "UNKNOWN", ifModSin: noTime, status: http.StatusNotImplemented, emptyBody: true, emptyType: true, emptyModT: true},
	{method: "GET", ifModSin: testTimeStamp.Add(-time.Hour), status: http.StatusOK, emptyBody: false, emptyType: false, emptyModT: false},
	{method: "GET", ifModSin: testTimeStamp, status: http.StatusNotModified, emptyBody: true, emptyType: true, emptyModT: false},
	{method: "GET", ifModSin: testTimeStamp.Add(time.Hour), status: http.StatusNotModified, emptyBody: true, emptyType: true, emptyModT: false},
	{method: "HEAD", ifModSin: testTimeStamp.Add(-time.Hour), status: http.StatusOK, emptyBody: true, emptyType: true, emptyModT: false},
	{method: "HEAD", ifModSin: testTimeStamp, status: http.StatusNotModified, emptyBody: true, emptyType: true, emptyModT: false},
	{method: "HEAD", ifModSin: testTimeStamp.Add(time.Hour), status: http.StatusNotModified, emptyBody: true, emptyType: true, emptyModT: false},
}

func TestStaticDataElement_ServeHTTP(t *testing.T) {
	mime := mimeh.MimeHTML
	data := "Hello test!"
	sde := StaticDataElement{
		UnixTime: testTimeStamp.Unix(),
		Mime:     mime,
		Data:     []byte(data),
	}
	url := "http://example.com/foo"

	for _, test := range sdeServeTests {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(test.method, url, nil)
		if !test.ifModSin.Equal(noTime) {
			req.Header.Add("If-Modified-Since", test.ifModSin.UTC().Format(http.TimeFormat))
		}
		sde.ServeHTTP(w, req)

		if w.Code != test.status {
			t.Errorf("expected status code is %v, but got %v", test.status, w.Code)
		}

		if test.emptyBody && w.Body.Len() != 0 {
			t.Errorf("expected empty body, but got '%v'", w.Body.String())
		}
		if !test.emptyBody && string(w.Body.Bytes()) != data {
			t.Errorf("expected body is '%v', but got '%v'", message, w.Body.String())
		}

		if r := w.Header().Get("Content-type"); test.emptyType && r != "" {
			t.Errorf("expect no mime, but got '%v'", r)
		} else if !test.emptyType && r != mime {
			t.Errorf("expected mime is '%v', but got '%v'", mime, r)
		}

		if r := w.Header().Get("Last-Modified"); test.emptyModT && r != "" {
			t.Errorf("expect no last-modified, but got '%v'", r)
		} else if r2, err := time.Parse(http.TimeFormat, r); !test.emptyModT && err != nil {
			t.Errorf("expect last-modified, but unable to parse '%v': %v", r, err)
		} else if !test.emptyModT && !r2.Equal(testTimeStamp) {
			t.Errorf("expected last-modified is '%v', but got '%v'", testTimeStamp, r2)
		}
	}
}
