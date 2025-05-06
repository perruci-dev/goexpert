# Requisitos:

- [x] O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
- [x] O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON).
- [x] Utilizando o package "context",o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
- [x] Os contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
- [x] O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}