package go_jsd

import "math"

// WFD holds the word file distribution for two file A and B
type WFD struct {
	A map[string]float64
	B map[string]float64
}

func CalculateJSD(wfd WFD) float64 {
	// calculate the avg of the distributions first
	initialCap := len(wfd.A)
	if len(wfd.B) > initialCap {
		initialCap = len(wfd.B)
	}
	wfdAvg := make(map[string]float64, initialCap)
	for word, wd := range wfd.A {
		wfdAvg[word] = (wd + wfd.B[word]) / 2.0
	}
	for word, wd := range wfd.B {
		if _, ok := wfd.A[word]; ok {
			continue
		}
		wfdAvg[word] = (wd + wfd.A[word]) / 2.0
	}

	// calculate the KLD for each distribution
	kldA := 0.0
	for word := range wfd.A {
		kldA += wfd.A[word] * math.Log2(wfd.A[word]/wfdAvg[word])
	}

	kldB := 0.0
	for word := range wfd.B {
		kldB += wfd.B[word] * math.Log2(wfd.B[word]/wfdAvg[word])
	}

	// return the distance
	return math.Sqrt(0.5*kldA + 0.5*kldB)
}
