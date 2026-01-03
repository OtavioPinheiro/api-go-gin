package main

import (
	"api-go-gin/models"
	"api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Letícia Natália Lorena Ribeiro", CPF: "238.552.101-60", RG: "20.741.598-5"},
		{Nome: "Otávio Augusto", CPF: "858.441.271-96", RG: "41.408.741-0"},
	}
	routes.HandleRequests()
}
