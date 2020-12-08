package model

import (
	"os"

	"github.com/withmandala/go-log"
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

func (t MyResult) String() string {
	return t.Name
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

	logger := log.New(os.Stdout).WithColor()
	logger.Info("Input %v", len(list))
	sqlStr := "INSERT INTO public.region (name) VALUES "

	for i := 0; i < len(list); i++ {
		logger.Infof("Item: %v, %v, %v", i, list[i].Name, list[i].ID)
		sqlStr += " ('"
		sqlStr += list[i].Name
		sqlStr += "'),"
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	//prepare the statement
	logger.Infof("Query: %v", sqlStr)
	res, err := db.Exec(sqlStr)
	if err != nil {
		logger.Errorf("&v", err)
		return nil, err
	}
	return res, nil
}

func GetAllRegions() ([]MyResult, error) {

	log := log.New(os.Stdout)
	sqlStr := "SELECT * FROM public.region "
	//prepare the statement
	log.Infof("Query: %v", sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var res []MyResult
	for rows.Next() {
		result := MyResult{}
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, result)
	}
	return res, nil

}
