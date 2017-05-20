package impl

import (
	"io/ioutil"
	"log"
)

type Buffers map[string][]byte

func (b Buffers) Overlay(path string) error {
	fi, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, info := range fi {
		log.Println(info.Name())
		if _, ok := b[info.Name()]; !ok && !info.IsDir() {
			bs, err := ioutil.ReadFile(info.Name())
			if err != nil {
				return err
			}
			b[info.Name()] = bs
		}
	}

	return nil
}
