package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Git struct {
}

type GitCoder interface {
	HandelCommand()
	Push(string)
	Commit(string)
}

var _ GitCoder = (*Git)(nil)

func (git *Git) Push(remote string) {
	os.Chdir("static")

	command_get_repository := exec.Command("git", "remote", "get-url", remote)
	url, _ := command_get_repository.Output()
	urlsplit := strings.Split(string(url), "/")

	file, _ := os.Open("data.json")

	defer file.Close()

	filedata, _ := io.ReadAll(file)

	var data []map[string]interface{}
	json.Unmarshal(filedata, &data)

	for index, value := range data {
		valuestr := value["repository"].(string)
		if len(valuestr) == 0 {
			data[index]["repository"] = strings.ReplaceAll(urlsplit[len(urlsplit)-1], "\n", "")
		}
	}
	modifiedData, _ := json.MarshalIndent(data, "", "  ")

	ioutil.WriteFile("data.json", modifiedData, 0644)
}

func (git *Git) Commit(message string) {
	message = strings.ReplaceAll(message, `"`, "")
	file, _ := os.Open("data.json")

	defer file.Close()

	filedata, _ := io.ReadAll(file)

	var data []map[string]interface{}
	json.Unmarshal(filedata, &data)

	for index, value := range data {
		valuestr := value["commit"].(string)
		if len(valuestr) == 0 {
			data[index]["commit"] = message
		}
	}
	modifiedData, _ := json.MarshalIndent(data, "", "  ")

	ioutil.WriteFile("data.json", modifiedData, 0644)
}

func (git *Git) HandelCommand() {
	if len(os.Args) < 3 {
		fmt.Println(fmt.Errorf("Not Found Args %s",os.Args[1]))
		return
	}

	if os.Args[2] == "push" {
		git.Push(os.Args[len(os.Args)-1])

	}

	if os.Args[2] == "commit" {
		git.Commit(os.Args[len(os.Args)-1])
	}

	command_pwd := exec.Command("pwd")
	pwd, _ := command_pwd.Output()

	os.Chdir(string(pwd))

	command_git := exec.Command("git", os.Args[2:]...)
	var output, err bytes.Buffer
	command_git.Stdout = &output
	command_git.Stderr = &err

	command_git.Run()

	if len(err.String()) != 0 {
		fmt.Println(err.String())
	} else {
		fmt.Println(output.String())

	}
}
