package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
//	"strconv"

	"fmt"

	//importaciones externas (descargadas)
//	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func UpdateCategoryQuery(body string, User string, pathParams int) (int, string) {
	//creamos la variable de la estructura que almacenará todo lo de la categoria relacionada
	var t entities.Product

	fmt.Println(t)

	return 200, "Update OK Very Good"

}
