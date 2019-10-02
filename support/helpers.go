package support

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// ReturnString writes a string msg to the HTTP response
func ReturnString(w http.ResponseWriter, statusCode int, msg string) error {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(msg))
	return err
}

// ReturnPrettyJSON will write encoded, prettified JSON onto the HTTP response
func ReturnPrettyJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.WriteHeader(statusCode)
	s := json.NewEncoder(w)
	s.SetIndent("", "    ")
	return s.Encode(data)
}

// GetQueryParam parses a request query params for a particular key
func GetQueryParam(r *http.Request, key string) []string {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		values = make(map[string][]string)
	}

	return values[key]
}
