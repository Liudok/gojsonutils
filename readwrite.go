package jsonutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// JSONResponse is a type used for sending JSON around
type JSONResponse struct {
	Error bool     `json:"error"`
	Message string `json:"message"`
	Data any       `json:"data,omitempty"`
}

// Type used to instantiate the module. Any variable of this type
// will have access to all the methods with the receiver *JsonRW
type JSONUtils struct{
	MaxSize int64
	AllowUnknownFields bool
}

func (ju *JSONUtils) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	
	maxBytes := int64(1048576) // 1 Mb
	
	if ju.MaxSize != 0 {
		maxBytes = ju.MaxSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	dec := json.NewDecoder(r.Body)

	if !ju.AllowUnknownFields {
		dec.DisallowUnknownFields()
	}

	err := dec.Decode(data)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshaleTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
			case errors.As(err, &syntaxError):
				return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

			case errors.Is(err, io.ErrUnexpectedEOF):
				return errors.New("body contains badly-formated JSON")

			case errors.As(err, &unmarshaleTypeError):
				if unmarshaleTypeError.Field != "" {
					return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshaleTypeError.Field)
				}
				return fmt.Errorf("body contains incorrect JSON type at character %d", unmarshaleTypeError.Offset)

			case strings.HasPrefix(err.Error(), "json: unknown field"):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
				return fmt.Errorf("body contains unknown key %s", fieldName)

			case err.Error() == "http: request body too large":
				return fmt.Errorf("body mast not be larger than %d bytes", maxBytes)
			
			case errors.As(err, &invalidUnmarshalError):
				return fmt.Errorf("error unmarshalling JSON: %s", err.Error())
		
			default:
				return err

		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

func (ju *JSONUtils) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header ) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (ju *JSONUtils) ErrorJSON(w http.ResponseWriter, err error, status ...int ) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()

	return ju.WriteJSON(w, statusCode, payload)
}
