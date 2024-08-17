package controllers

import (
	"net/http"
	"package-api/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var dbP *gorm.DB

func GetPackages(c echo.Context) error {
	var packages []models.Package
	dbP.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
	id := c.Param("id")
	var pack models.Package
	dbP.First(&pack, id)
	return c.JSON(http.StatusOK, pack)
}

func CreatePackage(c echo.Context) error {
	pack := new(models.Package)
	if err := c.Bind(pack); err != nil {
		return err
	}

	dbP.Create(&pack)
	return c.JSON(http.StatusCreated, pack)
}

func UpdatePackage(c echo.Context) error {
	id := c.Param("id")
	var pack models.Package

	// Find the package by ID
	if err := dbP.First(&pack, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Package not found")
	}

	if err := c.Bind(&pack); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data")
	}

	dbP.Save(&pack)
	return c.JSON(http.StatusOK, pack)
}

func DeletePackage(c echo.Context) error {
	id := c.Param("id")
	var pack models.Package

	if err := dbP.First(&pack, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Package not found")
	}

	dbP.Delete(&pack)
	return c.JSON(http.StatusOK, "Package deleted successfully")
}
