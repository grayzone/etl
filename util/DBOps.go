package util

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"strings"
	"time"
)

type DBOps struct {
	Db         *sql.DB
	DriverName string
	Dbname     string
	User       string
	Password   string
}

func (ops *DBOps) Init() {
	ops.Db = nil
	ops.DriverName = "postgres"
	ops.Dbname = "etl"
	ops.User = "postgres"
	ops.Password = "123456"
}

type CityLevel struct {
	City  string
	Total int
}

type ProvinceLevel struct {
	Province string
	Total    int
}

type CountryLevel struct {
	Country string
	Total   int
}

func (ops *DBOps) Open() (err error) {
	ops.Init()
	connstr := "user=" + ops.User + " password=" + ops.Password + " dbname=" + ops.Dbname + " sslmode=disable"
	ops.Db, err = sql.Open(ops.DriverName, connstr)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (ops *DBOps) Ping() (err error) {
	return ops.Db.Ping()
}

func (ops *DBOps) Close() (err error) {
	return ops.Db.Close()
}

func (ops *DBOps) AddOneCunstomer(line []string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("customer", "customerid", "name", "phone", "fax", "lastchangedate", "distributorflag", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddCunstomers(lines [][]string) error {

	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("customer", "customerid", "name", "phone", "fax", "lastchangedate", "distributorflag", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, line := range lines {
		_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), time.Now())
		if err != nil {
			log.Println(err.Error())
			return err
		}

	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddOneDevice(line []string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("device", "locationid", "customerid", "maintenanceexpirationdate", "serialnumber", "sku", "sourcesystem", "installcountrycode", "lastchangedate", "installationdate", "actualshipdate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddDevices(lines [][]string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("device", "locationid", "customerid", "maintenanceexpirationdate", "serialnumber", "sku", "sourcesystem", "installcountrycode", "lastchangedate", "installationdate", "actualshipdate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, line := range lines {
		_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.ToLower(strings.Trim(line[3], "\"")), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), time.Now())
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddOneLocation(line []string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("location", "locationid", "addressline1", "addressmodifier2", "addressmodifier3", "addressmodifier4", "city", "stateprovince", "postalcode", "countrycode", "addressmodifier1", "lastchangedate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), strings.Trim(line[12], "\""), time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddLocations(lines [][]string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("location", "locationid", "addressline1", "addressmodifier2", "addressmodifier3", "addressmodifier4", "city", "stateprovince", "postalcode", "countrycode", "addressmodifier1", "lastchangedate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, line := range lines {
		_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.ToUpper(strings.Trim(line[5], "\"")), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), strings.Trim(line[12], "\""), time.Now())
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddOneLocationRole(line []string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("locationrole", "customerid", "locationid", "locationrole", "lastchangedate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddLocationRoles(lines [][]string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("locationrole", "customerid", "locationid", "locationrole", "lastchangedate", "isdeleted", "batchnumber", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, line := range lines {
		_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), time.Now())
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddOneLog(line []string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("log_device_fail_sku", "locationid", "customerid", "maintenanceexpirationdate", "serialnumber", "sku", "sourcesystem", "installcountrycode", "installationdate", "actualshipdate", "isdeleted", "lastchangedate", "batchnumber", "errinfo", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), strings.Trim(line[12], "\""), time.Now())
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) AddLogs(lines [][]string) error {
	txn, err := ops.Db.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare(pq.CopyIn("log_device_fail_sku", "locationid", "customerid", "maintenanceexpirationdate", "serialnumber", "sku", "sourcesystem", "installcountrycode", "installationdate", "actualshipdate", "isdeleted", "lastchangedate", "batchnumber", "errinfo", "createtime"))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, line := range lines {
		//		line = split
		_, err = stmt.Exec(strings.Trim(line[0], "\""), strings.Trim(line[1], "\""), strings.Trim(line[2], "\""), strings.Trim(line[3], "\""), strings.Trim(line[4], "\""), strings.Trim(line[5], "\""), strings.Trim(line[6], "\""), strings.Trim(line[7], "\""), strings.Trim(line[8], "\""), strings.Trim(line[9], "\""), strings.Trim(line[10], "\""), strings.Trim(line[11], "\""), strings.Trim(line[12], "\""), time.Now())
		if err != nil {
			log.Println(err.Error())
			return err
		}

	}

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = txn.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (ops *DBOps) GetDeviceInContinent() ([]CountryLevel, error) {

	rows, err := ops.Db.Query("select countrycode,total from deviceinworld where countrycode != ''")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result []CountryLevel

	for rows.Next() {
		var record CountryLevel

		err := rows.Scan(&record.Country, &record.Total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ops *DBOps) GetDeviceInCountry() ([]ProvinceLevel, error) {
	rows, err := ops.Db.Query("select stateprovince, sum(total) as total from deviceinus group by stateprovince order by stateprovince")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result []ProvinceLevel

	for rows.Next() {
		var record ProvinceLevel

		err := rows.Scan(&record.Province, &record.Total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ops *DBOps) GetDeviceInProvince(province string) ([]CityLevel, error) {
	sql := fmt.Sprintf("select city, sum(total) as total from deviceinus where stateprovince = '%s' group by city order by city", province)
	//	log.Println(sql)
	rows, err := ops.Db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result []CityLevel

	for rows.Next() {
		var record CityLevel

		err := rows.Scan(&record.City, &record.Total)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ops *DBOps) GetCityListInProvince() ([][]string, error) {
	rows, err := ops.Db.Query("select city,stateprovince from deviceinus")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result [][]string

	for rows.Next() {
		var record []string
		var city string
		var province string
		err := rows.Scan(&city, &province)
		if err != nil {
			log.Fatal(err)
		}
		record = append(record, city)
		record = append(record, province)

		result = append(result, record)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ops *DBOps) UpdateDeviceNumByDeviceType(city [][]string) error {

	for _, v := range city {
		d1 := ops.GetDeviceNumInCity(v[0], v[1], 1)
		d2 := ops.GetDeviceNumInCity(v[0], v[1], 2)
		d3 := ops.GetDeviceNumInCity(v[0], v[1], 3)
		d4 := ops.GetDeviceNumInCity(v[0], v[1], 4)

		log.Printf("%v : %d, %d, %d, %d\n", v, d1, d2, d3, d4)
	}
	return nil
}

func (ops *DBOps) UpdateDeviceNumByDeviceTypeInOneCity(city string, province string) error {
	d1 := ops.GetDeviceNumInCity(city, province, 1)
	d2 := ops.GetDeviceNumInCity(city, province, 1)
	d3 := ops.GetDeviceNumInCity(city, province, 1)
	d4 := ops.GetDeviceNumInCity(city, province, 1)
	return nil
}

func (ops *DBOps) GetDeviceNumInCity(city string, province string, devicetype int) int {
	s := fmt.Sprintf("select count(serialnumber) from deviceindmp where city = '%s' and stateprovince='%s' and devicetype = %d", city, province, devicetype)
	log.Println(s)
	var result int
	err := ops.Db.QueryRow(s).Scan(&result)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
		result = 0
	case err != nil:
		log.Fatal(err)
	default:
		//		fmt.Printf("Username is %s\n", username)
	}

	return result
}
