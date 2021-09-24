package mercadopago

import (
	"time"

	"github.com/eduardo-mior/mercadopago-sdk-go/internal/request"
)

// PaymentResponse é a struct que é usada para receber os dados do request de novo pagamento do MercadoPago.
type PaymentResponse struct {
	CollectorID   int      `json:"collector_id"`   // Nosso ID do MercadoPago
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
	PictureURL  string  `json:"picture_url,omitempty"` // Imagem do item (imagem que fica nas OG MetaTags do HTML)
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
	ID          int              `json:"id"`           // ID único da resposta do Webhook
	LiveMode    bool             `json:"live_mode"`    // ???
	Type        string           `json:"type"`         // Tipo do evento, exemplo: payment
	DateCreated string           `json:"date_created"` // Data de criação do Webhook
	UserID      string           `json:"user_id"`      // Nosso ID do MercadoPago
	APIVersion  string           `json:"api_version"`  // Versão da API, exemplo: v1
	Action      string           `json:"action"`       // Ação do evento, exemplo: payment.created
	Data        WebhookPaymentID `json:"data"`         // Struct com o ID do pagamento
}

// WebhookPaymentID é a struct que contém o ID do pagamento que foi feito pelo pagador final, esse ID deve ser usado para consultar as informações do pagamento.
type WebhookPaymentID struct {
	ID string `json:"id"` // ID do pagamento que é usado para consultar o status
}

// PaymentConsultResponse é a struct que contém as informações do retorno da Consulta de um pagamento.
type PaymentConsultResponse struct {
	AdditionalInfo            PaymentConsultAdditionalInfo `json:"additional_info"`                // Informações adicionais do pagamento
	AuthorizationCode         *string                      `json:"authorization_code"`             // Código de autorização do pagamento
	BinaryMode                bool                         `json:"binary_mode"`                    // Indica se é o modo binaria de pagamento ou não
	Captured                  bool                         `json:"captured"`                       // Indica se o pagamento foi capturado ou não ???
	Card                      PaymentConsultCard           `json:"card"`                           // Informações do cartão de crédito do pagamento
	CollectorID               int                          `json:"collector_id"`                   // Nosso ID do MercadoPago
	CurrencyID                string                       `json:"currency_id"`                    // Identificador universal da moeda que será usada no pagamento no formato ISO-4217
	DateApproved              *time.Time                   `json:"date_approved"`                  // Data da aprovação do pagamento
	DateCreated               time.Time                    `json:"date_created"`                   // Data da criação do pagamento
	DateLastUpdated           *time.Time                   `json:"date_last_updated"`              // Data da ultima atualização do pagamento
	DateOfExpiration          *time.Time                   `json:"date_of_expiration"`             // Data de expiração de meios de pagamento em dinheiro
	Description               string                       `json:"description"`                    // Descrição do pagamento
	DifferentialPricingId     string                       `json:"differential_pricing_id"`        // Identificador único da configuração de preço diferenciado
	ExternalReference         string                       `json:"external_reference"`             // Nosso ID de controle interno
	FeeDetails                []FeeDetails                 `json:"fee_details"`                    // Informações sobre as taxas que foram aplicadas sobre o pagamento
	ID                        int                          `json:"id"`                             // Identificador único do pagamento
	Installments              int                          `json:"installments"`                   // Número máximo de parcelas
	LiveMode                  bool                         `json:"live_mode"`                      // ???
	MoneyReleaseDate          *time.Time                   `json:"money_release_date"`             // Data da liberação do dinheiro na nossa conta do MercadoPago
	MoneyReleaseSchema        *string                      `json:"money_release_schema"`           // Esquema da liberação do dinheiro
	NotificationURL           string                       `json:"notification_url"`               // URL do Webhook que é chamada quando o Status do pagamento é atualizado
	OperationType             string                       `json:"operation_type"`                 // Tipo do pagamento (consulte a documentação para saber oque significa)
	Payer                     Payer                        `json:"payer"`                          // Informações do pagador
	PaymentMethodID           string                       `json:"payment_method_id"`              // ID do método de pagamento (exemplo: master)
	PaymentTypeID             string                       `json:"payment_type_id"`                // Tipo do método de pagamento (exemplo: credit_card)
	ProcessingModes           string                       `json:"processing_modes"`               // Modo de processamento
	StatementDescriptor       string                       `json:"statement_descriptor,omitempty"` // Descrição do pagamento que ira aparecer no extrato do cartão
	TransactionAmount         float64                      `json:"transaction_amount"`             // Valor pago
	TransactionAmountRefunded float64                      `json:"transaction_amount_refunded"`    // Valor do reembolso
	TransactionDetails        TransactionDetails           `json:"transaction_details"`            // Detalhes da transação

	// Código de barras do boleto
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct só é retornada quando o pagamento é feito em boleto ou PEC(lotéricas), porém isso não esta documentado em nenhum lugar na documentação do MercadoPago
	Barcode *Barcode `json:"barcode"`

	// QRCode do Pix e Chave Copia-e-Cola do Pix
	// Nos testes que eu (Eduardo Mior) fiz, essa Struct só é retornada quando o pagamento é feito em pix, porém isso não esta documentado em nenhum lugar na documetação do MercadoPago
	PointOfInteraction *PointOfInteraction `json:"point_of_interaction"`

	// Status do pagamento (segundo documentação oficial do mercado pago)
	// approved - Pagamento aprovado
	// pending - Pagamento pendente
	// authorized - Pagamento autorizado porém ainda não aprovado
	// in_process - O pagamento esta sendo revisado pelo MercadoPago
	// in_mediation - Foi aberta uma disputa no pagamento e ele esta em revisão
	// rejected - Pagamento rejeitado (cartão de crédito)
	// cancelled - O pagamento foi cancelado por uma das partes ou porque ele expirou
	// refunded - O pagamento foi reembolsado para o usuário
	// charged_back - Foi feito um estorno do pagamento no cartão de crédito do usuário
	Status string `json:"status"`

	// Detalhes sobre o status do pagamento. A lista completa pode ser consultada em https://www.mercadopago.com.br/developers/pt/guides/online-payments/checkout-api/handling-responses
	// cc_rejected_card_disabled (cartão de crédito bloqueado)
	// cc_rejected_blacklist (cartão de crédito recusado)
	// cc_rejected_high_risk (cartão de crédito recusado)
	// rejected_high_risk (paypal recusado)
	// pending_review_manual (pagamendo sendo revisado pelo mercadopago)
	// pending_waiting_transfer (aguardando pagamento do pix / aguardando transferência do dinheiro)
	// pending_waiting_payment (aguardando pagamento do boleto ou do PEC(lotérica))
	// accredited (aprovado)
	// https://prnt.sc/1ta9w4b
	StatusDetail string `json:"status_detail"`
}

