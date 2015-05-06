package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"util"
)

type ImportData func([][]string) error

const DATA_SOURCE_FOLDER = "C:\\etl\\datasource\\20150419"
const READ_DATA_LINE = 5000

func getDataSourceFilepath(name string) string {
	filepath := DATA_SOURCE_FOLDER + "\\" + name + "_data.csv"
	return filepath
}

func ReadDataSourceFile(filepath string, fp ImportData) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	i := 0
	j := 0
	var list [][]string
	scanner := bufio.NewScanner(file)
	for {
		if scanner.Scan() {
			if i == 0 {
				i = i + 1
				continue
			}

			j = j + 1
			items := strings.Split(scanner.Text(), "|")
			list = append(list, items)

			if j == READ_DATA_LINE {
				//		log.Printf("%d:", i)
				err = fp(list)
				if err != nil {
					log.Fatal(err)
				}
				//		fmt.Println(".....ok")
				list = nil
				j = 0
			}

			i = i + 1

		} else {
			//		log.Printf("%d:", i)
			err = fp(list)
			if err != nil {
				log.Fatal(err)
			}
			//		fmt.Println(".....ok")
			list = nil
			break
		}

	}

}

func loadCustomerDataSouce() {
	log.Println("loadCustomerDataSouce")
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := getDataSourceFilepath("customer")
	ReadDataSourceFile(filepath, db.AddCunstomers)

	log.Println("loadCustomerDataSouce  ... ok")
}

func loadDeviceDataSouce() {
	log.Println("loadDeviceDataSouce")
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := getDataSourceFilepath("device")
	ReadDataSourceFile(filepath, db.AddDevices)

	log.Println("loadDeviceDataSouce  ... ok")
}

func loadLocationDataSouce() {
	log.Println("loadLocationDataSouce")
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := getDataSourceFilepath("location")
	ReadDataSourceFile(filepath, db.AddLocations)

	log.Println("loadLocationDataSouce  ... ok")

}

func loadLocationRoleDataSouce() {
	log.Println("loadLocationRoleDataSouce")
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := getDataSourceFilepath("location_role")
	ReadDataSourceFile(filepath, db.AddLocationRoles)

	log.Println("loadLocationRoleDataSouce  ... ok")
}

func loadETLDataSouce() {
	loadCustomerDataSouce()
	loadDeviceDataSouce()
	loadLocationDataSouce()
	loadLocationRoleDataSouce()
}

func loadETLlog() {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := "C:\\etl\\log\\20150419\\etl1429477709441_DEVICE_FAIL_SKU_VALIDATE.CSV"
	ReadDataSourceFile(filepath, db.AddLogs)
}

func main() {

	loadETLDataSouce()
	//loadETLlog()
}
