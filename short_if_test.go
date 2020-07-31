package funk

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestShortIf(t *testing.T) {
	is := assert.New(t)

	r := ShortIf(10>5 , 10, 5)
	is.Equal(r,10)

	r = ShortIf(10.0 == 10 , "yes", "no")
	is.Equal(r,"yes")

	r = ShortIf('a'=='b',"equal chars","unequal chars")
	is.Equal(r,"unequal chars")

	r = ShortIf("abc"=="abc","Same string","Different strings")
	is.Equal(r,"Same string")

	type testStruct struct{}
	a := testStruct{}
	b := testStruct{}
	r = ShortIf(a==b , &a, &b)
	is.Equal(r,&b)

}
