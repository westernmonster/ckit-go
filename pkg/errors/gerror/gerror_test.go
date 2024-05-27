package gerror_test

import (
	"ckit-go/pkg/errors/gerror"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nilError() error {
	return nil
}

func Test_Nil(t *testing.T) {
	assert.NotEqual(t, nil, gerror.New(""))
	// assert.Equal(t, gerror.Wrap(nilError(), "test"), nil)
}
