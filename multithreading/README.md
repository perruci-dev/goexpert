# Requisitos:

- [x] Realizar duas requisições simultaneamente para as seguintes APIs:
https://brasilapi.com.br/api/cep/v1/01153000 + cep
http://viacep.com.br/ws/" + cep + "/json/
- [x] Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- [x] O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- [x] Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.