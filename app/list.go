package app

import (
	"strconv"
	"math"
)

/*
 * Structure
 */
 
type List struct {
	Series []int
}

/*
 * Constructor
 */

func NewList(series []string) *List {
	newList := make([]int, 0)

	for _, element := range series {
		if element != "" {
			number, err := strconv.Atoi(element)
			handleError(err)

			newList = append(newList, number)
		}
	}

	return &List{
		Series: newList,
	}
}

/*
 * Methods
 */

// Return numbers of series
func (list *List) Length() int {
	return len(list.Series)
}

// Return the average of all results
func (list *List) StatAll() float64 {
	stat := 0

	for _, element := range list.Series {
			stat += element
	}

	return math.Round(float64(stat) / float64(list.Length()))
}

// Return the average of the 5 last results
func (list *List) StatLast() float64 {
	stat := 0
	series := list.Series[list.Length() - 5:] // Select the 5 last series

	for _, element := range series {
		stat += element
	}

	return math.Round(float64(stat) / 5.0)
}

// Add a new serie to the list
func (list *List) AddSerie(serie int) []int {
	list.Series = append(list.Series, serie)
	return list.Series
}

/*
 * Utils
 */

func handleError(e error) {
    if e != nil {
        panic(e)
    }
}
