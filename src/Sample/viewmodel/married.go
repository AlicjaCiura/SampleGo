package viewmodel

type Married struct {
	Title   string
	Active  string
	Dataset []ResultVm
	Details []DetailsVm
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

type DetailsVm struct {
	ID   string
	Name string
	Link string
}

func NewMarried(dataset []ResultVm) Married {
	result := Married{
		Active:  "malzenstwa",
		Title:   "Malzenstwa - Statystyki",
		Dataset: dataset,
		Details: nil,
	}
	return result
}

func NewMarried2(datasets []DetailsVm) Married {
	result := Married{
		Active:  "malzenstwa",
		Title:   "Malzenstwa - Statystyki",
		Dataset: nil,
		Details: datasets,
	}
	return result
}
