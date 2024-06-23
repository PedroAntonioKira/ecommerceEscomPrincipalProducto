package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
//	"strconv"

	"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/queries_category"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func GetProductRepositories(body string, request events.APIGatewayProxyRequest, pathParams int) (int, string) {

	status := 200
	response := "Vacio"


	fmt.Println("Entramos a GetCategoryRepositories")
	status, response = queries_category.GetProductQuery(body, request, pathParams)


	return status, response
}