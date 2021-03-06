package grpclb

import "google.golang.org/grpc/naming"

// NewAddUpdate new an add Update event for server registry
func NewAddUpdate(addr string, weight WeightLvl) naming.Update {
	return naming.Update{
		Op:       naming.Add,
		Addr:     addr,
		Metadata: weight,
	}
}

// NewDeleteUpdate new a delete Update event for server registry
func NewDeleteUpdate(addr string) naming.Update {
	return naming.Update{
		Op:   naming.Delete,
		Addr: addr,
	}
}
