package encore

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
)

type Entrypoint struct {
	JS []string `json:"js"`
	CSS []string `json:"css"`
}

const defaultEntrypointFile = "entrypoints.json"

type Entrypoints map[string]Entrypoint

func readEntrypoints(file fs.File) (Entrypoints, error) {
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var e struct {
		Entrypoints map[string]Entrypoint `json:"entrypoints"`
	}

	if err := json.Unmarshal(b, &e); err != nil {
		return nil, err
	}

	return e.Entrypoints, nil
}