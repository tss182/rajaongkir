package rajaongkir

import (
	"errors"
	"github.com/tss182/api"
)

type Rajaongkir struct {
	Token       string
	Province    string
	City        string
	District    string
	Origin      string
	Destination string
	Weight      int
	Courier     string
	Length      int
	Width       int
	Height      int
	Diameter    int
}
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
			City            string `json:"city"`
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
	err := apiInit.Do("GET")
	apiInit.Data = map[string]interface{}{}
	if dt.Province != "" {
		apiInit.Data["province"] = dt.Province
	}
	if dt.City != "" {
		apiInit.Data["id"] = dt.City
	}
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
	err := apiInit.Do("GET")
	apiInit.Data = map[string]interface{}{}
	if dt.District != "" {
		apiInit.Data["id"] = dt.District
	}
	if dt.City != "" {
		apiInit.Data["city"] = dt.City
	}
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

func (dt *Rajaongkir) GetCost() (APIResultAddress, error) {
	if dt.Origin == "" || dt.Destination == "" || dt.Weight == 0 || dt.Courier != "" {
		return APIResultAddress{}, errors.New("origin, destination, weight and courier is required")
	}
	apiInit := api.Init{}
	apiInit.Header = map[string]string{"key": dt.Token}
	apiInit.Url = "https://pro.rajaongkir.com/api/cost"
	err := apiInit.Do("GET")
	//if district != "" {
	//	apiInit.Data["id"] = district
	//}
	//if city != "" {
	//	apiInit.Data["city"] = city
	//}
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
