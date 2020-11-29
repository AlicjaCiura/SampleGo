package model

type Data struct {
	TotalRecords  int
	VariableID    int
	MeasureUnitID int
	AggregateID   int
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
}

type Overview struct {
	Results []Details
}

type Details struct {
	ID   string
	Name string
}
