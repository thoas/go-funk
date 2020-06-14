package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetStructOneLevel(t *testing.T) {
	is := assert.New(t)

	// copy here because we need to modify
	assignFoo := *foo

	err := Set(&assignFoo, 2, "ID")
	is.NoError(err)
	is.Equal(2, assignFoo.ID)

}

func TestSetStructTwoLevels(t *testing.T) {
	is := assert.New(t)

	// copy here because we need to modify
	assignFoo := *foo

	err := Set(&assignFoo, int64(2), "EmptyValue.Int64")
	is.NoError(err)
	is.Equal(int64(2), assignFoo.EmptyValue.Int64)

}
