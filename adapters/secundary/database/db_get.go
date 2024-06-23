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
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/adapters/secundary"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipalProducto/domain/entities"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/models"
//	"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
	//"github.com/PedroAntonioKira/EcommerceEscomAPIREST/tools"
)

func GetProductDatabase(p entities.Product, choice string, page int, pageSize int, orderType string, orderField string) (entities.Product, error) {

	fmt.Println("Comienza SelectProduct")

	//var Resp entities.ProductResp
	var Prod []entities.Product // Con esto devolveremos una colección de estructuras

	//Nos conectamos a la base de datos
	err := secundary.DbConnect()

	//Verificamos que no hayamos tenido un error para conectarnos a la base de datos.
	if err != nil {
		return p, err
	}

	// Generamos un "defer" para que se cierre la conexión a la base de datos hasta el final de la función
	defer secundary.Db.Close()

	var sentencia string

	sentencia = "SELECT Prod_Id, Prod_Title, Prod_Description, Prod_CreatedAt, Prod_Updated, Prod_Price, Prod_Path, Prod_CategoryId, Prod_Stock FROM products WHERE Prod_Id = " + strconv.Itoa(p.ProdId)

	//Imprimimos la sentencia SQL
	fmt.Println(sentencia)

	var rows *sql.Rows

	rows, err = secundary.Db.Query(sentencia)

	for rows.Next() {
		var p entities.Product
		var ProdId sql.NullInt32
		var ProdTitle sql.NullString
		var ProdDescription sql.NullString
		var ProdCreateAt sql.NullString // Cambiamos de NullTime a NullString
		var ProdUpdated sql.NullString
		var ProdPrice sql.NullFloat64
		var ProdPath sql.NullString
		var ProdCategId sql.NullInt32
		var ProdStock sql.NullInt32

		err := rows.Scan(&ProdId, &ProdTitle, &ProdDescription, &ProdCreateAt, &ProdUpdated, &ProdPrice, &ProdPath, &ProdCategId, &ProdStock)
		if err != nil {
			return p, err
		}

		p.ProdId = int(ProdId.Int32)
		p.ProdTitle = ProdTitle.String
		p.ProdDescrition = ProdDescription.String

		// Convertir el string de la fecha a time.Time
		if ProdCreateAt.Valid {
			p.ProdCreateAt = ProdCreateAt.String
		}

		// Asignar el valor de ProdUpdated.Time directamente a p.ProdUpdated si ProdUpdated es válido
		if ProdUpdated.Valid {
			p.ProdUpdated = ProdUpdated.String
		}

		p.ProdPrice = ProdPrice.Float64
		p.ProdPath = ProdPath.String
		p.ProdCategId = int(ProdCategId.Int32)
		p.ProdStock = int(ProdStock.Int32)

		Prod = append(Prod, p)
	}

	//Resp.Data = Prod

	fmt.Println(" Select Product > Ejecución Exitosa !")

	return Prod[0], nil
}