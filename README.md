# TCP Connection

## O que é o projeto ?

O projeto consiste em uma servido TCP, que fica esperando um cliente TCP se conectar e permite que o cliente execute duas ações LIST e RELAY, no qual:

- LIST é para listar todos os IDs dos clientes TCP conectados
- RELAY retorna para todos os clientes conectados, os dados enviados por um cliente

## Como rodar esse projeto de forma local ?

1. Clone esse repositorio
    - Via HTTP `git clone  https://github.com/RickChaves29/tcp_connection.git`
    - Via SSH `git@github.com:RickChaves29/tcp_connection.git`

2. Ainda no terminal, copie a variável de ambiente que está no arquivo .env.example e cole no arquivo .bashrc ou .profile adicionando a palavra chave export antes.

    > OBS: O Arquivo .bashrc fica na pasta raiz do seu úsuario
    - Exemplo no WSL ou linux

    ```bash
    export CONNECT_DB='<numero da porta>'
    ```

3. Voltando para pasta onde você clonou o projeto rode os seguintes comandos:

     - Baixar todas as dependências `go mod download`
     - Rodar o projeto `go run server/main.go`

    >OBS: caso não tenha setado a variavel de ambiente use o comando

    `export SERVER_PORT='<numero da porta>' && go run server/main.go`

## Como rodar o projeto apartir da imagem **Docker**

1. Puxe a imagem no [Docker Hub](https://hub.docker.com/r/rickchaves29/tcp_server)

    `docker pull rickchaves29/tcp_server:<tag de versão>`

2. Crie um container baseado na imagem

    ```bash
    docker run --name 'name of container' -e SERVER_PORT='<numero da porta>' \
    -p 4040:'<numero da porta>'/tcp rickchaves29/tcp_server:'tag version'
    ```
