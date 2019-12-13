package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*Test*/
func TestHello(t *testing.T) {
	emptyResult := Hello("")
	nameResult := Hello("Banana")

	assert.Equal(t, emptyResult, "Hello Dude!", "Namanya bukan Dude")
	assert.Equal(t, nameResult, "Hello Banana!", "Namanya bukan Banana")
}
