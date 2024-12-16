package terminalpage

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func readFile() {
	file, _ := os.Open("tasks.json")
	defer file.Close()
	var jsonData []map[string]interface{}
	data, _ := io.ReadAll(file)

	json.Unmarshal(data, &jsonData)
	tasks = []Task{}
	for _, value := range jsonData {
		task := Task{}
		task.Ac = value["ac"].(string)
		task.Deedline = value["deadline"].(string)
		task.Dic = value["dic"].(string)
		task.Title = value["title"].(string)
		task.Status = int(value["status"].(float64))
		tasks = append(tasks, task)
	}
}

func addFile() {
	file, _ := os.Create("tasks.json")
	defer file.Close()

	data := []map[string]interface{}{}
	for _, value := range tasks {
		task := map[string]interface{}{}
		task["ac"] = value.Ac
		task["deadline"] = value.Deedline
		task["dic"] = value.Dic
		task["status"] = value.Status
		task["title"] = value.Title
		data = append(data, task)
	}
	modifiedData, _ := json.MarshalIndent(data, "", "  ")

	ioutil.WriteFile("tasks.json", modifiedData, 0644)

}
