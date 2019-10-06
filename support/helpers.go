package support

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// PrintError simply prints the error message to stdout
func PrintError(err error) {
	log.Printf("ERROR: %v", err.Error())
}

// ReturnString writes a string msg to the HTTP response
func ReturnString(w http.ResponseWriter, statusCode int, msg string) error {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(msg))
	return err
}

// PrettyPrintJSON will pretty print the passed object
func PrettyPrintJSON(data interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ReturnJSON will write encoded JSON onto the HTTP response
func ReturnJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(json))
	return err
}

// ReturnPrettyJSON will write encoded, prettified JSON onto the HTTP response
func ReturnPrettyJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	json, err := PrettyPrintJSON(data)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(json))
	return err
}

// GetQueryParam parses a request query params for a particular key
func GetQueryParamDefault(r *http.Request, key string, fallback string) string {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return fallback
	}

	params, ok := values[key]
	if !ok || len(params) == 0 {
		return fallback
	}

	return params[0]
}

// GetQueryParam parses a request query params for a particular key
func GetQueryParam(r *http.Request, key string) []string {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		values = make(map[string][]string)
	}

	return values[key]
}

// UnmarshalJSON takes in a request body and a pointer to a struct with which to unmarshal JSON from the body
func UnmarshalJSON(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	return decoder.Decode(v)
}
