# URL Shortener
>
> Projeto baseado no Desafio Encurtador de URL da [Devgym](https://app.devgym.com.br/challenges/3ecd0771-981d-44dc-9eee-5ec69791a745)
>
## Tecnologias usadas

- Golang - 1.20
- PostgreSQL
- Docker

## Como rodar esse projeto de forma local

1. Clonar esse repositório pelo terminal
   - Via HTTP
     `git clone https://github.com/RickChaves29/url_shortener.git`
   - Via SSH
     `git clone git@github.com:RickChaves29/url_shortener.git`
2. Ainda no terminal, copie a variável de ambiente que está no arquivo .env.example e cole no arquivo .bashrc ou .profile adicionando a palavra chave **export** antes.

   > OBS: O Arquivo **.bashrc** fica na pasta raiz do seu úsuario

   - Exemplo no WSL ou Linux:

     ```bash
     export CONNECT_DB='postgres://user:password@host:port/dbname?sslmode=disable'
     ```

3. Voltando para basta onde você clonou o projeto rode os seguintes comandos:

   - Baixar todas as dependências `go mod download`
   - Rodar o projeto `go run cmd/main.go`

## Como rodar o projeto apartir da imagem **Docker**

1. Puxe a imagem no [Docker Hub](https://hub.docker.com/r/rickchaves29/url_shortener)

    `docker pull rickchaves29/url_shortener:<tag de versão>`

2. Crie um container baseado na imagem

    ```bash
    docker run --name 'name of container' -e CONNECT_DB='url from connect from database' \
    -p 3030:3030 rickchaves29/url_shortener:'tag version' 
    ```

## Rotas da API

### POST - /api/code

Recebe um json:

```json
{
  "originUrl": "http://example.com"
}
```

Retorna um json:

```json
{
  "hashUrl": "000000"
}
```

### GET - /api/code/:hashUrl

Recebe no caminho /api/code/**000000** o número de hashUrl gerado, e redireciona o úsuario para url original
