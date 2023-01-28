package store

import (
	"database/sql"
	"errors"

	"github.com/Lelo88/EjercicioGODB.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
	
	StoreInterface

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
	query:="INSERT INTO products (id,name,quantity,code_value,is_published,expiration,price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := repository.Database.Prepare(query)
	
	if err!=nil{
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.Id,product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err!=nil{
		return err
	}

	if err != nil {
        driverErr, ok := err.(*mysql.MySQLError)
        if !ok {
            return err
        }
		switch driverErr.Number {
        case 1062:
            return errors.New(driverErr.Message)
        default:
            return errors.New("error por otra cosa")
        }
	}
	

	rowsAffected , err := result.RowsAffected()

	if err != nil{
		return err
    }

	if rowsAffected != 1 {
		return err
	}

    productID, err := result.LastInsertId()
    if err != nil {
        return err
    }
	
	product.Id = int(productID)
	return nil
}

func (repository *MySqlRepository) Update(product domain.Product) error{

	return errors.New("error update product")
}

func (repository *MySqlRepository) Delete(id int) error{
	
	return errors.New("error delete product")
	
}

func (repository *MySqlRepository) Exists(code_value string) bool{
	
	var exists bool
    var id int
    query := "SELECT id FROM products WHERE code_value = ?;"
    row := repository.Database.QueryRow(query, code_value)
    err := row.Scan(&id)
    if err != nil {
        return false
    }
    if id > 0 {
        exists = true
		return exists
    }
    return exists
	
}