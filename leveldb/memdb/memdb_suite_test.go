package memdb

import (
	"testing"

	"github.com/FactomProject/goleveldbUp/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