// PaymentConsultCard é a struct que contém as informações do cartão de crédito que efetuou o pagamento
type PaymentConsultCard struct {
	Cardholder      Cardholder `json:"cardholder"`        // Informações do dono do cartão
	DateCreated     *time.Time `json:"date_created"`      // Data de criação do cartão
	DateLastUpdated *time.Time `json:"date_last_updated"` // Data da ultima atualização do cartão
	ExpirationMonth int        `json:"expiration_month"`  // Mês de expiração do cartão
	ExpirationYear  int        `json:"expiration_year"`   // Ano de expiração do cartão
	FirstSixDigits  string     `json:"first_six_digits"`  // Seis primeiros digitos do cartão
	LastFourDigits  string     `json:"last_four_digits"`  // Ultimos quatro digitos do cartão
}

// Cardholder é a struct que contém as informações do dono do cartão
type Cardholder struct {
	Identification PayerIdentification `json:"identification"` // Informações de identificação do cartão de crédito
	Name           string              `json:"name"`           // Nome do dono do cartão de crédito
}

// PaymentConsultAdditionalInfo é a struct que contém informações adicionais sobre o pagamento
type PaymentConsultAdditionalInfo struct {
	IPAddress string `json:"ip_address"` // IP do usuário que pagou

	// Items     []Item `json:"items"`
	// Por algum motivo eles retoram a Quantity do item e o UnityPrice como String e isso faz com que aconteça erro no Unmarshal do JSON.
	// Por esse motivo foi comentado o campo Items das informações adicionais.
}

// FeeDetails é a struct que contém as informações sobre a taxa que foi cobrada sobre o pagamento
type FeeDetails struct {
	Amount   float64 `json:"amount"`    // Porcentagem da taxa que foi paga
	FeePayer string  `json:"fee_payer"` // Indica que ira pagar a taxa
	Type     string  `json:"type"`      // Indica quem esta cobrando a taxa
}

// TransactionDetails é a struct que contém as informações sobre os detalhes da transação
type TransactionDetails struct {
	ExternalResourceURL      *string `json:"external_resource_url"`       // URL do Boleto ou do PEC caso a forma de pagamento for boleto ou lotéricas
	FinancialInstitution     *string `json:"financial_institution"`       // Instituição financeira responsavel pelo pagamento
	TotalPaidAmount          float64 `json:"total_paid_amount"`           // Valor total pago
	InstallmentAmount        float64 `json:"installment_amount"`          // Valor do pagamento / Valor da parcela
	NetReceivedAmount        float64 `json:"net_received_amount"`         // Valor liquido recebido com o valor descontado das taxas
	OverpaidAmount           float64 `json:"overpaid_amount"`             // Valor pago em excesso ???
	PayableDeferralPeriod    string  `json:"payable_deferral_period"`     // ????
	PaymentMethodReferenceID string  `json:"payment_method_reference_id"` // ID de referencia do método de pagamento
	TransactionID            *string `json:"transaction_id"`              // ID da transação
}

// Barcode é a struct que contém o código de barras do boleto
type Barcode struct {
	Content string `json:"content"` // Código de barras do boleto
}

// PointOfInteraction é a struct que contém as informações dos dados da transação no caso do pagamento ser em Pix (pelos testes que eu fiz essa struct só é preenchida quando é pix)
type PointOfInteraction struct {
	TransactionData TransactionData `json:"transaction_data"` // Informações do QRCode
}

// TransactionData é a struct que contém as informações do Base64 do QRCode e a chave Pix Copia-e-Cola
type TransactionData struct {
	QrCode       string `json:"qr_code"`        // Chave Pix Copia-e-Cola
	QrCodeBase64 string `json:"qr_code_base64"` // Base64 Do QRCode do Pix
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse é a struct que é usada para receber os retornos de erro do MercadoPago.
type ErrorResponse struct {
	Error   string `json:"error"`   // Slug do erro que retornou
	Message string `json:"message"` // Mensagem de erro relacinada ao campo
	Status  int    `json:"status"`  // Mensagem de erro relacinada ao campo
}
