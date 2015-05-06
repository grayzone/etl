package util

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"time"
	"strings"

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

	_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), time.Now())
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
		_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), time.Now())
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

	_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), time.Now())
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
		_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), time.Now())
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

	_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), strings.Trim(line[12],"\""), time.Now())
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
		_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), strings.Trim(line[12],"\""), time.Now())
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

	_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), time.Now())
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
		_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""),  time.Now())
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

	_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), strings.Trim(line[12],"\""), time.Now())
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
		_, err = stmt.Exec(strings.Trim(line[0],"\"") , strings.Trim(line[1],"\""), strings.Trim(line[2],"\""), strings.Trim(line[3],"\""), strings.Trim(line[4],"\""), strings.Trim(line[5],"\""), strings.Trim(line[6],"\""), strings.Trim(line[7],"\""), strings.Trim(line[8],"\""), strings.Trim(line[9],"\""), strings.Trim(line[10],"\""), strings.Trim(line[11],"\""), strings.Trim(line[12],"\""), time.Now())
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