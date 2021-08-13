package inmemory

import (
	"github.com/brentp/intintmap"
)

type AllIteratedNumbers struct {
	Numbers *intintmap.Map
}

func NewAllIteratedNumbers(maxI int64) *AllIteratedNumbers {
	return &AllIteratedNumbers{
		Numbers: intintmap.New(int(maxI*10), 0.99), //nolint:gomnd
	}
}

func (ain *AllIteratedNumbers) AddNumber(n int64) {
	i, _ := ain.Numbers.Get(n)
	i++
	ain.Numbers.Put(n, i)
}

func (ain *AllIteratedNumbers) AddNumbers(nn []int64) {
	for _, n := range nn {
		ain.AddNumber(n)
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

func (ain AllIteratedNumbers) Amount() int {
	return ain.Numbers.Size()
}
