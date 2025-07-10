package grabinterview

import (
	"sort"
	"strings"
	"time"
)

var (
	logEntries = []string{
		"2023-01-01 10:00:00, C1, P1",
		"2023-01-01 11:00:00, C1, P1", // Same page within 2 days
		"2023-01-02 09:00:00, C1, P2",
		"2023-01-03 10:00:00, C1, P1", // Outside 2-day window
	}
)

type CustomerVisit struct {
	PageViews map[string][]time.Time // PageID to slice of visit timestamps
}

func ReadData() map[string]*CustomerVisit {
	customerData := make(map[string]*CustomerVisit)
	for _, entry := range logEntries {
		parts := strings.Split(entry, ", ")
		timestamp, _ := time.Parse("2006-01-02 15:04:05", parts[0])
		customerId, pageId := parts[1], parts[2]

		if customerData[customerId] == nil {
			customerData[customerId] = &CustomerVisit{
				PageViews: make(map[string][]time.Time),
			}
		}
		customerData[customerId].PageViews[pageId] = append(customerData[customerId].PageViews[pageId], timestamp)
	}
	return customerData
}

func FindRepeatVisitor(data map[string]*CustomerVisit, durationInDay int, minCount int) []string {
	var result []string
	days := time.Duration(durationInDay) * 24 * time.Hour

	for customerId, CustomerVisit := range data {
		for _, timestamps := range CustomerVisit.PageViews {
			sort.Slice(timestamps, func(i, j int) bool {
				return timestamps[i].Before(timestamps[j])
			})
			for i := 0; i <= len(timestamps)-minCount; i++ {
				if timestamps[i+minCount-1].Sub(timestamps[i]) <= days {
					result = append(result, customerId)
					break
				}
			}
		}
	}
	// 2 pointers solutions
	for customerID, customerVisit := range data {
		for _, timestamps := range customerVisit.PageViews {
			if len(timestamps) < minCount {
				continue
			}
			left := 0
			for right := 0; right < len(timestamps); right++ {
				for timestamps[right].Sub(timestamps[left]) > days {
					left++
				}
				if right-left+1 >= minCount {
					result = append(result, customerID)
					break
				}
			}
		}
	}
	return result
}

// isit 2 days and 2 unique page id
func FindLoyaltyCustomer(data map[string]*CustomerVisit, days int, minPages int) []string {
	var result []string
	for customerID, customerVisit := range data {
		uniquePage := make(map[string]struct{})
		uniqueDates := make(map[string]struct{})
		for pageID, timestamps := range customerVisit.PageViews {
			uniquePage[pageID] = struct{}{}
			for _, t := range timestamps {
				date := t.Format("2006-01-02")
				uniqueDates[date] = struct{}{}
			}
			if len(uniquePage) > minPages && len(uniqueDates) > days {
				result = append(result, customerID)
			}
		}
	}
	return result
}
