package subfunction

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var outHeader *dpfm_api_output_formatter.Header
	var outHeaderPartner *[]dpfm_api_output_formatter.HeaderPartner
	var outHeaderPartnerPlant *[]dpfm_api_output_formatter.HeaderPartnerPlant
	var err error

	outHeader, err = dpfm_api_output_formatter.ConvertToHeader(sdc, psdc)
	if err != nil {
		return nil, err
	}
	outHeaderPartner, err = dpfm_api_output_formatter.ConvertToHeaderPartner(sdc, psdc)
	if err != nil {
		return nil, err
	}
	outHeaderPartnerPlant, err = dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(sdc, psdc)
	if err != nil {
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Header:             *outHeader,
		HeaderPartner:      *outHeaderPartner,
		HeaderPartnerPlant: *outHeaderPartnerPlant,
	}

	return osdc, nil
}
