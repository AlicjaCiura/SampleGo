package viewmodel

type Married struct {
	Title   string
	Active  string
	Dataset []ResultVm
}

/*
type Data struct {
	TotalRecords  int
	VariableID    int
	MeasureUnitID int
	AggregateID   int
	LastUpdate    time.Time
	Results       []Result
}

type Result struct {
	ID     string
	Name   string
	Values []Value
}

type Value struct {
	Year   string
	Val    float64
	AttrID int
}*/

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
