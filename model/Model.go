package model

import (
	"Covid-GO-API/constanta"
	"Covid-GO-API/settings"
	"fmt"
	//"net/http"
)

const ZONA_HIJAU = 1
const ZONA_KUNING = 2
const ZONA_MERAH = 3
const ZONA_HITAM = 4

type Model struct {
	db settings.DatabaseConfig
}

func zona_kecamatan(jumlah_positif int) int {
	if jumlah_positif < 1 {
		return ZONA_HIJAU
	} else if jumlah_positif > 0 && jumlah_positif < 17 {
		return ZONA_KUNING
	} else if jumlah_positif > 16 && jumlah_positif < 2049 {
		return ZONA_MERAH
	} else {
		return ZONA_HITAM
	}
}

// Insert DB
func (ExampleModel Model) InsertCovidDb(covid Covid) bool {

	sqlStatement := "INSERT INTO daerah_kecamatan (kecamatan, jumlah_penduduk_positif, jumlah_penduduk_pulih, jumlah_penduduk_wafat, keadaan_zona) " +
		"VALUES ($1, $2, $3, $4, $5)"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		covid.Kecamatan,
		covid.Jumlah_penduduk_positif,
		covid.Jumlah_penduduk_pulih,
		covid.Jumlah_penduduk_wafat,
		zona_kecamatan(covid.Jumlah_penduduk_positif),
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return true
	}

	return false
}

type Covid struct {
	Kecamatan               string `json:"kecamatan"`
	Jumlah_penduduk_positif int    `json:"jumlah_penduduk_positif"`
	Jumlah_penduduk_pulih   int    `json:"jumlah_penduduk_pulih"`
	Jumlah_penduduk_wafat   int    `json:"jumlah_penduduk_wafat"`
}

type CovidData struct {
	Id                      string `json:"id"`
	Kecamatan               string `json:"kecamatan"`
	Jumlah_penduduk_positif int    `json:"jumlah_penduduk_positif"`
	Jumlah_penduduk_pulih   int    `json:"jumlah_penduduk_pulih"`
	Jumlah_penduduk_wafat   int    `json:"jumlah_penduduk_wafat"`
	Keadaan_zona            string `json:"keadaan_zona"`
}

type CovidDataTotal struct {
	Total_penduduk_positif int `json:"total_penduduk_positif"`
	Total_penduduk_pulih   int `json:"total_penduduk_pulih"`
	Total_penduduk_wafat   int `json:"total_penduduk_wafat"`
}

type Covids struct {
	Total  CovidDataTotal `json:"totalCovid"`
	Covids []CovidData    `json:"covids"`
}

// GET From DB
func (ExampleModel Model) GetCovidDb() Covids {

	sqlStatement := "SELECT daerah_kecamatan.id, daerah_kecamatan.kecamatan, daerah_kecamatan.jumlah_penduduk_positif, daerah_kecamatan.jumlah_penduduk_pulih, daerah_kecamatan.jumlah_penduduk_wafat, zona_daerah.zona_daerah FROM daerah_kecamatan " +
		"INNER JOIN zona_daerah ON zona_daerah.id = daerah_kecamatan.keadaan_zona " +
		"ORDER BY daerah_kecamatan.id ASC"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()
	result := Covids{}
	totalCovid := CovidDataTotal{}

	for res.Next() {
		covid := CovidData{}
		err2 := res.Scan(&covid.Id, &covid.Kecamatan, &covid.Jumlah_penduduk_positif, &covid.Jumlah_penduduk_pulih, &covid.Jumlah_penduduk_wafat, &covid.Keadaan_zona)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
		totalCovid.Total_penduduk_positif += covid.Jumlah_penduduk_positif
		totalCovid.Total_penduduk_pulih += covid.Jumlah_penduduk_pulih
		totalCovid.Total_penduduk_wafat += covid.Jumlah_penduduk_wafat

		result.Covids = append(result.Covids, covid)
	}
	result.Total = totalCovid

	return result
}

