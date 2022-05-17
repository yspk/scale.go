package types_test

import (
	"github.com/yspk/scale.go/types"
	"testing"
)

func TestRegCustomTypes(t *testing.T) {
	types.RuntimeType{}.Reg()
}
