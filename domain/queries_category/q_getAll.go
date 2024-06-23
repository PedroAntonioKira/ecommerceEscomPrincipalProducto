package queries_category

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/json"
	"strconv"
	"strings"
	"fmt"

	//importaciones externas (descargadas)
	"github.com/aws/aws-lambda-go/events"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary/database"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/bd"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
)

func ListProductQuery(body string, request events.APIGatewayProxyRequest) (int, string) {

	// Variable donde almacenaremos la información recibida del producto.
	var t entities.Product

	// Variable para almacenar el error en caso de existir.
	var err error

	// "page" nos ayudará a saber el numero de pagina de productos y "pagesize" sera la cantidad de productos que entraran por pagina.
	var page, pageSize int

	// variables para ordenar "orderField" ordena por el campo especificado (Descripción, Precio, etc), "ordenType" ordena de forma ascendente o descendente.
	var orderType, orderField string

	//Variable que nos convierte a JSON los datos que manda el endpoint
	param := request.QueryStringParameters

	//Guarfamos los datos que vienen en el Json en variables.
	page, _ = strconv.Atoi(param["page"])

	pageSize, _ = strconv.Atoi(param["pageSize"])

	orderType = param["orderType"] // (D = Descendente) , (A o Nil = Ascendente).

	orderField = param["orderField"] // 'I' id, 'T' Title, 'D' Description, 'F' Created At,
	// 'P' Price, 'C' CategId, 'S' Stock

	//Valido que especifique al menos un tipo ordenamiento (en caso de existir)
	if !strings.Contains("ITDFPCS", orderField) {
		orderField = ""
	}

	var choice string

	if len(param["prodId"]) > 0 {
		fmt.Println("SI LLEGO EL prodId, SE ACTUALIZA")
		choice = "P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}

	if len(param["search"]) > 0 {
		fmt.Println("SI LLEGO EL search, SE ACTUALIZA")
		choice = "S"
		t.ProdSearch, _ = param["search"]
	}

	if len(param["categId"]) > 0 {
		fmt.Println("SI LLEGO EL categId, SE ACTUALIZA")
		choice = "C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(param["slug"]) > 0 {
		fmt.Println("SI LLEGO EL slug, SE ACTUALIZA")
		choice = "U"
		t.ProdPath, _ = param["slug"]
	}

	if len(param["slugCateg"]) > 0 {
		fmt.Println("SI LLEGO EL slugCateg, SE ACTUALIZA")
		choice = "K"
		t.ProdCategPath, _ = param["slugCateg"]
	}

	fmt.Println(param)

	result, err := database.ListProductDatabase(t, choice, page, pageSize, orderType, orderField)

	if err != nil {
		return 400, "Ocurrió un error al intentar capturar los resultados de la búsqueda de tipo '" + choice + "' en productos" + err.Error()
	}

	Product, err2 := json.Marshal(result)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON la busqueda de productos"
	}

	return 200, string(Product)
}