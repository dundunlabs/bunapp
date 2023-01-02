package router

import (
	"encoding/json"

	"github.com/uptrace/bunrouter"
)

func BindQuery(req bunrouter.Request, v any) error {
	query := make(map[string]any)
	for k, v := range req.URL.Query() {
		if len(v) == 1 {
			query[k] = v[0]
		} else {
			query[k] = v
		}
	}

	b, err := json.Marshal(query)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}

func BindBody(req bunrouter.Request, v any) error {
	return json.NewDecoder(req.Body).Decode(v)
}