// GET From DB
func (ExampleModel Model) GetCovidIdDb(id int) CovidData {

	sqlStatement := "SELECT daerah_kecamatan.id, daerah_kecamatan.kecamatan, daerah_kecamatan.jumlah_penduduk_positif, daerah_kecamatan.jumlah_penduduk_pulih, daerah_kecamatan.jumlah_penduduk_wafat, zona_daerah.zona_daerah FROM daerah_kecamatan " +
		"INNER JOIN zona_daerah ON zona_daerah.id = daerah_kecamatan.keadaan_zona " +
		"WHERE daerah_kecamatan.id = $1 " +
		"ORDER BY daerah_kecamatan.id ASC"

	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()

	covid := CovidData{}

	for res.Next() {
		err2 := res.Scan(&covid.Id, &covid.Kecamatan, &covid.Jumlah_penduduk_positif, &covid.Jumlah_penduduk_pulih, &covid.Jumlah_penduduk_wafat, &covid.Keadaan_zona)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	return covid
}

// GET From DB
func (ExampleModel Model) GetCovidPageDb(id int) Covids {

	sqlStatement := "SELECT daerah_kecamatan.id, daerah_kecamatan.kecamatan, daerah_kecamatan.jumlah_penduduk_positif, daerah_kecamatan.jumlah_penduduk_pulih, daerah_kecamatan.jumlah_penduduk_wafat, zona_daerah.zona_daerah FROM daerah_kecamatan " +
		"INNER JOIN zona_daerah ON zona_daerah.id = daerah_kecamatan.keadaan_zona " +
		"ORDER BY daerah_kecamatan.id ASC " +
		"OFFSET $1 " +
		"LIMIT $2"

	var pageNumber int
	if id != 0 {
		pageNumber = id
	} else {
		pageNumber = 1
	}

	totalRecords := 10

	offset := (pageNumber - 1) * totalRecords

	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		offset,
		totalRecords)
	if err != nil {
		fmt.Println(err)
	}

	getCovid := ExampleModel.GetCovidDb()
	defer res.Close()

	result := Covids{}

	for res.Next() {
		covid := CovidData{}
		err2 := res.Scan(&covid.Id, &covid.Kecamatan, &covid.Jumlah_penduduk_positif, &covid.Jumlah_penduduk_pulih, &covid.Jumlah_penduduk_wafat, &covid.Keadaan_zona)
		// Exit if we get an error
		if err2 != nil {
			fmt.Println(err2)
		}

		result.Covids = append(result.Covids, covid)
	}
	result.Total = getCovid.Total

	return result
}

type InfoCovid struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DetailInfoCovid struct {
	DetailTitle string      `json:"detail_title"`
	InfoCovids  []InfoCovid `json:"info_covid"`
}

// GET From DB
func (ExampleModel Model) GetCovidInfoDb() DetailInfoCovid {
	info := DetailInfoCovid{}
	covidInfo := InfoCovid{}
	covidInfo.Title = constanta.TITLE_INFO_1
	covidInfo.Description = constanta.DESCRIPTION_INFO_1
	info.InfoCovids = append(info.InfoCovids, covidInfo)

	covidInfo.Title = constanta.TITLE_INFO_2
	covidInfo.Description = constanta.DESCRIPTION_INFO_2
	info.InfoCovids = append(info.InfoCovids, covidInfo)

	covidInfo.Title = constanta.TITLE_INFO_3
	covidInfo.Description = constanta.DESCRIPTION_INFO_3
	info.InfoCovids = append(info.InfoCovids, covidInfo)

	covidInfo.Title = constanta.TITLE_INFO_4
	covidInfo.Description = constanta.DESCRIPTION_INFO_4
	info.InfoCovids = append(info.InfoCovids, covidInfo)

	info.DetailTitle = constanta.TITLE_INFO

	return info
}

// Update DB
func (ExampleModel Model) UpdateCovidDb(id int, covid Covid) bool {

	sqlStatement := "UPDATE daerah_kecamatan " +
		"SET kecamatan = $1, jumlah_penduduk_positif = $2, jumlah_penduduk_pulih = $3, jumlah_penduduk_wafat = $4, keadaan_zona = $5" +
		"WHERE id = $6"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		covid.Kecamatan,
		covid.Jumlah_penduduk_positif,
		covid.Jumlah_penduduk_pulih,
		covid.Jumlah_penduduk_wafat,
		zona_kecamatan(covid.Jumlah_penduduk_positif),
		id,
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return true
	}

	return false
}

// Delete From DB
func (ExampleModel Model) DeleteCovidDb(id int) bool {

	sqlStatement := "DELETE FROM daerah_kecamatan " +
		"WHERE id = $1"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		id,
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return true
	}

	return false
}
