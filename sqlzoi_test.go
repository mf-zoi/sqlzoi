package sqlzoi_test

import (
	"testing"

	"github.com/mf-zoi/sqlzoi"
)

func TestSample(t *testing.T) {
	if sqlzoi.Sample(1, 2) != 3 {
		t.Error("ng")
	}
}
