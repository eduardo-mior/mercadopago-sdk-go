package mercadopago

import (
	"encoding/json"
	"os"

	"github.com/eduardo-mior/mercadopago-sdk-go/internal/request"
)

const BASEURL = "https://api.mercadopago.com"

// CreatePayment é o método responsável por criar um pagamento no MercadoPago.
func CreatePayment(paymentRequest PaymentRequest, mercadoPagoAccessToken ...string) (*PaymentResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    paymentRequest,
		Headers: map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:     BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

// UpdatePayment é o método responsável por atualizar as informações de um pagamento no MercadoPago.
func UpdatePayment(paymentID string, paymentRequest PaymentRequest, mercadoPagoAccessToken ...string) (*PaymentResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:     "PUT",
		PathParams: request.PathParams{paymentID},
		Body:       paymentRequest,
		Headers:    map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:        BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

// GetPayment é o método responsável buscar todas as informações de um pagamento no MercadoPago.
func GetPayment(paymentID string, mercadoPagoAccessToken ...string) (*PaymentResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:     "GET",
		PathParams: request.PathParams{paymentID},
		Headers:    map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:        BASEURL + "/checkout/preferences",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse PaymentResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

// SearchPayments é o método responsável buscar todas as informações de um pagamento no MercadoPago.
// Como não existe nenhuma documentação completa sobre como esse EndPoint funciona então ele recebe por parametro qualquer filtro.
// Segundo oque consta nos SDKs oficiais e alguns não oficiais do MercadoPago, esse EndPoint é baseado em "Criteria Filters", ou seja,
// você pode filtrar por qualquer campo do pagamento usando qualquer operador, exemplo external_reference=525.
func SearchPayments(searchParams PaymentSearchParams, mercadoPagoAccessToken ...string) (*PaymentSearchResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:      "GET",
		QueryParams: request.QueryParams(searchParams),
		Headers:     map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:         BASEURL + "/checkout/preferences/search",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentSearchResponse PaymentSearchResponse
	err = json.Unmarshal(response.RawBody, &paymentSearchResponse)
	return &paymentSearchResponse, nil, err
}

// ConsultPayment é o método responsável consultar as informações atualizadas de um pagamento no MercadoPago, incluindo Status.
func ConsultPayment(paymentID string, mercadoPagoAccessToken ...string) (*PaymentConsultResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:     "GET",
		PathParams: request.PathParams{paymentID},
		Headers:    map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:        BASEURL + "/v1/payments",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentResponse PaymentConsultResponse
	err = json.Unmarshal(response.RawBody, &paymentResponse)
	return &paymentResponse, nil, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// GetIdentificationTypes é o método responsável retornar todos o tipos de documento de identificação do MercadoPago.
func GetIdentificationTypes(mercadoPagoAccessToken ...string) ([]IdentificationType, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:     BASEURL + "/v1/identification_types",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var identificationTypes []IdentificationType
	err = json.Unmarshal(response.RawBody, &identificationTypes)
	return identificationTypes, nil, err
}

// GetPaymentMethods é o método responsável retornar todos o tipos de documento de identificação do MercadoPago.
func GetPaymentMethods(mercadoPagoAccessToken ...string) ([]PaymentMethod, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"Authorization": "Bearer " + getAccessToken(mercadoPagoAccessToken...)},
		URL:     BASEURL + "/v1/payment_methods",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var paymentMethods []PaymentMethod
	err = json.Unmarshal(response.RawBody, &paymentMethods)
	return paymentMethods, nil, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// parseError é a função que pega os dados do erro do MercadoPago e retorna em formato de Struct.
func parseError(body []byte) (*ErrorResponse, error) {
	var errResponse ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}

// getAccessToken é a função responsável por retornar o AccessToken do mercado pago.
// Caso tenha sido passado um token por parametro pegamos o token passado por parametro, se não pegamos da variavel de ambiente MERCADO_PAGO_ACCESS_TOKEN.
func getAccessToken(mercadoPagoAccessToken ...string) string {
	if len(mercadoPagoAccessToken) > 0 {
		return mercadoPagoAccessToken[0]
	} else {
		return os.Getenv("MERCADO_PAGO_ACCESS_TOKEN")
	}
}
