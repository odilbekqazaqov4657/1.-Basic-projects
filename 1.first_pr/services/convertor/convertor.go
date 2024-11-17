package convertor

type Convertor struct {
}

func NewConvertor() Convertor {
	return Convertor{}
}

func (c Convertor) BinToDec(binNumber int) (decNumber int) {

	base := 1

	for binNumber > 0 {

		rem := binNumber % 10

		decNumber += rem * base
		binNumber /= 10

		base *= 2
	}
	return
}
