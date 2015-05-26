package main

import (
	"bufio"
	"encoding/json"
	"github.com/grayzone/etl/util"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type ImportData func([][]string) error

type GeoResponse struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
}

type Result struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAddress  string             `json:"formatted_address"`
	Geometry          Geometry           `json:"geometry"`
	PlaceID           string             `json:"place_id"`
	Types             []string           `json:"types"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometry struct {
	Bounds       Area       `json:"bounds"`
	Location     Coordinate `json:"location"`
	LocationType string     `json:"location_type"`
	Viewport     Area       `json:"viewport"`
}

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Area struct {
	Northeast Coordinate `json:"northeast"`
	Southwest Coordinate `json:"southwest"`
}

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

func UpdateDeviceNumByTypeInUS() {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	result, _ := db.GetCityListInProvince()
	//	log.Println(result)
	db.UpdateDeviceNumByDeviceType(result)
}

func VerifyOldCustomerID() {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := "C:\\etl\\old.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i = i + 1

		if db.IsOldCustomerDeleted(scanner.Text()) == false {
			log.Printf("%d:%s...error\n", i, scanner.Text())
		} else {
			log.Printf("%d:%s.....ok\n", i, scanner.Text())
		}

	}

}

func VerifyNewCustomerID() {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filepath := "C:\\etl\\new.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i = i + 1

		if db.IsNewCustomerExisting(scanner.Text()) == false {
			log.Printf("%d:%s...error\n", i, scanner.Text())
		} else {
			if i%1000 == 0 {
				log.Printf("%d:%s.....ok\n", i, scanner.Text())
			}
			//			log.Printf("%d:%s.....ok\n",i,scanner.Text())
		}

	}

}

//http://maps.googleapis.com/maps/api/geocode/json?address=%E2%80%98ABBEVILLE%20LA%20US%E2%80%99
func UpdateCityCoordinate() {
	var db util.DBOps
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	googleapi := "http://maps.googleapis.com/maps/api/geocode/json?address="
	u, err := url.Parse(googleapi)
	q := u.Query()

	citylist, _ := db.GetCityListInProvince()
	var i int64
	i = 1798
	j := 0

	for _, city := range citylist {
		cityid, _ := strconv.ParseInt(city[2], 10, 64)
		log.Println(cityid)
		if cityid < i {
			continue
		}

		j = j + 1
		if j > 1000 {
			break
		}

		//		log.Println(city[2])
		param := city[0] + " " + city[1] + " " + "US"
		q.Set("address", param)
		u.RawQuery = q.Encode()
		//		log.Println(u)

		u.RawQuery = q.Encode()
		log.Println(u)
		res, err := http.Get(u.String())
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var f GeoResponse
		json.Unmarshal(body, &f)

		if f.Status != "OK" {
			continue
		}
		//		log.Println(city[2])
		//		log.Println(f.Results[0].Geometry.Location)
		loc := f.Results[0].Geometry.Location
		db.UpdateCityCoordinateByID(city[2], loc.Lat, loc.Lng)
		time.Sleep(1000 * time.Millisecond)

	}

}

func TestFunc() {

	result := 9999
	for result != 0 {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		b := rand.Perm(9)
		for i, _ := range a {
			a[i] = b[i] + 1
		}
		if (13*a[1])%a[2] == 0 && (a[6]*a[7])%a[8] == 0{
			result = a[0] + (13*a[1])/a[2] + a[3] + 12*a[4] - a[5] + (a[6]*a[7])/a[8] - 87
		}else{
			result = 9999
		}
		log.Println(result)

	}

}

func main() {

	// loadETLDataSouce()
	//loadETLlog()
	//	UpdateDeviceNumByTypeInUS()

	//	VerifyOldCustomerID()
	//	VerifyNewCustomerID()

	//	UpdateCityCoordinate()
	TestFunc()
}
