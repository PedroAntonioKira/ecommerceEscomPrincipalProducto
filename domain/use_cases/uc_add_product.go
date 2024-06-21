package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/repositories"
)

func AddProductUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Use Case"

	status, response = repositories.AddProductRepositories(body, user)

	return status, response
}