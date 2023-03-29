package httpjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewRequest(method, url string, body any) (*http.Request, error) {
	p, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal body: %w", err)
	}

	return http.NewRequest(method, url, bytes.NewBuffer(p))
}
