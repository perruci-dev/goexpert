package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type apiURL struct {
	api string
	url string
}

func main() {
	cep := "01153000"

	urls := []apiURL{
		// {"BrasilAPIv2", fmt.Sprintf("https://brasilapi.com.br/api/cep/v2/%s", cep)},
		{"BrasilAPIv1", fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)},
		{"ViaCEP", fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string, 1)

	for _, url := range urls {
		go consultarCEP(ctx, ch, url)
	}

	resposta := <-ch
	fmt.Println(resposta)
}

func consultarCEP(ctx context.Context, ch chan string, apiURL apiURL) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case ch <- fmt.Sprintf("Resposta da API " + apiURL.api + "\n" + string(body)):
	case <-ctx.Done():
	}
}
