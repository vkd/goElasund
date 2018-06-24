package dice

// SumDices - sum of some dicer
type SumDices struct {
	dd []Dicer
}

var _ Dicer = SumDices{}

// NewSumDices - return sum of rolling all dices
func NewSumDices(dices ...Dicer) SumDices {
	return SumDices{dd: dices}
}

// Roll - throw dices
func (s SumDices) Roll() int {
	var sum int
	for _, d := range s.dd {
		sum += d.Roll()
	}
	return sum
}
