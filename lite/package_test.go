package testify

import (
	"testing"

	"github.com/stretchr/testify/lite/assert"
)

func TestImports(t *testing.T) {
	if assert.Equal(t, 1, 1) != true {
		t.Error("Something is wrong.")
	}
}
