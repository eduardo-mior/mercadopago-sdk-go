package mercadopago

import (
	"time"

	"github.com/eduardo-mior/mercadopago-sdk-go/internal/request"
)

// PaymentResponse é a struct que é usada para receber os dados do request de novo pagamento do MercadoPago.
type PaymentResponse struct {
	CollectorID   int      `json:"collector_id"`   // ID do collector (fornecedor) ??
	OperationType string   `json:"operation_type"` // Tipo da operação (regular_payment, money_transfer)
	Items         []Item   `json:"items"`          // Itens vendidos
	Payer         Payer    `json:"payer"`          // Informações do pagador da cobrança
	BackUrls      BackUrls `json:"back_urls"`      // URLS de redirecionamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn         string         `json:"auto_return"`
	PaymentMethods     PaymentMethods `json:"payment_methods"`      // Configurações das condições de pagamento do pagamento
	ClientID           string         `json:"client_id"`            // ID do cliente do MercadoPago
	Marketplace        string         `json:"marketplace"`          // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee     float64        `json:"marketplace_fee"`      // Comissão de Mercado cobrada pelo proprietario do aplicativo
	Shipments          Shipments      `json:"shipments"`            // Informações de envio dos itens
	NotificationURL    string         `json:"notification_url"`     // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	ExternalReference  string         `json:"external_reference"`   // Nosso ID de controle interno
	AdditionalInfo     string         `json:"additional_info"`      // Informações adicionais do pagamento
	Expires            bool           `json:"expires"`              // Indica se o pagamento possui possui data de expiração
	DateOfExpiration   *time.Time     `json:"date_of_expiration"`   // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom *time.Time     `json:"expiration_date_from"` // A partir de qual data o pagamento estara ativo
	ExpirationDateTo   *time.Time     `json:"expiration_date_to"`   // Até qual data o pagamento estara ativo
	DateCreated        time.Time      `json:"date_created"`         // Data de criação do pagamento (gerado pelo MercadoPago)
	ID                 string         `json:"id"`                   // ID do pagamento do MercadoPago (gerado pelo MercadoPago)
	InitPoint          string         `json:"init_point"`           // Link de pagamento do pagamento
	SandboxInitPoint   string         `json:"sandbox_init_point"`   // Link de pagamento de staging do pagamento
	SiteID             string         `json:"site_id"`              // ID do site do pagamento
}

