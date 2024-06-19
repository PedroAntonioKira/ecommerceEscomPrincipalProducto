package use_cases

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/repositories"
)

func AddCategoryUC(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Use Case"

	status, response = repositories.AddCategoryRepositories(body, user)

	return status, response
}