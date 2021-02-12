package logistics

import (
	"fmt"
	"math/big"
	"sort"
)

type logisticsNetwork struct{}

type Loc struct {
	X int64
	Y int64
}

type provider struct {
	id       int
	loc      Loc
	capacity *big.Rat
	requests map[*requester]*big.Rat
}

func (p *provider) String() string {
	return fmt.Sprintf("provider: %v", p.id)
}

func (p *provider) receiveRequest(r *requester, requestAmount *big.Rat) {
	fmt.Printf("%v received request for %v from %v\n", p, requestAmount, r)
	p.requests[r] = requestAmount
	totalRequested := new(big.Rat)
	for _, a := range p.requests {
		totalRequested.Add(totalRequested, a)
	}
	if totalRequested.Cmp(p.capacity) > 0 {
		fmt.Printf("%v needed to scale the fulfillment\n", p)
		scale := new(big.Rat)
		scale.Quo(p.capacity, totalRequested)
		for rr, rra := range p.requests {
			fulfilled := new(big.Rat)
			fulfilled.Mul(rra, scale)
			for i, rrp := range rr.providers {
				if rrp == p {
					fmt.Printf("%v fulfilled %v to %v\n", p, fulfilled, rr)
					rr.fulfillments[i] = fulfilled
					rr.updateFromIndex = i + 1
				}
			}
		}
		return
	}
	for i, rp := range r.providers {
		if rp == p {
			fmt.Printf("%v fulfilled %v to %v\n", p, requestAmount, r)
			r.fulfillments[i] = big.NewRat(requestAmount.Num().Int64(), requestAmount.Denom().Int64())
			r.updateFromIndex = i + 1
		}
	}
}

type requester struct {
	id              int
	loc             Loc
	demand          *big.Rat
	providers       []*provider
	fulfillments    []*big.Rat
	updateFromIndex int
}

func (r *requester) String() string {
	return fmt.Sprintf("requester: %v", r.id)
}

type providerDistanceSlice requester

func (pds *providerDistanceSlice) Len() int {
	return len(pds.providers)
}

func (pds *providerDistanceSlice) Less(i, j int) bool {
	iProvider := pds.providers[i]
	jProvider := pds.providers[j]
	ix := iProvider.loc.X
	iy := iProvider.loc.Y
	jx := jProvider.loc.X
	jy := jProvider.loc.Y
	if calcDistance(pds.loc, pds.providers[i].loc).Cmp(calcDistance(pds.loc, pds.providers[j].loc)) < 0 {
		fmt.Println("less by distance")
		return true
	}
	if ix < jx {
		fmt.Println("less by x")
		return true
	}
	if iy < jy {
		fmt.Println("less by y")
		return true
	}
	fmt.Println("not less")
	return false
}

func (pds *providerDistanceSlice) Swap(i, j int) {
	tmp := pds.providers[i]
	pds.providers[i] = pds.providers[j]
	pds.providers[j] = tmp
}

func (r *requester) makeRequestsFromIndex(i int) bool {
	fmt.Printf("%v making request from %v\n", r, i)
	if i == len(r.providers) {
		fmt.Printf("%v has no more providers\n", r)
		return false
	}
	for j, p := range r.providers[i : i+1] {
		totalFulfillments := new(big.Rat)
		for _, f := range r.fulfillments[:i+j] {
			fmt.Printf("%v is summing %v\n", r, f)
			totalFulfillments.Add(totalFulfillments, f)
		}
		if totalFulfillments.Cmp(r.demand) == 0 {
			fmt.Printf("%v has everything it could possibly need\n", r)
			return false
		}
		remainingDemand := new(big.Rat)
		remainingDemand.Sub(r.demand, totalFulfillments)
		p.receiveRequest(r, remainingDemand)
	}
	return true
}

func (r *requester) makeRequests() bool {
	return r.makeRequestsFromIndex(r.updateFromIndex)
}

func calcDistance(l1, l2 Loc) *big.Rat {
	return new(big.Rat).Add(new(big.Rat).Abs(new(big.Rat).Sub(big.NewRat(l1.X, 1), big.NewRat(l2.X, 1))), new(big.Rat).Abs(new(big.Rat).Sub(big.NewRat(l1.Y, 1), big.NewRat(l2.Y, 1))))
}

func (r *requester) cost() *big.Rat {
	totalCost := new(big.Rat)
	for i, f := range r.fulfillments {
		distance := calcDistance(r.loc, r.providers[i].loc)
		cost := new(big.Rat)
		cost.Mul(distance, f)
		totalCost.Add(totalCost, cost)
	}
	return totalCost
}

type Subnetwork interface {
	Cost() *big.Rat
}

type subnetwork struct {
	providers  []*provider
	requesters []*requester
}

func (s *subnetwork) Cost() *big.Rat {
	for {
		fmt.Printf("cost loop\n")
		neededUpdates := false
		for _, r := range s.requesters {
			neededUpdates = r.makeRequests() || neededUpdates
		}
		if !neededUpdates {
			break
		}
	}

	totalCost := new(big.Rat)
	for _, r := range s.requesters {
		rCost := r.cost()
		totalCost.Add(totalCost, rCost)
	}
	return totalCost
}

type SubnetworkBuilder interface {
	AddProvider(l Loc, capacity *big.Rat)
	AddRequester(l Loc, capacity *big.Rat)
	Build() Subnetwork
}

type subnetworkBuilder struct {
	providers  []*provider
	requesters []*requester
}

func NewSubnetworkBuilder() SubnetworkBuilder {
	return &subnetworkBuilder{}
}

func (sb *subnetworkBuilder) AddProvider(l Loc, capacity *big.Rat) {
	newProvider := &provider{
		id:       len(sb.providers),
		loc:      l,
		capacity: capacity,
		requests: make(map[*requester]*big.Rat),
	}
	sb.providers = append(sb.providers, newProvider)
}

func (sb *subnetworkBuilder) AddRequester(l Loc, demand *big.Rat) {
	newRequester := &requester{
		id:     len(sb.requesters),
		loc:    l,
		demand: demand,
	}
	sb.requesters = append(sb.requesters, newRequester)
}

func (sb *subnetworkBuilder) Build() Subnetwork {
	for _, r := range sb.requesters {
		r.providers = make([]*provider, len(sb.providers))
		for i := range sb.providers {
			r.providers[i] = sb.providers[i]
		}
		pds := providerDistanceSlice(*r)
		sort.Sort(&pds)
		r.fulfillments = make([]*big.Rat, len(r.providers))
		for i := range r.fulfillments {
			r.fulfillments[i] = big.NewRat(0, 1)
		}
	}
	for _, r := range sb.requesters {
		fmt.Printf("%v requests from:\n", r)
		for _, p := range r.providers {
			fmt.Printf("%v\n", p)
		}
	}
	return &subnetwork{sb.providers, sb.requesters}
}
