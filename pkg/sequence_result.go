package pkg

type SequenceResult struct {
	InitialNumber int64
	Steps         int
	CurrentNumber int64
	MaxNumber     int64
	Numbers       []int64
}

func (sr *SequenceResult) CollatzIteration() {
	sr.Steps++

	if sr.CurrentNumber%2 != 0 {
		sr.CurrentNumber = sr.CurrentNumber*3 + 1 //nolint: gomnd
	} else {
		sr.CurrentNumber /= 2
	}

	if sr.MaxNumber < sr.CurrentNumber {
		sr.MaxNumber = sr.CurrentNumber
	}

	sr.Numbers = append(sr.Numbers, sr.CurrentNumber)
}

func (sr SequenceResult) IsEnd() bool {
	return sr.CurrentNumber == 1
}

func (sr SequenceResult) GetCurrentNumber() int64 {
	return sr.CurrentNumber
}

func (sr SequenceResult) GetNumbers() []int64 {
	return sr.Numbers
}

func (sr SequenceResult) GetMaxNumber() int64 {
	return sr.MaxNumber
}

func (sr SequenceResult) GetInitialNumber() int64 {
	return sr.InitialNumber
}
