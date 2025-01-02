package luatool

import "github.com/harnyk/gena"

type Manifest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Params      gena.H `json:"params"`
}
