package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleJoin(t *testing.T) {
	strJoin := ConsoleJoin([]Console{PC, XboxX, XboxOne})
	assert.Equal(t, "PC,XBOX_SERIES_X,XBOX_ONE", strJoin)

	strJoin = ConsoleJoin([]Console{XboxOne})
	assert.Equal(t, "XBOX_ONE", strJoin)

	strJoin = ConsoleJoin([]Console{})
	assert.Equal(t, "", strJoin)
}
