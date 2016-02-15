package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func YamlFileDecode(path string, out interface{}) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, out)
	if err != nil {
		return
	}
	return
}
