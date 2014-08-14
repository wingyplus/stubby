package stubby

import (
	"gopkg.in/yaml.v1"
	"io"
	"io/ioutil"
)

func Parse(reader io.Reader) ([]Stub, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var stubs []Stub
	err = yaml.Unmarshal(b, &stubs)

	return stubs, err 
}
