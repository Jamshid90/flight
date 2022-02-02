package sort

import (
	"github.com/Jamshid90/flight/internal/entity"
)

func SortByNumber(list []*entity.Flight) []*entity.Flight {
	result := make(chan []*entity.Flight)
	defer close(result)
	go MergeSortByNumber(list, result)
	return <-result
}

func MergeSortByNumber(data []*entity.Flight, r chan []*entity.Flight) {
	if len(data) <= 1 {
		r <- data
		return
	}

	middle := len(data) / 2
	leftChan := make(chan []*entity.Flight)
	rightChan := make(chan []*entity.Flight)

	defer close(leftChan)
	defer close(rightChan)

	go MergeSortByNumber(data[:middle], leftChan)
	go MergeSortByNumber(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	r <- Merge(ldata, rdata)
}

func Merge(ldata []*entity.Flight, rdata []*entity.Flight) (result []*entity.Flight) {
	result = make([]*entity.Flight, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx].Number < rdata[ridx].Number:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}