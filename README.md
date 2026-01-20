<div align="center">

# AbacatePay Go Types

Tipagens oficiais e helpers modernos para integrar com a API da AbacatePay.

<img src="https://res.cloudinary.com/dkok1obj5/image/upload/v1767631413/avo_clhmaf.png" width="100%" alt="AbacatePay Open Source"/>

## Instalação

```bash
go get github.com/AbacatePay/go-types@latest
```

## Como a AbacatePay API Types documenta

Antes de tudo, você deve específicar a versão da API que você deseja importar os tipos. Coloque `/v*` no final da importação, sendo `*` a versão que deseja usar:

</div>

```go
import "github.com/AbacatePay/go-types/v2"
```

- Prefixo `API*`
Representa estruturas gerais da API (Objetos retornados, modelos internos etc.).

- Prefixo `Webhook*`
Representa payloads recebidos pelos eventos de webhook.
Documentação: https://docs.abacatepay.com/pages/webhooks

- Prefixo `REST<HTTPMethod>*`
Tipos usados em requisições diretas à API.
  - Sufixo Body → corpo enviado na requisição
  Ex.: `RESTPostCreateNewChargeBody`

  - Sufixo `QueryParams` → parâmetros de query
  Ex.: `RESTGetCheckQRCodePixStatusQueryParams`

  - Sufixo `Data` → dados retornados pela API
  Ex.: `RESTGetListCouponsData`

- O pacote **NÃO adiciona tipos além do que existe na documentação oficial**.
Cada tipo reflete exatamente o que está documentado aqui: https://docs.abacatepay.com/pages/introduction

<div align="center">

## Quickstart

**Crie um novo cupom**

</div>

```go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	types "github.com/AbacatePay/go-types/v2"
)

func CreateCoupon(body types.RESTPostCreateCouponBody) (*types.APICoupon, error) {
	url := types.APIBaseURL + types.RouteCreateCoupon

	payload, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	req.Header.Set("Authorization", "...")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result v1.RESTPostCreateCouponData

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data, nil
}
```

<div align="center">

**Escute eventos de Webhooks da AbacatePay**
</div>

```go
package main

import (
	"encoding/json"
	"net/http"

	types "github.com/AbacatePay/go-types/v2"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var event types.WebhookEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	switch event.Event {
	case types.WebhookEventBillingPaid:
		// Pagamento confirmado
	case types.WebhookEventPayoutDone:
		// Saque concluído
	case types.WebhookEventPayoutFailed:
		// Saque falhou
	}

	w.WriteHeader(http.StatusOK)
}
```

<p align="center">Feito com 🥑 pela equipe AbacatePay
</br>Open source, de verdade.</p>
