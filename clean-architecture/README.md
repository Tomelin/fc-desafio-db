# DESAFIO CLEAN ARCHITECTURE

Desafio clean architecture consiste em listar ordens cadastradas no BD postgres através de HTTP, gRPC e GraphQL.

São 3 itens cadastrados de forma automatica ao subir o docker-compose, podendo adicionar novos itens

## OVERVIEW CONFIGS
### DIRETÓRIOS

Segue uma breve descrição de cada diretorio desse micro sistema:

```
.
|-- build       # Contém Dockerfile, docker-compose e script para iniciar os itens no BD
|   |-- build
|   |   -- scripts
|   -- scripts
|-- cmd         # start do sistema
|-- configs     # carrega as configs através do arquivo config.yaml
|-- docs        # documentações, nesse caso temos apenas do OpenAPIv3
|   -- swagger
|-- internal    # diretório internal do sistema
|   |-- core    # aqui, atratamso do core, usecase
|   |   |-- entity
|   |   |-- repository
|   |   -- service
|   -- infra    # aqui é onde se encontra os handlers e storage (db, redis, ....)
|       |-- handler
|       |   |-- graphql
|       |   |   -- graph
|       |   |       -- model
|       |   |-- grpc
|       |   -- rest
|       -- storage
|           |-- cache
|           -- database
|-- pkg        # pacotes compilados por terceiros (protobuf)
|   |-- grpc
|   |   -- order_pb
|   -- rest
|       -- httpserver
-- proto       # configurações do protobuf
```

###  Makefile

O makefile, contém os atalhos para executar o sistema ou gerar arquivos tais como swagger

###  Protobuf

O diretório proto, contém informações do protobuf e é gerado através do comando **buf generate .**

### CMD

O diretório cmd contém o arquivo .confg/config.yaml, onde contempla todas as configs do sistema.

O arquivo config consiste na seguinte configuração:

```
---
database:
  host: "database"
  port: "5432"
  database: desafio
  username: admin
  password: UE9TVEdSRVNfUEFTU1dPUkQK
webserver:
  port: "8443"
  ssl_enabled: "false"
  name: "api"
  host: "0.0.0.0"
  version: ""
cache:
  port: "6379"
  host: "redis"
  password: "cmVkaXMtc2VydmVyCg"
  database: 0

```

### Docker compose

O docker compose, incia aos seguintes containers:

**app**, que precisa ter a variável PATH_CONFIG apontando para onde está o arquivo config.yaml.

**Portas expostas:**
```
http: 8443
grpc: 50051
graphql: 8082
```
;

**database**, o banco de dados postgres

**redis**, redis, que utlizamos para cache nas consultas

**redis-commander**, redis-commander é utilizado apenas para visualizar se tem as chaves no redis, serve como troubleshooting

## INICIALIZANDO

### Clone
git clone https://github.com/Tomelin/fc-desafio-db.git  
cd fc-desafio-db/clean-architecture  

### Executando docker compose
docker-compose -f build/docker-compose.yaml up -d  

### Acessando os serviços

- **HTTP:** http://localhost:8443/api
- **gRPC**: tcp://localhost:50051
- **GraphQL**: http://localhost:8082

#### GraphQL

Example de querys:

```
query Orders{
  orders{
    id
    name
  }
}

query filter{
  order(id: "b34f408d-7067-4b84-8782-3c8e5b2f893d"){
    id
    name
  }
}

```

#### gRPC

1. 127.0.0.1:50051> package pb
2. service OrderService
3. call FindByFilter
  3.1 value (TYPE_STRING) => item
