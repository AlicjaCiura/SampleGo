package model

import (
	"log"
)

type Data struct {
	TotalRecords  int
	VariableID    int
	MeasureUnitID int
	AggregateID   int
	Results       []MyResult
}

type MyResult struct {
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

func SaveDb(list []MyResult) (Result, error) {

	log.Printf("Input %v", len(list))
	sqlStr := "INSERT INTO public.region (name) VALUES "

	for i := 0; i < len(list); i++ {
		log.Printf("Item: %v, %v, %v", i, list[i].Name, list[i].ID)
		sqlStr += " (' "

		sqlStr += list[i].Name
		sqlStr += "'),"

	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	log.Printf("Query: %v", sqlStr)
	res, err := db.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil

}
