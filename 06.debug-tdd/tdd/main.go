package main

import (
	"math"
	"sort"
)

var (
	person = map[string]float64{}
)

func inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}
	person[name] = minFatRate
}

func getRand(name string) (rank int, fatRate float64) {
	person2 := map[float64][]string{}
	rankArr := make([]float64, 0)
	for nameItem, fatRateItem := range person {
		person2[fatRateItem] = append(person2[fatRateItem], nameItem)
		rankArr = append(rankArr, fatRateItem)
	}
	sort.Float64s(rankArr)
	for i, fatValue := range rankArr {
		names := person2[fatValue]
		for _, j := range names {
			if name == j {
				rank = i + 1
				fatRate = fatValue
				return rank, fatRate
			}

		}
	}
	return 0, 0
}
