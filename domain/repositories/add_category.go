package repositories

import (
	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/queries_category"
)

func AddCategoryRepositories(body string, user string) (int, string) {
	status := 200
	response := "Vacio Add Category Repositories"

	status, response = queries_category.AddCategoryQuery(body, user)

	return status, response
}