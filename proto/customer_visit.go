package proto

import "time"

type CustomerVisit struct {
	PageViews map[string][]time.Time // PageID to slice of visit timestamps
}
