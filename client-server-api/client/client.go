package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	valorCotacao := pedirCotacao()
	salvarCotacao(valorCotacao)
}

func pedirCotacao() string {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Tempo do contexto excedido")
		}
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Fatal(err)
	}

	println("Cotacao: " + cotacao.Bid)
	return cotacao.Bid
}

func salvarCotacao(valor string) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte("Dólar: " + valor))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cotação salva com sucesso!")
	defer f.Close()
}
