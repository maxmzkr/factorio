package logistics_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/maxmzkr/factoriogo/logistics"
)

func TestNetwork(t *testing.T) {
	sb1 := logistics.NewSubnetworkBuilder()
	sb1.AddProvider(logistics.Loc{0, 0}, big.NewRat(2, 1))
	sb1.AddProvider(logistics.Loc{2, 0}, big.NewRat(2, 1))
	sb1.AddProvider(logistics.Loc{4, 0}, big.NewRat(2, 1))
	sb1.AddRequester(logistics.Loc{1, 0}, big.NewRat(3, 1))
	sb1.AddRequester(logistics.Loc{3, 0}, big.NewRat(3, 1))
	s1 := sb1.Build()

	sb2 := logistics.NewSubnetworkBuilder()
	sb2.AddProvider(logistics.Loc{0, 0}, big.NewRat(1, 1))
	sb2.AddProvider(logistics.Loc{2, 0}, big.NewRat(2, 1))
	sb2.AddProvider(logistics.Loc{4, 0}, big.NewRat(1, 1))
	sb2.AddProvider(logistics.Loc{1, 1}, big.NewRat(1, 1))
	sb2.AddProvider(logistics.Loc{3, 1}, big.NewRat(1, 1))
	sb2.AddRequester(logistics.Loc{1, 0}, big.NewRat(3, 1))
	sb2.AddRequester(logistics.Loc{3, 0}, big.NewRat(3, 1))
	s2 := sb2.Build()

	testCases := []struct {
		s  logistics.Subnetwork
		ec *big.Rat
	}{
		{
			s1,
			big.NewRat(7, 1),
		},
		{
			s2,
			big.NewRat(6, 1),
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("network%v", i), func(t *testing.T) {
			t.Parallel()
			cost := tc.s.Cost()
			if tc.ec.Cmp(cost) != 0 {
				t.Errorf("%v was not equal to %v", cost, tc.ec)
			}
		},
		)
	}
}
