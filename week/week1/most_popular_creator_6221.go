package week1

import "sort"

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	res := make([][]string, 0)
	// key：人名，value：作品和热度
	statistic := make(map[string][]VideoAndHot)
	viewStat := make(map[string]int)
	max := 0
	for i, c := range creators {
		if _, exist := statistic[c]; !exist {
			statistic[c] = make([]VideoAndHot, 0)
		}
		statistic[c] = append(statistic[c], VideoAndHot{
			id:  ids[i],
			hot: views[i],
		})
		viewStat[c] += views[i]
		if max < viewStat[c] {
			max = viewStat[c]
		}
	}
	for c, hot:= range viewStat {
		if hot == max {
			sort.Slice(statistic[c], func(i, j int) bool {
				if statistic[c][i].hot == statistic[c][j].hot {
					return statistic[c][i].id < statistic[c][j].id
				}
				return statistic[c][i].hot > statistic[c][j].hot
			})
			res = append(res, []string{c, statistic[c][0].id})
		}
	}
	return res
}

type VideoAndHot struct {
	id string
	hot int
}
