# Go Chat Backend

Este é um projeto de backend para um chat em tempo real, implementado em Go. O servidor utiliza WebSockets para comunicação bidirecional entre clientes e permite que mensagens sejam enviadas de um usuário para outro.

## Funcionalidades

- Conexão de múltiplos clientes via WebSocket.
- Envio de mensagens entre usuários.
- Identificação de usuários através de nomes de usuário.

## Pré-requisitos

- Go 1.23 ou superior
- Dependências do Go (instaladas automaticamente com `go mod`)

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu_usuario/go-chat-backend.git
   cd go-chat-backend
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Execute o servidor:

   ```bash
   go run main.go
   ```

4. O servidor estará disponível em `http://localhost:8080/ws`.

## Uso

Para se conectar ao servidor WebSocket, você pode usar um cliente WebSocket ou um navegador. A URL de conexão deve incluir o nome de usuário como um parâmetro de consulta:

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir um problema ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
