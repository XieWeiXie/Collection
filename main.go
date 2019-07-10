package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	RECYCLABLE = iota + 1
	HAZARDOUS
	HOUSEHOLDFOOD
	RESIDUAL
)
const (
	RecyclableWasteCN    = "可回收垃圾"
	HazardousWasteCN     = "有害垃圾"
	HouseholdFoodWasteCN = "湿垃圾"
	ResidualWasteCN      = "干垃圾"
)

var WasteMap = make(map[string]string)

func init() {
	WasteMap[strconv.Itoa(RECYCLABLE)] = RecyclableWasteCN
	WasteMap[strconv.Itoa(HAZARDOUS)] = HazardousWasteCN
	WasteMap[strconv.Itoa(HOUSEHOLDFOOD)] = HouseholdFoodWasteCN
	WasteMap[strconv.Itoa(RESIDUAL)] = ResidualWasteCN
}

func main() {
	//readCSV()
	//csvMap()
	write()
}

type Waste struct {
	Name  string
	Class string
}

type Wastes []Waste

func readCSV() Wastes {
	f, err := os.Open("waste.csv")
	if err != nil {
		log.Println(err)
		return nil
	}
	reader := csv.NewReader(f)
	rows, _ := reader.ReadAll()
	var ws Wastes
	for index, line := range rows {
		if index == 0 {
			continue
		}
		var one Waste
		one = Waste{
			Name:  strings.TrimSpace(line[1]),
			Class: strings.TrimSpace(WasteMap[line[3]]),
		}
		ws = append(ws, one)
	}
	//fmt.Println(ws)
	return ws
}

func csvMap() string {
	rubbish := "map[string]string{"
	var uniqueMap = make(map[string]bool)
	for index, line := range readCSV() {
		if index == 0 {
			continue
		}
		if uniqueMap[line.Name] {
			continue
		}
		uniqueMap[line.Name] = true
		rubbish += fmt.Sprintf(`"%s": "%s",`, strings.TrimSpace(line.Name), strings.TrimSpace(line.Class)) + "\n"
	}
	rubbish += "}"
	fmt.Println(rubbish)
	return rubbish
}

func write() {
	dir := "./waste/waste.go"
	if _, fErr := os.Stat(dir); os.IsExist(fErr) {
		err := os.Mkdir("waste", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	code := `package waste
var (
	Waste = map[string]string{}
)
func init(){
	Waste = %s
}
`
	codeMap := csvMap()
	goCode := fmt.Sprintf(code, codeMap)
	err := ioutil.WriteFile(dir, []byte(goCode), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	cmd := exec.Command("go", "fmt", dir)
	cmd.Start()
	return
}
