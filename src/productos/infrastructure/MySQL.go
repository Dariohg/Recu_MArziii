package infrastructure

import (
	"233338-R-C2/src/core"
	"233338-R-C2/src/productos/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Guardar(product *entities.Product) error {
	query := `INSERT INTO productos 
              (nombre, precio, codigo, descuento, email) 
              VALUES (?, ?, ?, ?, ?)`

	result, err := mysql.conn.ExecutePreparedQuery(query,
		product.Nombre,
		product.Precio,
		product.Codigo,
		product.Descuento,
		product.Email)

	if err != nil {
		return fmt.Errorf("error al guardar el producto: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID del producto insertado: %v", err)
	}

	product.ID = int(id)

	return nil
}

func (mysql *MySQL) ObtenerTodos() ([]entities.Product, error) {
	query := `SELECT id, nombre, precio, codigo, descuento, email, fecha_creacion 
              FROM productos
              ORDER BY fecha_creacion DESC`

	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.ID,
			&product.Nombre,
			&product.Precio,
			&product.Codigo,
			&product.Descuento,
			&product.Email,
			&product.FechaCreacion); err != nil {
			return nil, fmt.Errorf("error al escanear producto: %v", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (mysql *MySQL) ObtenerPorId(id int) (*entities.Product, error) {
	query := `SELECT id, nombre, precio, codigo, descuento, email, fecha_creacion 
              FROM productos 
              WHERE id = ?`

	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var product entities.Product
	if rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Nombre,
			&product.Precio,
			&product.Codigo,
			&product.Descuento,
			&product.Email,
			&product.FechaCreacion)
		if err != nil {
			return nil, fmt.Errorf("error al escanear producto: %v", err)
		}
		return &product, nil
	}
	return nil, fmt.Errorf("producto no encontrado")
}

func (mysql *MySQL) ObtenerUltimoProducto() (*entities.Product, error) {
	query := `SELECT id, nombre, precio, codigo, descuento, email, fecha_creacion 
              FROM productos 
              ORDER BY fecha_creacion DESC 
              LIMIT 1`

	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var product entities.Product
	if rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Nombre,
			&product.Precio,
			&product.Codigo,
			&product.Descuento,
			&product.Email,
			&product.FechaCreacion)
		if err != nil {
			return nil, fmt.Errorf("error al escanear producto: %v", err)
		}
		return &product, nil
	}
	return nil, nil // No hay productos aún
}

func (mysql *MySQL) ContarProductosConDescuento() (int, error) {
	query := `SELECT COUNT(*) FROM productos WHERE descuento = true`

	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, fmt.Errorf("error al contar productos con descuento: %v", err)
		}
	}

	return count, nil
}
