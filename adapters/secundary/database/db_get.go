package database

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"strconv"
//	"strings"

	//"strings"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalCategoria/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func GetCategoryQuery(CategId int, Slug string) ([]entities.Category, error) {
	fmt.Println("Comienza SelectCategories")

	//Creamos la variable que almacenara cada registro devuelto de cada categoria de la base de datos.
	var Categ []entities.Category

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return Categ, err
	}

	//Declaramos la sentencia SQL para insertar la categoria
	sentencia := "SELECT Categ_Id, Categ_Name, Categ_Path FROM category  WHERE Categ_Id = " + strconv.Itoa(CategId)

/*
	//Validamos si nos pidieron buscar por un ID en particular en caso que se haya especificado.
	if CategId > 0 {
		sentencia += " WHERE Categ_Id = " + strconv.Itoa(CategId)
	} else {
		// Validamos que nos permita buscar por una ruta en particular en caso que se haya especificado.
		if len(Slug) > 0 {
			sentencia += " WHERE Categ_Path LIKE '%" + Slug + "%'"
		}
	}
*/

	//Imprimimos la sentencia SQL
	fmt.Println(sentencia)

	var rows *sql.Rows

	var err003 error
	rows, err003 = secundary.Db.Query(sentencia)

	fmt.Println(" Mostramos error de  Db.Query(sentencia): ")
	fmt.Println(err003)

	for rows.Next() {
		var c entities.Category
		var categId sql.NullInt32
		var categName sql.NullString
		var categPath sql.NullString

		err := rows.Scan(&categId, &categName, &categPath)

		//Validamos que no haya surgido ningun error y en caso de existir abortamos y salimos regresando el error
		if err != nil {
			return Categ, err
		}

		c.CategID = int(categId.Int32)
		c.CategName = categName.String
		c.CategPath = categPath.String

		Categ = append(Categ, c)

	}

	fmt.Println("Select Categories > Ejecución Exitosa : Optimizado01")

	return Categ, nil
}