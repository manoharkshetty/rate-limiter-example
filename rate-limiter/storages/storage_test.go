package storages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetStorage(t *testing.T) {
	assert.NotNil(t, GetStorage())
}
