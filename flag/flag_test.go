package astiflag_test

import (
	"os"
	"testing"

	"github.com/sgraves66/go-astitools/flag"
	"github.com/stretchr/testify/assert"
)

func TestSubcommand(t *testing.T) {
	os.Args = []string{"bite"}
	assert.Equal(t, "", astiflag.Subcommand())
	os.Args = []string{"bite", "-caca"}
	assert.Equal(t, "", astiflag.Subcommand())
	os.Args = []string{"bite", "caca"}
	assert.Equal(t, "caca", astiflag.Subcommand())
}
