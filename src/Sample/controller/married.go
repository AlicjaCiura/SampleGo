package controller

import (
	"SampleGo/src/Sample/model"
	"SampleGo/src/Sample/viewmodel"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/withmandala/go-log"
)

type married struct {
	marriedTemplate  *template.Template
	overviewTemplate *template.Template
}

func (m married) registerRoutes() {
	http.HandleFunc("/married", m.handleMarried)
	http.HandleFunc("/overview", m.handleMarried2)
}

func (m married) handleMarried(w http.ResponseWriter, r *http.Request) {
	datasets := test()
	update, _ := model.GetAllRegions()
	diff := difference(update, datasets.Results)
	if len(diff) != 0 {
		model.SaveDb(diff)
	}
	vm := viewmodel.NewMarried(prepare(datasets.Results))
	m.marriedTemplate.Execute(w, vm)
}

func (m married) handleMarried2(w http.ResponseWriter, r *http.Request) {
	dataList := test2().Results
	vm := viewmodel.NewMarried2(prepareOveriview(dataList))
	m.overviewTemplate.Execute(w, vm)
}

func prepare(data []model.MyResult) []viewmodel.ResultVm {
	r := make([]viewmodel.ResultVm, len(data))
	for i := 0; i < len(data); i++ {
		vm := dataToVm(data[i])
		r[i] = vm
	}
	return r
}

func dataToVm(d model.MyResult) viewmodel.ResultVm {
	return viewmodel.ResultVm{
		Name:   d.Name,
		Size:   len(d.Values),
		Values: prepare2(d.Values),
	}
}

func prepare2(data []model.Value) []viewmodel.ValueVm {
	r := make([]viewmodel.ValueVm, len(data))
	for i := 0; i < len(data); i++ {
		vm := dataToVm2(data[i])
		r[i] = vm
	}
	return r
}

func dataToVm2(d model.Value) viewmodel.ValueVm {
	return viewmodel.ValueVm{
		Year: d.Year,
		Val:  d.Val,
	}
}

func test() model.Data {
	resp2, err2 := http.Get("https://bdl.stat.gov.pl/api/v1/data/by-variable/450543?aggregate-id=1&format=json")
	body2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		panic(err2.Error())
	}
	var data2 model.Data
	json.Unmarshal(body2, &data2)
	return data2
}

func test2() model.Overview {
	resp2, err2 := http.Get("https://bdl.stat.gov.pl/api/v1/subjects?parent-id=G535&page=0&page-size=50&lang=pl&format=json")
	body2, err2 := ioutil.ReadAll(resp2.Body)
	if err2 != nil {
		panic(err2.Error())
	}
	var data2 model.Overview
	json.Unmarshal(body2, &data2)
	return data2
}

func prepareOveriview(data []model.Details) []viewmodel.DetailsVm {
	r := make([]viewmodel.DetailsVm, len(data))
	for i := 0; i < len(data); i++ {
		vm := dataToVM2(data[i])
		r[i] = vm
	}
	return r
}

func dataToVM2(d model.Details) viewmodel.DetailsVm {
	return viewmodel.DetailsVm{
		Name: d.Name,
		ID:   d.ID,
		Link: "/" + d.Name,
	}
}

func difference(slice1 []model.MyResult, slice2 []model.MyResult) []model.MyResult {
	logger := log.New(os.Stdout)
	var diff []model.MyResult
	for _, s1 := range slice2 {
		found := false
		for _, s2 := range slice1 {
			if strings.Compare(strings.TrimSpace(s1.Name), strings.TrimSpace(s2.Name)) == 0 {
				found = true
				logger.Infof("	-> Znalezino: %s\n", s1.Name)
				break
			}
		}
		if !found {
			diff = append(diff, s1)
		}
	}
	return diff
}
