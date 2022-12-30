package lib

type History []float64

func (h *History) Saldo() float64 {
	var sumHistory float64
	for _, loss := range *h {
		sumHistory -= loss
	}
	return sumHistory
}

type StepFunc func(History) float64

func Two(h History) float64 {
	return h[len(h)-1] * 2
}

func Fib(h History) float64 {
	var secondPrevWage float64
	if len(h) >= 2 {
		secondPrevWage = h[len(h)-2]
	}
	return h[len(h)-1] + secondPrevWage
}
