# DesafioQ2Bank
## _Desafio Q2BANK: API de Transação entre usuários

DesafioQ2Bank, é um teste técnino oferecido pela Quero 2 Pay como step de um processo seletivo de Desenvolvedor Backend Pleno/Senior.

A aplicação API tem o objetivo de fazer transações entre usuários

## Decisões técnincas

As decisões tomadas para desenvolver esta aplicação foi de criar uma aplicação no modelo "scaffold" para que assim o projeto possa evoluir futuramente.

## Instalação

A aplicação necessita de um ambiente com [Golang](https://go.dev/doc/install) 1.18+ para rodar.


### Inicializar o projeto - Passo 1

Alterar o arquivo connection.go (api/db) e inserir as credencias do seu banco local "db, err := sql.Open("mysql", "usuario:senha@tcp(localhost)/dbname")"

Rodar os comandos do arquivo init.sql (api/db) no seu mysql local.

Rodar o comando:
- go run main.go


### Passo 2:
Utilizando o Postman: 
- http://localhost:9000/getUserID (Metodo GET para buscar um usuário)
Insira o ID do usuário

- http://localhost:9000/createUser (Metodo Post para criar um usuario)
Insira o Nome, Type, Cpf_Cnpj, Email, Password e Balance para criar um usuário

- http://localhost:9000/getTransactions (Metodo GET para buscar uma transação)
Insira o ID do user, para buscar transações

- http://localhost:9000/transaction (Metodo POST para criar uma transação)
Insira o IDPayer, IDPayee e Valor da transferência


## Estrutura de Pastas

- db = Conexão do banco de dados e script para criar o mysql
- transaction = contém todos os arquivos relacionados a transação
- user = contém todos os arquivos relacionados ao usuário
    - models = contém a struct transaction/user e as query para table transaction/user
    - repository = comunicação das funções com os dados do banco
    - service = serviços, funções com as regras de negócio
    - transport = decode dos dados de requisição para o código

## Observação

Não está funcionado os métodos POST e GET, estou procurando os possíveis erros para corrigir e deixar uma aplicação funcional

