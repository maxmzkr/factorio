package logistics_test

import (
	"math/big"
	"testing"

	"github.com/maxmzkr/factoriogo/pkg/logistics"
)

func TestNetwork(t *testing.T) {
	sb := logistics.NewSubnetworkBuilder()
	sb.AddProvider(logistics.Loc{0, 0}, big.NewRat(2, 1))
	sb.AddProvider(logistics.Loc{0, 0}, big.NewRat(2, 1))
	t.Error()
}
