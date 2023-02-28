package srverrors

import (
	"fmt"
	"net/http"
)

const (
	ConversionError = "conversion error"
	RecordNotFound  = "record not found"
	Duplicate       = "duplicate"
	BadQueryParam   = "malformed query parameters"
)

func HandleError(err error, w http.ResponseWriter) {
	errstr := err.Error()
	if errstr == ConversionError {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, ConversionError)
	} else if errstr == RecordNotFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, RecordNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
	}
}
