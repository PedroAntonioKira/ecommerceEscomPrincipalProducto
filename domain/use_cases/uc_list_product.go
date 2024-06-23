package use_cases

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"encoding/json"
//	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/repositories"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListProductUC(body string, request events.APIGatewayProxyRequest) (int, string) {

	status := 200
	response := "Vacio"


	status, response = repositories.ListProductRepositories(body, request)


	return status, response
}