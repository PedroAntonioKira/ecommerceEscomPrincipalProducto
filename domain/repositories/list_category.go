package repositories

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
//	"strconv"

	"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/queries_category"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListCategoryRepositories(body string, request events.APIGatewayProxyRequest) (int, string) {

	status := 200
	response := "Vacio"


	fmt.Println("Entramos a ListCategoryRepositories")
	status, response = queries_category.ListCategoryQuery(body, request)


	return status, response
}