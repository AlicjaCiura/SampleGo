package controller

import (
	"SampleGo/src/Sample/model"
	"SampleGo/src/Sample/viewmodel"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type married struct {
	marriedTemplate *template.Template
}

func (m married) registerRoutes() {
	http.HandleFunc("/married", m.handleMarried)
}

func (m married) handleMarried(w http.ResponseWriter, r *http.Request) {
	datasets := test()
	vm := viewmodel.NewMarried(prepare(datasets.Results))
	m.marriedTemplate.Execute(w, vm)
}

func prepare(data []model.Result) []viewmodel.ResultVm {
	r := make([]viewmodel.ResultVm, len(data))
	for i := 0; i < len(data); i++ {
		vm := dataToVm(data[i])
		r[i] = vm
	}
	return r
}

func dataToVm(d model.Result) viewmodel.ResultVm {
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
