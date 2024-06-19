package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
//	"strconv"

	//"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func GetCategoryQuery(body string, request events.APIGatewayProxyRequest, pathParams int) (int, string) {
//	var err error
	var CategId int
	var Slug string // Es el path (CategoryPath), solo que se suele llamar así en un ecommerce


	//Verificamos si recibimos el "CategId" o recibimos el "Slug"
	/*
	if len(request.QueryStringParameters["categId"]) > 0 {

		CategId, err = strconv.Atoi(request.QueryStringParameters["categId"])

		//Verificamos que la conversión haya sido correcta
		if err != nil {
			return 500, "Ocurrio un error al intentar convertir en entero el valor" + request.QueryStringParameters["categId"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	*/

	CategId = pathParams
	Slug = "Dato Especifico"
	//Si no se especifica por que debemos filtrar (id o path) devolvemos todas las categorias de la Base de Datos.
	lista, err2 := database.GetCategoryQuery(CategId, Slug)

	//Validamos si no tuvimos un error al capturar categorias
	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar categoría/s > " + err2.Error()
	}

	// Obtenemos la información en un JSON
	Categ, err3 := json.Marshal(lista)

	// Verificamos no haya existido ningun error al pasar la estructura y obtener la información
	if err3 != nil {
		return 400, "Ocurrio un error al intentar capturar la Categoría/s > " + err3.Error()
	}

	return 200, string(Categ)

}