package rajaongkir

import (
	"errors"
	"github.com/tss182/api"
)

type Rajaongkir struct {
	Token           string
	Province        string
	City            string
	District        string
	Origin          string
	Destination     string
	Weight          int
	Courier         string
	Length          int
	Width           int
	Height          int
	Diameter        int
	DestinationType string
}

const (
	TypeKec  = "subdistrict"
	TypeKota = "city"
)

type APIResultAddress struct {
	Rajaongkir struct {
		//Query struct {
		//	City string `json:"city"`
		//} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results []struct {
			SubdistrictID   string `json:"subdistrict_id"`
			ProvinceID      string `json:"province_id"`
			Province        string `json:"province"`
			CityID          string `json:"city_id"`
			City            string `json:"city_name"`
			City2           string `json:"city"`
			Type            string `json:"type"`
			SubdistrictName string `json:"subdistrict_name"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func (dt *Rajaongkir) GetProvince() (APIResultAddress, error) {
	apiInit := api.Init{}
	apiInit.Header = map[string]string{"key": dt.Token}
	apiInit.Url = "https://pro.rajaongkir.com/api/province"
	apiInit.Data = map[string]interface{}{}
	if dt.Province != "" {
		apiInit.Data = map[string]interface{}{"id": dt.Province}
	}
	err := apiInit.Do("GET")
	if err != nil {
		return APIResultAddress{}, err
	}
	var r APIResultAddress
	err = apiInit.Get(&r)
	if err != nil {
		return APIResultAddress{}, err
	}

	if r.Rajaongkir.Status.Code != 200 {
		return APIResultAddress{}, errors.New(r.Rajaongkir.Status.Description)
	}
	return r, nil
}

func (dt *Rajaongkir) GetCity() (APIResultAddress, error) {
	apiInit := api.Init{}
	apiInit.Header = map[string]string{"key": dt.Token}
	apiInit.Url = "https://pro.rajaongkir.com/api/city"
	apiInit.Data = map[string]interface{}{}
	if dt.Province != "" {
		apiInit.Data["province"] = dt.Province
	}
	if dt.City != "" {
		apiInit.Data["id"] = dt.City
	}
	err := apiInit.Do("GET")
	if err != nil {
		return APIResultAddress{}, err
	}
	var r APIResultAddress
	err = apiInit.Get(&r)
	if err != nil {
		return APIResultAddress{}, err
	}
	//fmt.Println(apiInit.GetRaw())

	if r.Rajaongkir.Status.Code != 200 {
		return APIResultAddress{}, errors.New(r.Rajaongkir.Status.Description)
	}
	return r, nil
}

func (dt *Rajaongkir) GetSUbDistrict() (APIResultAddress, error) {
	apiInit := api.Init{}
	apiInit.Header = map[string]string{"key": dt.Token}
	apiInit.Url = "https://pro.rajaongkir.com/api/subdistrict"
	apiInit.Data = map[string]interface{}{}
	if dt.District != "" {
		apiInit.Data["id"] = dt.District
	}
	if dt.City != "" {
		apiInit.Data["city"] = dt.City
	}
	err := apiInit.Do("GET")
	if err != nil {
		return APIResultAddress{}, err
	}
	var r APIResultAddress
	err = apiInit.Get(&r)
	if err != nil {
		return APIResultAddress{}, err
	}
	if r.Rajaongkir.Status.Code != 200 {
		return APIResultAddress{}, errors.New(r.Rajaongkir.Status.Description)
	}
	for i, v := range r.Rajaongkir.Results {
		r.Rajaongkir.Results[i].City = v.City2
	}
	return r, nil
}

func (dt *Rajaongkir) GetCost() (APIResultAddress, error) {
	if dt.Origin == "" || dt.Destination == "" || dt.Weight == 0 || dt.Courier != "" || dt.DestinationType != "" {
		return APIResultAddress{}, errors.New("origin, destination, weight and courier is required")
	}
	apiInit := api.Init{}
	apiInit.Header = map[string]string{"key": dt.Token}
	apiInit.Url = "https://pro.rajaongkir.com/api/cost"
	apiInit.Data = map[string]interface{}{
		"origin":          dt.Origin,
		"destionation":    dt.Destination,
		"weight":          dt.Weight,
		"courier":         dt.Courier,
		"destinationType": dt.DestinationType,
	}

	err := apiInit.Do("POST")

	if err != nil {
		return APIResultAddress{}, err
	}
	var r APIResultAddress
	err = apiInit.Get(&r)
	if err != nil {
		return APIResultAddress{}, err
	}
	if r.Rajaongkir.Status.Code != 200 {
		return APIResultAddress{}, errors.New(r.Rajaongkir.Status.Description)
	}
	return r, nil
}
