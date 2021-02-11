package logistics

import (
	"math/big"
	"sort"
)

type logisticsNetwork struct{}

type Loc struct {
	x int64
	y int64
}

type provider struct {
	loc      Loc
	capacity *big.Rat
	requests map[*requester]*big.Rat
}

func (p *provider) receiveRequest(r *requester, requestAmount *big.Rat) (fulfilled *big.Rat) {
	p.requests[r] = requestAmount
	totalRequested := new(big.Rat)
	for _, a := range p.requests {
		totalRequested.Add(totalRequested, a)
	}
	if totalRequested.Cmp(p.capacity) > 0 {
		scale := new(big.Rat)
		scale.Quo(p.capacity, totalRequested)
		fulfilled := new(big.Rat)
		fulfilled.Mul(requestAmount, scale)
		return fulfilled
	}
	return big.NewRat(requestAmount.Num().Int64(), requestAmount.Denom().Int64())
}

type requester struct {
	loc          Loc
	demand       *big.Rat
	providers    []provider
	fulfillments []*big.Rat
}

type providerDistanceSlice requester

func (pds *providerDistanceSlice) Len() int {
	return len(pds.providers)
}

func (pds *providerDistanceSlice) Less(i, j int) bool {
	iProvider := pds.providers[i]
	jProvider := pds.providers[j]
	ix := iProvider.loc.x
	iy := iProvider.loc.y
	jx := jProvider.loc.x
	jy := jProvider.loc.y
	distance := func(x, y int64) *big.Rat {
		return new(big.Rat).Add(new(big.Rat).Abs(new(big.Rat).Sub(big.NewRat(pds.loc.x, 1), big.NewRat(x, 1))), new(big.Rat).Abs(new(big.Rat).Sub(big.NewRat(pds.loc.y, 1), big.NewRat(y, 1))))
	}
	if distance(ix, iy).Cmp(distance(jx, jy)) < 0 {
		return true
	}
	if ix < jx {
		return true
	}
	if iy < jy {
		return true
	}
	return false
}

func (pds *providerDistanceSlice) Swap(i, j int) {
	tmp := pds.providers[i]
	pds.providers[i] = pds.providers[j]
	pds.providers[j] = tmp
}

func (r *requester) makeRequests() {
	pds := providerDistanceSlice(*r)
	sort.Sort(&pds)
	for i, p := range r.providers {
		totalFulfillments := new(big.Rat)
		for _, f := range r.fulfillments[:i] {
			totalFulfillments.Add(totalFulfillments, f)
		}
		remainingDemand := new(big.Rat)
		remainingDemand.Sub(r.demand, totalFulfillments)
		p.receiveRequest(r, remainingDemand)
	}
}

type Subnetwork interface {
	Cost() *big.Rat
}

type subnetwork struct {
	providers  []provider
	requesters []requester
}

func (s *subnetwork) Cost() *big.Rat {
	return big.NewRat(1, 1)
}

type SubnetworkBuilder interface {
	AddProvider(l Loc, capacity *big.Rat)
	AddRequester(l Loc, capacity *big.Rat)
	Build() Subnetwork
}

type subnetworkBuilder struct {
	providers  []provider
	requesters []requester
}

func NewSubnetworkBuilder() SubnetworkBuilder {
	return &subnetworkBuilder{}
}

func (sb *subnetworkBuilder) AddProvider(l Loc, capacity *big.Rat) {
	sb.providers = append(sb.providers, provider{loc: l, capacity: capacity})
}

func (sb *subnetworkBuilder) AddRequester(l Loc, demand *big.Rat) {
	sb.requesters = append(sb.requesters, requester{loc: l, demand: demand})
}

func (sb *subnetworkBuilder) Build() Subnetwork {
	for _, r := range sb.requesters {
		r.providers = sb.providers
	}
	return &subnetwork{sb.providers, sb.requesters}
}
