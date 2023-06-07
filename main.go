package main

import (
	"encoding/json"
	"fmt"

	"github.com/praise579/project-manager/config"
)

func main() {
	byte, _ := json.Marshal(config.Conf)
	fmt.Println(string(byte))
}
