# digital-bank
API de transferencia entre contas Internas de um banco digital.

## Pré-requisitos
- [Git](https://git-scm.com/)
- [Docker e Docker Compose](https://www.docker.com/get-started/)
## Getting started
1. Clone o repositório
    ```
        git clone https://github.com/yanwr/digital-bank.git
   ```
2. Rodando a API 
    ```
        cd digital-bank
        docker-compose up --build
   ```

## Por dentro da API
- **/app/login**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
POST | /app/login | - | Autentica utilizando CPF e Secret. E retorna o token. |

Request | Response
--|--
 `{`<br>`"cpf": string,`<br>`"secret": string `<br>`}` | `{`<br>`"token": string`<br>`}`

- **/app/accounts**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
GET  | /app/accounts | - | Retorna todas contas. |

Request | Response
:--:|--
| - | `[`<br>`{`<br>`"id": string,`<br>`"name": string,`<br>`"cpf": "string,`<br>`"balance": number,`<br>`"created_at": time `<br>`}`<br>`]`

- **/app/accounts**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
POST | /app/accounts | - | Cria uma nova conta. |

Request | Response
--|--
 `{`<br>`"name": string,`<br>`"cpf": string,`<br>`"secret": "string,`<br>`"balance": number`<br>`}` | `{`<br>`"id": string,`<br>`"name": string,`<br>`"cpf": "string,`<br>`"balance": number,`<br>`"created_at": time `<br>`}`

- **/app/accounts/:account_id/balance**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
GET  | /app/accounts/:account_id/balance | - | Retorna o saldo da conta.

Request | Response
:--:|--
| Query Param | number

- **/app/transfers**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
GET  | /app/transfers | Bearer ${token} | Retorna a lista de transferências da conta  logada. |

Request | Response
:--:|--
| - | `[`<br>`{`<br>`"id": string,`<br>`"account_destination_id": string,`<br>`"amount": "number,`<br>`"created_at": time `<br>`}`<br>`]`

- **/app/transfers**

Método | Rota | Header Authorization | Função
:--:|--|:--:|--
POST | /app/transfers | Bearer ${token} | Faz uma transferência da conta logadoo para  uma conta informada. |

Request | Response
--|--
 `{`<br>`"account_destination_id": string,`<br>`"amount": number `<br>`}` | `{`<br>`"id": string,`<br>`"account_destination_id": string,`<br>`"amount": "number,`<br>`"created_at": time `<br>`}`