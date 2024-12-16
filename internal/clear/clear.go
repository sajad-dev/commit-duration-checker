package clear

import "os"

func Clear() {
	os.Chdir("static")

	os.Create("data.json")
}
