package pkg

type NumbersStorageInterface interface {
	AddNumber(n int64)
	AddNumbers(nn []int64)
	InMap(n int64) bool
	Amount() int
}

type SequenceInterface interface {
	RunForNumber(i int64)
	CollatzSequence(sr SequenceResultInterface) SequenceResultInterface
	AlreadyCounted(sr SequenceResultInterface) bool
}

type SequenceResultInterface interface {
	CollatzIteration()
	IsEnd() bool
	GetCurrentNumber() int64
	GetNumbers() []int64
	GetMaxNumber() int64
	GetInitialNumber() int64
}
