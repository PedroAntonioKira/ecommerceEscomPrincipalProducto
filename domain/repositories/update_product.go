package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/queries_category"
)

func UpdateProductRepositories(body string, user string, pathParams int) (int, string) {
	status := 200
	response := "Vacio Update Category Repositories"

	status, response = queries_category.UpdateProductQuery(body, user, pathParams)

	return status, response
}