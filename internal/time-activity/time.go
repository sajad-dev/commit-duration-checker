package timeactivity

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var sc = 0

func CheckTime(keypress chan int) int {
	sc++
	if len(keypress) != 0 {
		sc = 0
	}
	return sc
}

var second int = 0

func GetWindowWrite() string {
	cmd := exec.Command("xdotool", "getactivewindow", "getwindowname")
	out, _ := cmd.Output()
	return strings.TrimSpace(string(out))
}

type ArrMap struct {
	window string
	secend int
}

var arr []ArrMap

var Activity = make(chan []ArrMap, 10)

func HandelFile() {
	os.Chdir("static")

	var arrmap = []map[string]interface{}{}
	for _, v := range arr {
		hashmap := map[string]any{}
		hashmap["window"] = v.window
		hashmap["secend"] = strconv.Itoa(v.secend)
		hashmap["repository"] = ""
		hashmap["commit"] = ""
		arrmap = append(arrmap, hashmap)
	}

	file, _ := os.Open("data.json")

	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var jsonData []map[string]interface{}
	json.Unmarshal(data, &jsonData)

	for i, v := range jsonData {
		for ind, val := range arrmap {
			if v["window"] == val["window"] && v["commit"] == val["commit"] && v["repository"] == val["repository"] {
				str1, _ := strconv.Atoi(jsonData[i]["secend"].(string))
				str2, _ := strconv.Atoi(arrmap[ind]["secend"].(string))
				jsonData[i]["secend"] = strconv.Itoa(str1 + str2)
				arrmap = append(arrmap[:ind], arrmap[ind+1:]...)

			}
		}
	}

	jsonData = append(jsonData, arrmap...)

	modifiedData, _ := json.MarshalIndent(jsonData, "", "  ")

	ioutil.WriteFile("data.json", modifiedData, 0644)

}

func TimeCa(keypress chan int) {
	time.Sleep(time.Millisecond * 1)
	if CheckTime(keypress) <= 60000 {
		second++
		data := ArrMap{
			window: GetWindowWrite(),
			secend: 1,
		}
		for index, item := range arr {
			if item.window == data.window {
				arr[index].secend = item.secend + 1
				return
			}
		}
		arr = append(arr, data)
		HandelFile()
		// Activity <- arr
	}
}
