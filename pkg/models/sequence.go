package models

type Sequence struct {
	InitialNumber int64
	Steps         int64
	CurrentNumber int64
	MaxNumber     int64
}

type SeqResult struct {
	Number    int64
	Steps     int64
	MaxNumber int64
}

func (s *Sequence) Iterate(passedNumbers *CheckedNumbers, res chan Sequence) {
	s.Steps++

	if s.CurrentNumber%2 != 0 {
		s.CurrentNumber = s.CurrentNumber*3 + 1
	} else {
		s.CurrentNumber /= 2
	}

	if s.MaxNumber < s.CurrentNumber {
		s.MaxNumber = s.CurrentNumber
	}

	if s.CurrentNumber == 1 || passedNumbers.IsChecked(s.CurrentNumber) {
		res <- *s
		return
	}

	go s.Iterate(passedNumbers, res)
}
