package model

import (
	"os"

	"github.com/withmandala/go-log"
)

//Data is struct for data from stats.gov.pl
type Data struct {
	TotalRecords  int
	VariableID    int
	MeasureUnitID int
	AggregateID   int
	Results       []MyResult
}

//MyResult is struct for result by region, Name is name of region, Values is list of values
type MyResult struct {
	ID     string
	Name   string
	Values []Value
}

func (t MyResult) String() string {
	return t.Name
}

//Value is struct for values by region, Year is year for values, Val ist values
type Value struct {
	Year   string
	Val    float64
	AttrID int
}

//Overview is struct for all informtion from stats.gov.pl
type Overview struct {
	Results []Details
}

//Details is struct
type Details struct {
	ID   string
	Name string
}

//SaveDb function to save regions to db
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

//GetAllRegions function to get all regions from db
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
