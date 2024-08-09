package controllers

import (
	"assignment1/config"
	"assignment1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPackages(c echo.Context) error {
	var packages []models.Package
	config.DB.Find(&packages)
	return c.JSON(http.StatusOK, packages)
}

func GetPackageByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pkg models.Package
	if result := config.DB.First(&pkg, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}
	return c.JSON(http.StatusOK, pkg)
}

func CreatePackage(c echo.Context) error {
	var pkg models.Package
	if err := c.Bind(&pkg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}
	config.DB.Create(&pkg)
	return c.JSON(http.StatusCreated, pkg)
}

func UpdatePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pkg models.Package
	if result := config.DB.First(&pkg, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}

	if err := c.Bind(&pkg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	config.DB.Save(&pkg)
	return c.JSON(http.StatusOK, pkg)
}

func DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if result := config.DB.Delete(&models.Package{}, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Package not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
