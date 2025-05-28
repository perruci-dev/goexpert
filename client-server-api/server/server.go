package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CotacaoUSDBRL struct {
	USDBRL Cotacao `json:"USDBRL"`
}

type Cotacao struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

var db *gorm.DB

func main() {
	initDB()
	http.HandleFunc("/cotacao", buscarCotacao)
	http.ListenAndServe(":8080", nil)
}

func initDB() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Table("cotacoes").AutoMigrate(&Cotacao{})
}

func buscarCotacao(w http.ResponseWriter, r *http.Request) {
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	cotacao := pedirCotacao()
	salvarCotacao(cotacao)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotacao)

}

func pedirCotacao() Cotacao {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Tempo do contexto para resposta da API excedido")
		}
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cotacaoUSDBRL CotacaoUSDBRL
	err = json.Unmarshal(body, &cotacaoUSDBRL)
	if err != nil {
		log.Fatal(err)
	}

	return cotacaoUSDBRL.USDBRL
}

func salvarCotacao(cotacao Cotacao) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := db.WithContext(ctx).Table("cotacoes").Create(cotacao).Error
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Tempo do contexto para persistir no db excedido")
		}
		log.Fatal(err)
	}
	fmt.Printf("Cotação salva com sucesso!")
}
