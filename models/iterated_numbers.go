package models

import (
	"fmt"
	"sort"

	"github.com/brentp/intintmap"
)

type AllIteratedNumbers struct {
	Numbers *intintmap.Map
}

func NewAllIteratedNumbers(maxI int64) AllIteratedNumbers {
	return AllIteratedNumbers{
		Numbers: intintmap.New(int(maxI*5), 0.99), //nolint:gomnd
	}
}

func (ain *AllIteratedNumbers) AddNumbers(nn []int64) {
	for _, n := range nn {
		i, _ := ain.Numbers.Get(n)
		i++
		ain.Numbers.Put(n, i)
	}
}

func (ain *AllIteratedNumbers) InMap(n int64) bool {
	num, ok := ain.Numbers.Get(n)

	if ok && num > 0 {
		num++
		ain.Numbers.Put(n, num)
		return true
	}

	return false
}

func (ain AllIteratedNumbers) GetMostUsedNumbers(amount int) []string {
	numbersMap := make([][2]int64, 0)
	for couple := range ain.Numbers.Items() {
		numbersMap = append(numbersMap, [2]int64{couple[0], couple[1]})
	}

	sort.Slice(numbersMap, func(i, j int) bool {
		return numbersMap[i][1] > numbersMap[j][1]
	})

	mostUsed := make([]string, 0)
	for i := 0; i < amount; i++ {
		mostUsed = append(mostUsed, fmt.Sprintf("%d (%d)", numbersMap[i][0], numbersMap[i][1]))
	}

	return mostUsed
}
