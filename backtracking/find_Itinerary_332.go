package backtracking

import "sort"

func findItinerary(tickets [][]string) []string {
	resInFindItinerary = make([]string, 0)
	record := make(map[string][]PairInFindItinerary)
	for i := 0; i < len(tickets); i++ {
		if _, exist := record[tickets[i][0]]; !exist {
			record[tickets[i][0]] = make([]PairInFindItinerary, 0)
		}
		record[tickets[i][0]] = append(record[tickets[i][0]], PairInFindItinerary{
			Des:     tickets[i][1],
			Visited: false,
		})
	}
	for from := range record {
		sort.Slice(record[from], func(i, j int) bool {
			return record[from][i].Des < record[from][j].Des
		})
	}
	endNum := len(tickets) + 1
	resInFindItinerary = append(resInFindItinerary, "JFK")
	backTrackingInFindItinerary(record, "JFK", endNum)
	return resInFindItinerary
}

var resInFindItinerary []string
func backTrackingInFindItinerary(record map[string][]PairInFindItinerary, from string, endNum int) bool {
	if len(resInFindItinerary) == endNum {
		return true
	}
	next := record[from]
	for i := 0; i < len(next); i++ {
		if !next[i].Visited {
			next[i].Visited = true
			resInFindItinerary = append(resInFindItinerary, next[i].Des)
			if backTrackingInFindItinerary(record, next[i].Des, endNum) {
				return true
			}
			resInFindItinerary = resInFindItinerary[:len(resInFindItinerary) - 1]
			next[i].Visited = false
		}
	}
	return false
}

type PairInFindItinerary struct {
	Des string
	Visited bool
}