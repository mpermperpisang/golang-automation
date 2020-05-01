package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*Test*/
func TestHello(t *testing.T) {
	assert := assert.New(t)
	emptyResult := Hello("")
	nameResult := Hello("Banana")

	assert.Equal(emptyResult, "Hello Dude!", "Namanya bukan Dude")
	assert.Equal(nameResult, "Hello Banana!", "Namanya bukan Banana")
}
