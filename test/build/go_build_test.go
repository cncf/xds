package go_build_test

import (
	"testing"

	_ "github.com/cncf/xds/go/udpa/data/orca/v1"
	_ "github.com/cncf/xds/go/udpa/service/orca/v1"
	_ "github.com/cncf/xds/go/udpa/type/v1"
)

func TestNoop(t *testing.T) {
	// Noop test that verifies the successful importation of xDS protos
}