// PaymentRequest é a struct que é usada para fazer a request de um novo pagamento para o MercadoPago
type PaymentRequest struct {
	ExternalReference string `json:"external_reference"` // Nosso ID de controle interno
	Items             []Item `json:"items"`              // Itens vendidos
	AdditionalInfo    string `json:"additional_info"`    // Informações adicionais do pagamento
	// Indica se o comprador será redirecionado automaticamente para o back_urls após a compra
	// Use "approved" para redirecionar apenas no caso de sucesso
	// User "all" para todos os casos
	AutoReturn          string               `json:"auto_return"`
	BackUrls            BackUrls             `json:"back_urls"`                      // URLS de redirecionamento
	DateOfExpiration    *time.Time           `json:"date_of_expiration"`             // Data de expiração de meios de pagamento em dinheiro
	ExpirationDateFrom  *time.Time           `json:"expiration_date_from"`           // A partir de qual data o pagamento estara ativo
	ExpirationDateTo    *time.Time           `json:"expiration_date_to"`             // Até qual data o pagamento estara ativo
	Expires             bool                 `json:"expires"`                        // Indica se o pagamento possui possui data de expiração
	DifferentialPricing *DifferentialPricing `json:"differential_pricing,omitempty"` // Configuração do preço diferenciado para este pagamento
	Marketplace         string               `json:"marketplace"`                    // Indica de qual marketplace foi feito pagamento (padrão NENHUM)
	MarketplaceFee      float64              `json:"marketplace_fee"`                // Comissão de Mercado cobrada pelo proprietario do aplicativo
	NotificationURL     string               `json:"notification_url"`               // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	Payer               Payer                `json:"payer"`                          // Informações do pagador da cobrança
	PaymentMethods      PaymentMethods       `json:"payment_methods"`                // Configurações das condições de pagamento do pagamento
	StatementDescriptor string               `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	Shipments           Shipments            `json:"shipments"`                      // Informações de envio dos itens
	Tracks              []Track              `json:"tracks,omitempty"`               // Lista trackeamentos que serão executados durante a interação do fluxo de pagamento
}

// BackUrls é a struct que contém as URLs que são utilizadas para redicionar o usuário após a pagamentor ser realizado ou após acontecer algum erro
type BackUrls struct {
	Success string `json:"success"` // URL de redirecionamento para pagamentos aprovados (exemplo pix)
	Pending string `json:"pending"` // URL de redirecionamento para pagamentos que estão com o pagamento pendente (exemplo boleto)
	Failure string `json:"failure"` // URL de redirecionamento para pagamentos que falharão (exemplo cartão de crédito)
}

// DifferentialPricing é a struct que contém o identificador único da configuração de preço diferenciado que sera aplicado ao pagamento
type DifferentialPricing struct {
	ID string `json:"id"` // Identificador único da configuração de preço diferenciado
}

// Payer é a struct que é usada para identificar para quem é o pagamento.
type Payer struct {
	Phone          PayerPhone          `json:"phone"`          // Número de telefone do pagador
	Identification PayerIdentification `json:"identification"` // Documento de identificação do pagador
	Address        PayerAddress        `json:"address"`        // Endereço do pagador
	Email          string              `json:"email"`          // Email do pagador
	Name           string              `json:"name"`           // Nome do pagador (não colocar caracteres especiais)
	Surname        string              `json:"surname"`        // Sobrenome do pagador (não colocar caracteres especiais)
}

// PayerPhone é a struct que contém as informações do telefone do pagador
type PayerPhone struct {
	AreaCode string `json:"area_code"` // Código de área do telefone
	Number   string `json:"number"`    // Número do telefone
}

// PayerIdentification é a struct que contém as informações de identificação do pagador
type PayerIdentification struct {
	Type   string `json:"type"`   // Tipo do documento de identificação (use CPF ou CNPJ)
	Number string `json:"number"` // Número do documentação de identificação
}

// PayerAddress é a struct que contém as informações do endereço do pagador
type PayerAddress struct {
	ZipCode      string  `json:"zip_codigo"`   // CEP do endereço do pagador
	StreetName   string  `json:"street_name"`  // Nome da rua do endereço do pagador
	StreetNumber *string `json:"stree_number"` // Número do endereço do pagador
}

// PaymentMethods é a struct que contém as informações das configurações do pagamento
type PaymentMethods struct {
	ExcludedPaymentMethods []PaymentMethodID `json:"excluded_payment_methods"`  // Lista dos meios de pagamentos bloqueados para esse pagamento
	ExcludedPaymentTypes   []PaymentMethodID `json:"excluded_payment_types"`    // Lista dos tipos de pagamento bloqueados para esse pagamento
	DefaultPaymentMethodID *string           `json:"default_payment_method_id"` // Método de pagamento padrão (preferido, preferencial) para esse pagamento
	Installments           *int              `json:"installments"`              // Número máximo de parcelas
	DefaultInstallments    *int              `json:"default_installments"`      // Número máximo de parcelas padrão
}

// PaymentMethodID é a struct que contém o ID do método de pagamento
type PaymentMethodID struct {
	ID string `json:"id"` // Identificador do método de pagamento do MercadoPago
}

// Item é a struct que contém as informações dos itens que serão cobrados no pagamento
type Item struct {
	ID          string  `json:"id"`                    // Identificador interno nosso de controle
	Title       string  `json:"title"`                 // Titulo do item que é exibido na hora do pagamento
	Description string  `json:"description"`           // Descrição do item (não é exibido na hora do pagamento)
	PictureURL  string  `json:"picture_url,omitempty"` // Imagem do item (não é exibido na hora do pagamento)
	CategoryID  string  `json:"category_id"`           // Identificador da categoria interno nosso de controle
	Quantity    float64 `json:"quantity"`              // Quantidade do item vendido (obrigatório)
	CurrencyID  string  `json:"currency_id"`           // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	UnitPrice   float64 `json:"unit_price"`            // Preço unitário do item vendido (obrigatório)
}

// Shipments é a struct que contém as informações de entrega/envio dos itens do pagamento
type Shipments struct {
	Mode                  string   `json:"mode"`                    // Modo de envio (custom, me2, not_specified)
	LocalPickup           bool     `json:"local_pickup"`            // Preferência de remoção dos pacotes naagência (somente me2)
	Dimensions            string   `json:"dimensions"`              // Tamanho do pacote em cm x cm x cm (somente me2)
	DefaultShippingMethod int      `json:"default_shipping_method"` // Método de envio padrão no _checkout_ (somente me2)
	FreeMethos            []int    `json:"free_methods"`            // IDs dos métodos de envio com frete grátis (somente me2)
	Cost                  float64  `json:"cost"`                    // Custo do frete (somente custom)
	FreeShipping          bool     `json:"free_shipping"`           // Preferência por frete grátis (somente custom)
	ReceiverAddress       *Address `json:"receiver_address"`        // Endereço de envio
}

// Address é a struct que contém as informações do endereço de envio dos itens de um pagamento
type Address struct {
	ZipCode      string `json:"zip_codigo"`   // CEP do endereço de envio
	StreetName   string `json:"street_name"`  // Nome da rua do endereço de envio
	CityName     string `json:"city_name"`    // Nome da cidade do endereço de envio
	StateName    string `json:"state_name"`   // Nome do estado do endereço do envio
	StreetNumber *int   `json:"stree_number"` // Número do endereço de envio
	Floor        string `json:"floor"`        // Número do andar
	Apartment    string `json:"apartment"`    // Número do apartamento
}

// Track é a struct que contém as informações dos trackeamentos que serão executados durante o fluxo de pagamento
type Track struct {
	Type   string `json:"type"`  // Tipo do trackamento (google_ad ou facebook_ad)
	Values string `json:"value"` // Valores de configuração de acordo com o tipo do track (para mais informações consulte a documentação oficial)
}

// PaymentSearchParams é um map[string]interface{} contém os filtros usados no método de Search de pagamentos
type PaymentSearchParams request.QueryParams

// PaymentSearchResponse é a struct que contém todas as informações que são retornadas pela API no método de Search de pagamentos
type PaymentSearchResponse struct {
	NextOffset int                            `json:"next_offset"` // Número de ínicio da próxima busca
	Total      int                            `json:"total"`       // Total de itens encontrados na busca
	Elements   []PaymentSearchElementResponse `json:"elements"`    // Pagamentos retornados da busca
}

// PaymentSearchElementResponse é a struct que contém toda as informações do pagamentos que são retornados no método de Search de pagamentos
type PaymentSearchElementResponse struct {
	ShippingMode       string     `json:"shipping_mode"`
	ID                 string     `json:"id"`
	CollectorID        int        `json:"collector_id"`
	CorporationID      string     `json:"corporation_id"`
	ExternalReference  string     `json:"external_reference"`
	PayerID            *string    `json:"payer_id"`
	PayerEmail         string     `json:"payer_email"`
	ProcessingModes    []string   `json:"processing_modes"`
	ProductID          string     `json:"product_id"`
	DateCreated        time.Time  `json:"date_created"`
	ExpirationDateFrom *time.Time `json:"expiration_date_from"`
	ExpirationDateTo   *time.Time `json:"expiration_date_to"`
	Marketplace        string     `json:"marketplace"`
	ClientID           string     `json:"client_id"`
	SiteID             string     `json:"site_id"`
	Expires            bool       `json:"expires"`
	Items              []string   `json:"items"`
	OperationType      string     `json:"operation_type"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// IdentificationType é a struct que contém as informações de um tipo de documento de identificação
type IdentificationType struct {
	ID        string `json:"id"`         // ID único do tipo de documento de identificação
	Name      string `json:"name"`       // Nome do tipo de documento de identificação
	Type      string `json:"type"`       // Tipo do dado do documento de identificação
	MinLength int    `json:"min_length"` // Tamanho mínimo do documento de identificação
	MaxLength int    `json:"max_length"` // Tamanho máximo do documento de identificação
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// PaymentMethod é a struct que contém as informações de um método de pagamento
type PaymentMethod struct {
	ID                    string                               `json:"id"`                     // ID único do método de pagamento
	Name                  string                               `json:"name"`                   // Nome do método de pagamento
	PaymentTypeID         string                               `json:"payment_type_id"`        // Tipo do meio de pagamento (ticket, atm, credit_card, debit_card, prepaid_card)
	Status                string                               `json:"status"`                 // Status do meio de pagamento (active, deactive, temporally_deactive)
	SecureThumbnail       string                               `json:"secure_thumbnail"`       // Logo do método de pagamento que deve ser mostrada em sites seguros
	Thumbnail             string                               `json:"thumbnail"`              // Logo do método de pagamento
	DeferredCapture       string                               `json:"deferred_capture"`       // Indica se a captura pode ser lenta ou não
	Settings              []PaymentMethodSettings              `json:"settings"`               // Configurações do método de pagamento
	AdditionalInfoNeeded  []string                             `json:"additional_info_needed"` // Lista de informações que devem ser fornecidas pelo pagador
	MinAllowedAmount      float64                              `json:"min_allowed_amount"`     // Mínimo valor que pode ser processado com este meio de pagamento
	MaxAllowedAmount      int                                  `json:"max_allowed_amount"`     // Máxilo valor que pode ser processado com este meio de pagamento
	AccreditationTime     int                                  `json:"accreditation_time"`     // Tempo de processamento do pagamento
	FinancialInstitutions []PaymentMethodFinancialInstitutions `json:"financial_institutions"` // Instituições financeiras de processamento do meio de pagamento
	ProcessingModes       []string                             `json:"processing_modes"`       // Modos de processamento
}

// PaymentMethodSettings é a struct que contém as informações das configurações de um método de pagamento
type PaymentMethodSettings struct {
	Bin          PaymentMethodSettingsBin          `json:"bin"`
	CardNumber   PaymentMethodSettingsCardNumber   `json:"card_number"`   // Informações de configuração do cartão de crédito
	SecurityCode PaymentMethodSettingsSecurityCode `json:"security_code"` // Informações de configuração do código de segurança
}

type PaymentMethodSettingsBin struct {
	Pattern             string `json:"pattern"`              // Expressão regular, representando bines aceitos
	ExclusionPattern    string `json:"exclusion_pattern"`    // Expressão regular, representando bines excluídos
	InstallmentsPattern string `json:"installments_pattern"` // Expressão regular, representando bines aceitos para pagar com mais de uma cota
}

// PaymentMethodSettingsCardNumber é a struct que contém as informações das configurações do cartão de crédito
type PaymentMethodSettingsCardNumber struct {
	Length     int    `json:"length"`     // Comprimento do núemro do cartão de crédito
	Validation string `json:"validation"` // Validação do número do cartão de crédito (standart, none)
}

// PaymentMethodSettingsSecurityCode é a struct que contém as informações do código de seguraça
type PaymentMethodSettingsSecurityCode struct {
	Mode         string `json:"mode"`          // Indica se o código de segurança é obrigatório ou opcional (mandatory, optional)
	Length       int    `json:"length"`        // Comprimento do código de segurança
	CardLocation string `json:"card_location"` // Localização do código de segurança do cartão de crédito (back, front)
}

// PaymentMethodFinancialInstitutions é a struct que contém as informações das instituições financeiras que processamento o método de pagamento
type PaymentMethodFinancialInstitutions struct {
	ID          string `json:"id"`          // ID único da instituição financeira (exemplo atm)
	Description string `json:"description"` // Nome descritivo da instituição financeira
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// WebhookResponse é a struct que é usada para receber os dados do request que o MercadoPago faz para o nosso webhook.
type WebhookResponse struct {
	ID            int              `json:"id"`
	LiveMode      bool             `json:"live_mode"`
	Type          string           `json:"type"`
	DateCreated   string           `json:"date_created"`
	ApplicationID int              `json:"application_id"`
	UserID        int              `json:"user_id"`
	Version       int              `json:"version"`
	APIVersion    string           `json:"api_version"`
	Action        string           `json:"action"`
	Data          WebhookPaymentID `json:"data"`
}

// WebhookPaymentID é a struct que contém o ID do pagamento que foi feito pelo pagador final, esse ID deve ser usado para consultar as informações do pagamento.
type WebhookPaymentID struct {
	ID string `json:"id"`
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse é a struct que é usada para receber os retornos de erro do MercadoPAgo.
type ErrorResponse struct {
	Error   string `json:"error"`   // Slug do erro que retornou
	Message string `json:"message"` // Mensagem de erro relacinada ao campo
	Status  int    `json:"status"`  // Mensagem de erro relacinada ao campo
}
