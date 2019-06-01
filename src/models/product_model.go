package models

import (
	"database/sql"
)

type ProductModel struct {
	DB *sql.DB
}

func (productModel ProductModel) FindAll() (product []Product, err error) {
	rows, err := productModel.DB.Query("select * from product")
}
