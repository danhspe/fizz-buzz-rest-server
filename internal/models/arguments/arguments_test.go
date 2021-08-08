package arguments_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danhspe/fizz-buzz-rest-server/internal/models/arguments"
)

var args = arguments.Arguments{
	Int1:  0,
	Int2:  0,
	Limit: 0,
	Str1:  "",
	Str2:  "",
}

func TestNew(t *testing.T) {
	assert.EqualValues(t, args, arguments.New(0, 0, 0, "", ""))
}

func TestNewFromJson(t *testing.T) {
	argsAsJson, err := args.AsJson()

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, argsAsJson, "argsAsJson should not be nil")

	argsFromJson, err := arguments.NewFromJson(argsAsJson)

	assert.Nil(t, err, "Error should be nil")
	assert.EqualValuesf(t, args, argsFromJson, "args do not match: expected %+v - got %+v", args, argsFromJson)
}

func TestNewFromJson_WhenJsonDecodingFails_ThenReturnError(t *testing.T) {
	_, err := arguments.NewFromJson("")
	assert.NotNil(t, err, "Expected error")
}
