# Go Cleanarch: Instruções para execução

## Banco de dados e RabbitMQ
Antes de executar, tenha certeza de que os serviços do mysql e do rabbitmq estão em execução. Para isto, utilize o comando 

`docker compose up -d`

na pasta raiz do projeto. Isto fará com que os containeres correspondentes a este serviço estejam operacionais.

## Variáveis de ambiente
Copie os conteúdos do arquivo `.env.example` para um arquivo `.env`, na raiz do projeto. Estas são as variáveis necessárias para a conexão com o banco de dados e com a fila, além de definirem as portas nas quais cada servidor será executado. Mude os valores de portas conforme desejar.

## Execução da aplicação
Na pasta raiz, execute o seguinte comando no terminal:

`go run cmd/ordersystem/wire_gen.go cmd/ordersystem/main.go`

Isto fará com que a aplicação entre em estado de execução, disponibilizando os servidores web, gRPC e GraphQL em suas respectivas portas (definidas no seu arquivo .env).

## Recomendações
- Acesso via GraphQL: Pelo playground via web, acessando a URL do servidor;
- Acesso via gRPC: Utilitário [evans](https://github.com/ktr0731/evans);
- Acesso via Web: Instale o [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) no seu VSCode, ou utilize o [Postman](https://www.postman.com/)