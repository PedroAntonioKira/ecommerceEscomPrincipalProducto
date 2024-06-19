package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/repositories"
)

func UpdateCategoryUC(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Use Case"

	status, response = repositories.UpdateCategoryRepositories(body, user, pathParams)

	return status, response
}