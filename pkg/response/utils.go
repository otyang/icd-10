package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

func ParseBody(c *fiber.Ctx, payload any) error {
	err := c.BodyParser(payload)
	err = DetailedJsonError(err)
	return BadRequest(err.Error(), nil)
}

// DetailedJsonError checks the type of json Error
func DetailedJsonError(err error) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	if err != nil {
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		default:
			return err
		}
	}
	return nil
}

// JSON is for default Golang writer.
// it marshals 'response'struct to JSON, escapes HTML & sets the
// Content-Type as application/json.
func JSON(w http.ResponseWriter, rsp *Response, headers map[string]string) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(rsp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for index, value := range headers {
		w.Header().Add(index, value)
	}
	w.WriteHeader(rsp.StatusCode)
	w.Write(buf.Bytes())
}

func getRuntimeCallerInfo(pc uintptr, file string, lineNo int, ok bool) string {
	if ok {
		details := runtime.FuncForPC(pc)
		if details != nil {
			return fmt.Sprintf("%s#%d (%s)", file, lineNo, details.Name())
		}

		return fmt.Sprintf("%s#%d", file, lineNo)
	}

	return "runtime.Caller() info not available"
}
