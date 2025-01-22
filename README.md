## Desafio GoLang Pós GoExpert - Clean Architecture

### Setup
1. Clone o repositório
2. No terminal, execute o comando `docker-compose up -d` para subir o banco de dados
3. No terminal, execute o comando `go mod tidy` para baixar as dependências
4. Caminhe até a pasta `cmd/ordersystem`
5. No terminal, execute o comando `go run main.go wire_gen.go` para rodar a aplicação
6. Para validar o _usecase_ de listagem de _orders_, crie alguns registros diretamente no banco de dados ou utilize as rotas do _usecase_ de criação de _orders_ 


### Endpoints

#### REST porta :8080
Utilize o arquivo `api/list_orders.http` 

#### gRPC porta :50051
Utilize o evans com os comandos abaixo para listagem via gRPC
```shell
$ evans -r repl
$ package pb
$ service OrderService
$ call ListOrders
```
#### GraphQL porta :8080
1. Em seu browser, abra o link http://localhost:8080/
2. Cole a _query_ abaixo e execute:
```graphql
query queryOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

