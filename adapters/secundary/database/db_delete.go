package database

import (
	//Importaciones de go (vienen incluidas al instalar)
//	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
//	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
//	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/utils"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func DeleteProductQuery(pathParams int) error {

	fmt.Println(" Comienza Delete Product")

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	//Declaramos la sentencia SQL para insertar el Producto.
	sentencia := "DELETE FROM products WHERE Prod_Id = " + strconv.Itoa(pathParams)

	//Ejecutamos la sentencia SQL
	_, err = secundary.Db.Exec(sentencia)

	//Verificamos no haya existido un error al ejecutar la sentencia SQL
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Delete Produc > Ejecución Exitosa!")

	return nil

}