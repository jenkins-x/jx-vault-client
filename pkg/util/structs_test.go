package util_test

import (
	"testing"

	"github.com/jenkins-x/jx-vault-client/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestToGenericMapAndBack(t *testing.T) {
	t.Parallel()

	s := TestStruct{
		C: "Sea",
		InnerStruct: struct {
			A string
			B string
		}{
			A: "Aye",
			B: "Bee",
		},
	}

	marshalled, err := util.ToMapStringInterfaceFromStruct(s)
	assert.NoError(t, err)

	unmarshalled := TestStruct{}
	err = util.ToStructFromMapStringInterface(marshalled, &unmarshalled)

	assert.NoError(t, err)
	assert.Equal(t, s, unmarshalled)
}

type TestStruct struct {
	C           string
	InnerStruct struct {
		A string
		B string
	}
}
