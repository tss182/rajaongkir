package rajaongkir

import (
	"errors"
	"github.com/tss182/api-go"
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
	Waybill     string
	Length      int
	Width       int
	Height      int
	Diameter    int
	TypeOrigin  string
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
			Code            string `json:"code"`
			Name            string `json:"name"`
			Costs           []struct {
				Service     string `json:"service"`
				Description string `json:"description"`
				Cost        []struct {
					Value int    `json:"value"`
					Etd   string `json:"etd"`
					Note  string `json:"note"`
				} `json:"cost"`
			} `json:"costs"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

type APITracking struct {
	Rajaongkir struct {
		Query struct {
			Waybill string `json:"waybill"`
			Courier string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Result struct {
			Delivered bool `json:"delivered"`
			Summary   struct {
				CourierCode   string `json:"courier_code"`
				CourierName   string `json:"courier_name"`
				WaybillNumber string `json:"waybill_number"`
				ServiceCode   string `json:"service_code"`
				WaybillDate   string `json:"waybill_date"`
				ShipperName   string `json:"shipper_name"`
				ReceiverName  string `json:"receiver_name"`
				Origin        string `json:"origin"`
				Destination   string `json:"destination"`
				Status        string `json:"status"`
			} `json:"summary"`
			Details struct {
				WaybillNumber    string      `json:"waybill_number"`
				WaybillDate      string      `json:"waybill_date"`
				WaybillTime      string      `json:"waybill_time"`
				Weight           string      `json:"weight"`
				Origin           string      `json:"origin"`
				Destination      string      `json:"destination"`
				ShippperName     string      `json:"shippper_name"`
				ShipperAddress1  string      `json:"shipper_address1"`
				ShipperAddress2  interface{} `json:"shipper_address2"`
				ShipperAddress3  interface{} `json:"shipper_address3"`
				ShipperCity      string      `json:"shipper_city"`
				ReceiverName     string      `json:"receiver_name"`
				ReceiverAddress1 string      `json:"receiver_address1"`
				ReceiverAddress2 string      `json:"receiver_address2"`
				ReceiverAddress3 string      `json:"receiver_address3"`
				ReceiverCity     string      `json:"receiver_city"`
			} `json:"details"`
			DeliveryStatus struct {
				Status      string `json:"status"`
				PodReceiver string `json:"pod_receiver"`
				PodDate     string `json:"pod_date"`
				PodTime     string `json:"pod_time"`
			} `json:"delivery_status"`
			Manifest []struct {
				ManifestCode        string `json:"manifest_code"`
				ManifestDescription string `json:"manifest_description"`
				ManifestDate        string `json:"manifest_date"`
				ManifestTime        string `json:"manifest_time"`
				CityName            string `json:"city_name"`
			} `json:"manifest"`
		} `json:"result"`
	} `json:"rajaongkir"`
}

type APIResultAddressSingle struct {
	Rajaongkir struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Results struct {
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
	apiInit := api.Api{}
	apiInit.Method = api.MethodGET
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/province"
	if dt.Province != "" {
		apiInit.Body = map[string]interface{}{"id": dt.Province}
	}
	err := apiInit.Do()
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
	apiInit := api.Api{}
	apiInit.Method = api.MethodGET
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/city"
	data := map[string]interface{}{}
	if dt.Province != "" {
		data["province"] = dt.Province
	}
	if dt.City != "" {
		data["id"] = dt.City
	}
	apiInit.Body = data
	err := apiInit.Do()
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

func (dt *Rajaongkir) GetSubDistrict() (APIResultAddress, error) {
	apiInit := api.Api{}
	apiInit.Method = api.MethodGET
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/subdistrict"
	data := map[string]interface{}{}
	if dt.City != "" {
		data["city"] = dt.City
	}
	apiInit.Body = data
	err := apiInit.Do()
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

func (dt *Rajaongkir) GetSubDistrictSingle() (APIResultAddressSingle, error) {
	apiInit := api.Api{}
	apiInit.Method = api.MethodGET
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/subdistrict"
	data := map[string]interface{}{}
	if dt.District != "" {
		data["id"] = dt.District
	}
	apiInit.Body = data
	err := apiInit.Do()
	if err != nil {
		return APIResultAddressSingle{}, err
	}
	var r APIResultAddressSingle
	err = apiInit.Get(&r)
	if err != nil {
		return APIResultAddressSingle{}, err
	}
	if r.Rajaongkir.Status.Code != 200 {
		return APIResultAddressSingle{}, errors.New(r.Rajaongkir.Status.Description)
	}
	r.Rajaongkir.Results.City = r.Rajaongkir.Results.City2
	return r, nil
}

func (dt *Rajaongkir) GetCost() (APIResultAddress, error) {
	if dt.Origin == "" || dt.Destination == "" || dt.Weight == 0 || dt.Courier == "" || dt.TypeOrigin == "" {
		return APIResultAddress{}, errors.New("origin, destination, weight and courier is required")
	}
	apiInit := api.Api{}
	apiInit.Method = api.MethodPOST
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/cost"
	apiInit.Body = map[string]interface{}{
		"origin":          dt.Origin,
		"destination":     dt.Destination,
		"weight":          dt.Weight,
		"courier":         dt.Courier,
		"originType":      dt.TypeOrigin,
		"destinationType": dt.TypeOrigin,
	}

	err := apiInit.Do()

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

func (dt *Rajaongkir) Tracking() (APITracking, error) {
	if dt.Waybill == "" || dt.Courier == "" {
		return APITracking{}, errors.New("waybill and courier is required")
	}
	apiInit := api.Api{}
	apiInit.Method = api.MethodPOST
	apiInit.ContentType = api.TypeUrlEncode
	apiInit.HeaderAdd("key", dt.Token)
	apiInit.Url = "https://pro.rajaongkir.com/api/waybill"
	apiInit.Body = map[string]interface{}{
		"waybill": dt.Waybill,
		"courier": dt.Courier,
	}

	err := apiInit.Do()

	if err != nil {
		return APITracking{}, err
	}
	var r APITracking
	err = apiInit.Get(&r)
	if err != nil {
		return APITracking{}, err
	}
	if r.Rajaongkir.Status.Code != 200 {
		return APITracking{}, errors.New(r.Rajaongkir.Status.Description)
	}
	return r, nil
}
