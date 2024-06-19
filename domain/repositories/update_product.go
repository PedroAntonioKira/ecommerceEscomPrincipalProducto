package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/queries_product"
)

func UpdateCategoryRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Repositories"

	status, response = queries_product.UpdateCategoryQuery(body, user, pathParams)

	return status, response
}