![GoLang MercadoPago](https://i.imgur.com/fUzaPBC.png)
# MercadoPago SDK para Go
SDK (não oficial) para consumir os serviços do MercadoPago em Golang.

## 📲  Instalação
Para baixar o SDK basta utilizar o seguinte comando:
```bash
$ go get -u github.com/eduardo-mior/mercadopago-sdk-go
```

## 🛠 Funcionalidades do SDK
Funcionalidades disponíbilizadas no SDK:
- Criação de um pagamento
- Atualização de um pagamento
- Buscar informações de um pagamento
- Consultar situação de um pagamento
- Listagem/Busca de pagamento
- Listagem de tipos de documento de identificação do MercadoPago
- Listagem de métodos de pagamento e suas confiurações

## 🌟  Começando 
Para começar você deve fazer o `import` do SDK, para isso basta adicionar a seguinte linha no seu código:
```go
import "github.com/eduardo-mior/mercadopago-sdk-go"
```
Feito isso já esta tudo pronto para você começar a utilizaro SDK!

## 💻  Exemplos de uso
Criando um pagamento:
```go
response, mercadopagoErr, err := mercadopago.CreatePayment(mercadopago.PaymentRequest{
    ExternalReference: "seu-id-interno-0001",
    Items: []mercadopago.Item{
        {
            Title:     "Pagamendo mensalidade PagueTry",
            Quantity:  1,
            UnitPrice: 50,
        },
    },
    Payer: mercadopago.Payer{
        Identification: mercadopago.PayerIdentification{
            Type:   "CPF",
            Number: "12345678909",
        },
        Name:    "Eduardo",
        Surname: "Mior",
        Email:   "eduardo-mior@hotmail.com",
    },
    NotificationURL:   "https://localhost/webhook/mercadopago",
})

if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

Atualizando um pagamento:
```go
response, mercadopagoErr, err := mercadopago.UpdatePayment("825927174-5423394f-06f1-4d2b-8545-35ebecf70008", mercadopago.PaymentRequest{
    ExternalReference: "seu-id-interno-0001",
    Items: []mercadopago.Item{
        {
            Title:     "Pagamendo semestralidade PagueTry",
            Quantity:  1,
            UnitPrice: 300,
        },
    },
    Payer: mercadopago.Payer{
        Identification: mercadopago.PayerIdentification{
            Type:   "CPF",
            Number: "12345678909",
        },
        Name:    "Eduardo",
        Surname: "De Bastiani Mior",
        Email:   "eduardo-mior@hotmail.com",
    },
})

if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

Buscando as informações de um pagamento:
```go
response, mercadopagoErr, err := mercadopago.GetPayment("825927174-5423394f-06f1-4d2b-8545-35ebecf70008")

if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

Pesquisando um pagamento:
```go
// Na pesquisa de pagamento pode ser aplicado filtro por qualquer campo/propriedade do pagamento.
// No exemplo abaixo a pesquisa é feita pelo campo external_reference que é o nosso ID interno de conrole.
response, mercadopagoErr, err := mercadopago.SearchPayments(mercadopago.PaymentSearchParams{"external_reference": "seu-id-interno-0001"})


if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

Consultando a lista de tipos de documento e identificação:
```go
identificationTypes, mercadopagoErr, err := mercadopago.GetIdentificationTypes()

if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

Consultando a lista de meios e pagamento:
```go
identificationTypes, mercadopagoErr, err := mercadopago.GetPaymentMethods()

if err != nil {
    // Erro inesperado
} else if mercadopagoErr != nil {
    // Erro retornado do MercadoPago
} else {
    // Sucesso!
}
```

## 🙋🏻‍♂️  Ajuda
O SDK atualmente possui suporte as seguintes funções:
- Criação de um pagamento
- Atualização de um pagamento
- Consulta de um pagamento
- Listagem/Busca de pagamento
- Listagem de tipos de documento de identificação do MercadoPago
- Listagem de métodos de pagamento e suas confiurações

De acordo como forem surgindo as necessídades mais funções serão implementadas no SDK. Sinta-se livre para fazer um PullRequest ou uma Issue para novas funcionalidades.
###
O SDK precisa obrigatóriamente para funcionar, de uma variavel de ambiente chamada `MERCADO_PAGO_ACCESS_TOKEN` que contém o seu Token de integração do MercadoPago. Esse Token pode ser obtido na [página "Suas Integrações" na Dashboard do painel de Desenvolvedores do MercadoPago](https://www.mercadopago.com.br/developers/panel). Para setar a variavel ambiente você pode usar a função `os.Setenv("MERCADO_PAGO_ACCESS_TOKEN", "seu-token...")` ou você pode usar um arquivo `.env` e usar um pacote para gerenciar as variaveis de ambiente, como por exemplo o [Gotenv](https://github.com/subosito/gotenv).
###
Todas as funções do SDK podém retornar um `error` genérico do GO e um `ErrorResponse` do MercadoPago. O `error` sempre relacionado a erros do GO, como por exemplo falha ao tentar dar parse em um JSON, já o `ErrorResponse` que é a Struct de erro retornada do MercadoPago, sempre esta relacionada a erros que foram retornados da API, como por exemplo quando você não envia um campo obrigatório por exemplo.
###
Após criar um pagamento, o link para efetuar o pagamento esta na posição `InitPoint`, do model `PaymentResponse`.
###
Atenção! Você deve implementar manualmente o Webhook que recebe as atualizações de Status do pagamento usando o seu Framework WEB de prefencia (lembrando que o SDK possui a Struct `WebhookResponse` que pode ajudar no recebimento dos dados).

## 📚 Documentação oficial
Para mais duvidas consulte a [documentação oficial do MercadoPago](https://www.mercadopago.com.br/developers/pt/reference).
