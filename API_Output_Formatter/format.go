package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-orders-headers-creates-subfunc-rmq-kube/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*Header, error) {
	calculateOrderID := psdc.CalculateOrderID
	headerBPCustomerSupplier := psdc.HeaderBPCustomerSupplier

	header := Header{}
	inputHeader := sdc.Orders
	inputData, err := json.Marshal(inputHeader)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inputData, &header)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(headerBPCustomerSupplier)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &header)
	if err != nil {
		return nil, err
	}

	header.OrderID = calculateOrderID.OrderIDLatestNumber

	return &header, nil
}

func ConvertToHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartner, error) {
	var headerPartners []HeaderPartner
	calculateOrderID := psdc.CalculateOrderID
	headerPartnerFunction := psdc.HeaderPartnerFunction
	headerPartnerBPGeneral := psdc.HeaderPartnerBPGeneral
	headerPartnerFunctionMap := make(map[int]api_processing_data_formatter.HeaderPartnerFunction, len(*headerPartnerFunction))
	headerPartnerBPGeneralMap := make(map[int]api_processing_data_formatter.HeaderPartnerBPGeneral, len(*headerPartnerBPGeneral))

	for _, v := range *headerPartnerFunction {
		headerPartnerFunctionMap[*v.BusinessPartner] = v
	}

	for _, v := range *headerPartnerBPGeneral {
		headerPartnerBPGeneralMap[*v.BusinessPartner] = v
	}

	for key := range headerPartnerFunctionMap {
		headerPartners = append(headerPartners, HeaderPartner{
			OrderID:                 calculateOrderID.OrderIDLatestNumber,
			PartnerFunction:         headerPartnerFunctionMap[key].PartnerFunction,
			BusinessPartner:         headerPartnerBPGeneralMap[key].BusinessPartner,
			BusinessPartnerFullName: headerPartnerBPGeneralMap[key].BusinessPartnerFullName,
			BusinessPartnerName:     headerPartnerBPGeneralMap[key].BusinessPartnerName,
			Country:                 headerPartnerBPGeneralMap[key].Country,
			Language:                headerPartnerBPGeneralMap[key].Language,
			Currency:                headerPartnerBPGeneralMap[key].Currency,
			AddressID:               headerPartnerBPGeneralMap[key].AddressID,
		})
	}

	return &headerPartners, nil
}

func ConvertToHeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlants []HeaderPartnerPlant
	calculateOrderID := psdc.CalculateOrderID
	headerPartnerPlant := psdc.HeaderPartnerPlant

	for i := range *headerPartnerPlant {
		headerPartnerPlants = append(headerPartnerPlants, HeaderPartnerPlant{
			OrderID:         calculateOrderID.OrderIDLatestNumber,
			PartnerFunction: (*headerPartnerPlant)[i].PartnerFunction,
			BusinessPartner: (*headerPartnerPlant)[i].BusinessPartner,
			Plant:           (*headerPartnerPlant)[i].Plant,
		})
	}

	return &headerPartnerPlants, nil
}
