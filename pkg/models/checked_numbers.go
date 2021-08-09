package models

import (
	"sync"
)

type CheckedNumbers struct {
	tenPower2  map[int64]int
	tenPower4  map[int64]int
	tenPower6  map[int64]int
	tenPower8  map[int64]int
	tenPower10 map[int64]int
	tenPower12 map[int64]int
	m          sync.RWMutex
}

func NewCheckedNumbers() CheckedNumbers {
	return CheckedNumbers{
		tenPower2:  make(map[int64]int),
		tenPower4:  make(map[int64]int),
		tenPower6:  make(map[int64]int),
		tenPower8:  make(map[int64]int),
		tenPower10: make(map[int64]int),
		tenPower12: make(map[int64]int),
	}
}

func (cn *CheckedNumbers) IsChecked(n int64) bool {
	switch {
	case n >= 1 && n < 10e2:
		cn.m.RLock()
		if _, ok := cn.tenPower2[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	case n >= 10e2 && n < 10e4:
		cn.m.RLock()
		if _, ok := cn.tenPower4[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	case n >= 10e4 && n < 10e6:
		cn.m.RLock()
		if _, ok := cn.tenPower6[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	case n >= 10e6 && n < 10e8:
		cn.m.RLock()
		if _, ok := cn.tenPower8[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	case n >= 10e8 && n < 10e10:
		cn.m.RLock()
		if _, ok := cn.tenPower10[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	case n >= 10e10 && n < 10e12:
		cn.m.RLock()
		if _, ok := cn.tenPower12[n]; ok {
			cn.m.RUnlock()
			return true
		}
		cn.m.RUnlock()
	}

	return false
}

func (cn *CheckedNumbers) Set(n int64) {
	cn.m.Lock()
	switch {
	case n >= 1 && n < 10e2:
		cn.tenPower2[n] = 1
	case n >= 10e2 && n < 10e4:
		cn.tenPower4[n] = 1
	case n >= 10e4 && n < 10e6:
		cn.tenPower6[n] = 1
	case n >= 10e6 && n < 10e8:
		cn.tenPower8[n] = 1
	case n >= 10e8 && n < 10e10:
		cn.tenPower10[n] = 1
	case n >= 10e10 && n < 10e12:
		cn.tenPower12[n] = 1
	}
	cn.m.Unlock()
}
