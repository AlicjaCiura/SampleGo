package viewmodel

type Married struct {
	Title   string
	Active  string
	Dataset []ResultVm
}

type ResultVm struct {
	Name   string
	Size   int
	Values []ValueVm
}

type ValueVm struct {
	Year string
	Val  float64
}

func NewMarried(dataset []ResultVm) Married {
	result := Married{
		Active:  "malzenstwa",
		Title:   "Malzenstwa - Statystyki",
		Dataset: dataset,
	}
	return result
}
