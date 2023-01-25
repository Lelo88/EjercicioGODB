package store

import (
	"database/sql"
	//"github.com/go-sql-driver/mysql"
	"github.com/Lelo88/EjercicioGODB.git/internal/domain"
)

type MySqlRepository struct {
	Database *sql.DB
}

func NewSQLStore (db *sql.DB) StoreInterface {
	return &MySqlRepository{
		Database: db,
	}
}

func (repository *MySqlRepository) Read(id int) (domain.Product, error){
	
	var product domain.Product

	query := "SELECT * FROM products WHERE id = ?;"
	
	row := repository.Database.QueryRow(query, id)
	err := row.Scan(&product.Id, 
					&product.Name, 
					&product.Quantity,
					&product.CodeValue,
					&product.IsPublished,
					&product.Expiration,
					&product.Price)

	if err!= nil {
		return domain.Product{}, err
	}
	
	return product,nil
}

func (repository *MySqlRepository) Create(product domain.Product) error{
	query:="INSERT INTO products (name,quantity,code_value,is_published,expiration,price) VALUES (?,?,?,?,?,?);"
	stmt, err := repository.Database.Prepare(query)
	
	if err!=nil{
		return err
	}


	result, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product)
	if err!=nil{
		return err
	}

	_, err = result.RowsAffected()

	if err!=nil{
		return err
	}
	
	return nil
}

func (repository *MySqlRepository) Update(product domain.Product) error{

	return nil
}

func (repository *MySqlRepository) Delete(id int) error{

	return nil
}

func (repository *MySqlRepository) Exists(code_value string) bool{

	return false
}