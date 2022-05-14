package go_build_test

import (
	"testing"

	_ "github.com/cncf/xds/go/xds/data/orca/v3"
	_ "github.com/cncf/xds/go/xds/service/orca/v3"
	_ "github.com/cncf/xds/go/xds/type/v3"
)

func TestNoop(t *testing.T) {
	// Noop test that verifies the successful importation of xDS protos
}
