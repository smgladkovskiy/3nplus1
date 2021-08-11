package models

type Sequence struct {
	InitialNumber int64
	Steps         int
	CurrentNumber int64
	MaxNumber     int64
	Numbers       []int64
}

type SeqResult struct {
	Number    int64
	Steps     int
	MaxNumber int64
}

func (s *Sequence) Iterate(iteratedNumbers *AllIteratedNumbers, res chan Sequence) {
	s.Steps++

	if s.CurrentNumber%2 != 0 {
		s.CurrentNumber = s.CurrentNumber*3 + 1 //nolint: gomnd
	} else {
		s.CurrentNumber /= 2
	}

	if s.MaxNumber < s.CurrentNumber {
		s.MaxNumber = s.CurrentNumber
	}

	if s.CurrentNumber == 1 || iteratedNumbers.InMap(s.CurrentNumber) {
		res <- *s

		return
	}

	s.Numbers = append(s.Numbers, s.CurrentNumber)

	go s.Iterate(iteratedNumbers, res)
}
