package inmemory

import (
	"github.com/smgladkovskiy/3nplus1/pkg"
)

type Sequence struct {
	numbersStorage pkg.NumbersStorageInterface
	resultsChan    chan pkg.SequenceResultInterface
}

func NewSequence(iNumbers pkg.NumbersStorageInterface, res chan pkg.SequenceResultInterface) pkg.SequenceInterface {
	return &Sequence{
		numbersStorage: iNumbers,
		resultsChan:    res,
	}
}

func (s *Sequence) RunForNumber(i int64) {
	sr := &pkg.SequenceResult{
		InitialNumber: i,
		CurrentNumber: i,
		Numbers:       []int64{i},
	}
	s.resultsChan <- s.CollatzSequence(sr)
}

func (s *Sequence) CollatzSequence(sr pkg.SequenceResultInterface) pkg.SequenceResultInterface {
	sr.CollatzIteration()

	if sr.IsEnd() || s.AlreadyCounted(sr) {
		return sr
	}

	return s.CollatzSequence(sr)
}

func (s *Sequence) AlreadyCounted(sr pkg.SequenceResultInterface) bool {
	return s.numbersStorage.InMap(sr.GetCurrentNumber())
}
