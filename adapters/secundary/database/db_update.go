package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/entities"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func UpdateProductDatabase(p entities.Product) error {
	fmt.Println(" Comienza Update Product")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "Update products SET "

	//Verificamos si nos pasaron como parametro a actualizar el titulo del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_Title", "S", 0, 0, p.ProdTitle)
	//Verificamos si nos pasaron como parametro a actualizar la descripción del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_Description", "S", 0, 0, p.ProdDescrition)
	//Verificamos si nos pasaron como parametro a actualizar el precio del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_Price", "F", 0, p.ProdPrice, "")
	//Verificamos si nos pasaron como parametro a actualizar la categoria del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_CategoryId", "N", p.ProdCategId, 0, "")
	//Verificamos si nos pasaron como parametro a actualizar el stock del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_Stock", "N", p.ProdStock, 0, "")
	//Verificamos si nos pasaron como parametro a actualizar la ruta del producto.
	sentencia = utils.ArmoSentencia(sentencia, "Prod_Path", "S", 0, 0, p.ProdPath)

	//Terminamos la sentencia indicando el id del registro que se va actualizar
	sentencia += "WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	//Ejecutamos la sentencia SQL
	_, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update Produc > Ejecución Exitosa!")

	return nil

}