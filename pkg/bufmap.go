package impl

import (
	"io/ioutil"
	"log"
	"path"
)

type Buffers map[string][]byte

func (b Buffers) Overlay(p string) error {
	fi, err := ioutil.ReadDir(p)
	if err != nil {
		return err
	}

	for _, info := range fi {
		log.Println(info.Name())
		if _, ok := b[info.Name()]; !ok && !info.IsDir() {
			bs, err := ioutil.ReadFile(path.Join(p, info.Name()))
			if err != nil {
				return err
			}
			b[info.Name()] = bs
		}
	}

	return nil
}
