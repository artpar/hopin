package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/astaxie/beego"
)

type Configuration struct {
	LogFileLocation string
	Database        ConnectionParam
	Server          Server
}

type ConnectionParam struct {
	User     string
	Password string
	Port     int
	Host     string
	Database string
}

type Server struct {
	Host string
	Port int
}

func GetConfiguration(env string) Configuration {
	file, _ := os.Open("config_" + env + ".json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

var argumentMap map[uint8]string = map[uint8]string {
	'P': "profile",
}

func String(name string) string {
	val, _ := beego.GetConfig("string", name)
	return val.(string)
}

func ProcessArguments(args []string) map[string]string {
	var argMap map[string]string = make(map[string]string)
	for _, argString := range args {
		if strings.HasPrefix(argString, "--") {
			i := strings.Index(argString, "=")
			if i != -1 {
				name := argString[2:i]
				value := argString[i+1:]
				argMap[name] = value
			}
		} else if (strings.HasPrefix(argString, "-")) {
			argMap[argumentMap[argString[1]]] = argString[2:]
		}
	}
	return argMap
}
