package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/repositories"
)

func UpdateProductUC(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Use Case"

	status, response = repositories.UpdateProductRepositories(body, user, pathParams)

	return status, response
}