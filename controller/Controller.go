package controller

import (
	"Covid-GO-API/model"
	"Covid-GO-API/responsegraph"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Controller struct {
	model model.Model
}

// Interface
type ExampleControllerInterface interface {
	GetPostsController(c echo.Context) error
}

// Interface
type ControllerInsertDB interface {
	GetPostsDB(c echo.Context) error
}

type Covid struct {
	Kecamatan               string `json:"kecamatan"`
	Jumlah_penduduk_positif int    `json:"jumlah_penduduk_positif"`
	Jumlah_penduduk_pulih   int    `json:"jumlah_penduduk_pulih"`
	Jumlah_penduduk_wafat   int    `json:"jumlah_penduduk_wafat"`
}

// Post Data DB
func (Controller Controller) PostsCovidDB(c echo.Context) error {
	cv := new(Covid)
	if err := c.Bind(cv); err != nil {
		return err
	}

	cfm := model.Covid{
		Kecamatan:               cv.Kecamatan,
		Jumlah_penduduk_positif: cv.Jumlah_penduduk_positif,
		Jumlah_penduduk_pulih:   cv.Jumlah_penduduk_pulih,
		Jumlah_penduduk_wafat:   cv.Jumlah_penduduk_wafat,
	}

	posts := Controller.model.InsertCovidDb(cfm)

	if posts {

		res := responsegraph.ResponseGeneric{
			Status:  "Success",
			Message: "Berhasil Input Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal Input Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)

	}
}

// Get Data DB
func (Controller Controller) GetCovidDB(c echo.Context) error {
	get := Controller.model.GetCovidDb()

	res := responsegraph.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil Mendapatkan Data",
		Data:    get,
	}
	return c.JSON(http.StatusOK, res)
}

// Get Data DB
func (Controller Controller) GetCovidIdDB(c echo.Context) error {
	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err == nil {
		fmt.Println(ids)
	}

	get := Controller.model.GetCovidIdDb(ids)

	if get.Kecamatan != "" {
		res := responsegraph.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil Mendapatkan Data",
			Data:    get,
		}
		return c.JSON(http.StatusOK, res)
	}
	res := responsegraph.ResponseGeneric{
		Status:  "Error",
		Message: "Gagal Mendapatkan Data",
	}
	return c.JSON(http.StatusOK, res)
}

// Get Data DB
func (Controller Controller) GetCovidPageDB(c echo.Context) error {
	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err == nil {
		fmt.Println(ids)
	}

	get := Controller.model.GetCovidPageDb(ids)

	res := responsegraph.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil Mendapatkan Data",
		Data:    get.Covids,
	}
	return c.JSON(http.StatusOK, res)
}

// Get Data DB
func (Controller Controller) GetCovidTotalDB(c echo.Context) error {
	get := Controller.model.GetCovidDb()

	res := responsegraph.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil Mendapatkan Data",
		Data:    get.Total,
	}
	return c.JSON(http.StatusOK, res)
}

// Get Data DB
func (Controller Controller) GetCovidInfoDB(c echo.Context) error {
	get := Controller.model.GetCovidInfoDb()

	res := responsegraph.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil Mendapatkan Data",
		Data:    get,
	}
	return c.JSON(http.StatusOK, res)
}

// Update Data DB
func (Controller Controller) UpdateCovidDB(c echo.Context) error {
	cv := new(Covid)
	if err := c.Bind(cv); err != nil {
		return err
	}

	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err == nil {
		fmt.Println(ids)
	}

	cfm := model.Covid{
		Kecamatan:               cv.Kecamatan,
		Jumlah_penduduk_positif: cv.Jumlah_penduduk_positif,
		Jumlah_penduduk_pulih:   cv.Jumlah_penduduk_pulih,
		Jumlah_penduduk_wafat:   cv.Jumlah_penduduk_wafat,
	}

	updates := Controller.model.UpdateCovidDb(ids, cfm)

	if updates {

		res := responsegraph.ResponseGeneric{
			Status:  "Success",
			Message: "Berhasil Mengupdate Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegraph.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal Mengupdate Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (Controller Controller) DeleteCovidDb(c echo.Context) error {
	id := c.Param("id")
	ids, err := strconv.Atoi(id)
	if err == nil {
		fmt.Println(ids)
	}

	deletes := Controller.model.DeleteCovidDb(ids)

	if deletes {
		res := responsegraph.ResponseGeneric{
			Status:  "Success",
			Message: "Berhasil Menghapus Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res := responsegraph.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal Menghapus Data Kecamatan",
		}
		return c.JSON(http.StatusOK, res)
	}
}
