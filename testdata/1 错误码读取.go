package main

import (
	"encoding/json"
	"fmt"
	"gvb_server/models/res"
	"os"

	"github.com/sirupsen/logrus"
)

const file = "models/res/err_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	bytedata, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap = ErrMap{}
	err = json.Unmarshal(bytedata, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[res.SettingsError])
}
