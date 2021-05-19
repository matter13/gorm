package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Barang struct {
	//gorm.Model
	Id_barang   int    `gorm:"primaryKey"`
	Nama_barang string `gorm:"column:nama_barang"`
	Deskripsi   string `gorm:"column:deskripsi"`
	Stok        int    `gorm:"column:stok"`
	Harga       int    `gorm:"column:harga"`
	Tgl_masuk   string `gorm:"column:tgl_masuk"`
}

type Tabler interface {
	TableName() string
}

func (Barang) TableName() string {
	return "barang"
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", Indexz)

	e.Logger.Fatal(e.Start(":7000"))

}

func Indexz(c echo.Context) error {
	con := "root@tcp(127.0.0.1:3306)/db_kasir?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})
	if err != nil {
		panic("Tidak Terkoneksi dengan Database")
	} else {
		fmt.Println("Database Terkoneksi")
	}

	var barang []Barang

	// mengambil semua data
	result := db.Find(&barang)
	// SELECT * FROM barang;

	//result.RowsAffected // returns data berupa count
	//result.Error //returns data error

	return c.JSON(http.StatusOK, result)
}
