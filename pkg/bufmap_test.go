package impl

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferOverlay(t *testing.T) {
	asrt := assert.New(t)

	b := Buffers{"foo": []byte("barbaz")}
	b.Overlay(path.Join(os.Getenv("PWD"), "test"))

	asrt.Len(b, 2)
	asrt.NotNil(b["testmain.go"])
}

func TestBufferModified(t *testing.T) {
	asrt := assert.New(t)

	b := Buffers{"testmain.go": []byte("overlayed")}
	b.Overlay(path.Join(os.Getenv("PWD"), "test"))

	asrt.Len(b, 1)
	asrt.Equal([]byte("overlayed"), b["testmain.go"])
}
