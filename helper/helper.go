package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/astaxie/beego"
	"net/http"
	"io/ioutil"
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

func MapArrayToMapString(vals map[string][]string) map[string]string {
	ret := make(map[string]string)
	for key, val := range vals {
		ret[key] = val[0]
	}
	return ret
}

func String(name string) string {
	val, _ := beego.GetConfig("string", name)
	return val.(string)
}

func GooglePlacesAutocompleteApi(requestParamMap map[string]string) map[string]interface{} {
	requestParamString := ""
	for key, val := range requestParamMap {
		requestParamString += key+"="+val+"&"
	}
	requestParamString += "key="+GOOGLE_PLACES_API_KEY
	url := "https://maps.googleapis.com/maps/api/place/autocomplete/json?" + requestParamString
	return JsonUrlToMap(url)
}

func GetJson(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		beego.Error("error while fetching places ", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	return bytes
}

func JsonUrlToMap(url string) map[string]interface{} {
	dec := make(map[string]interface{})
	bytes := GetJson(url)
	//	beego.Info("data ", string(bytes))
	json.Unmarshal(bytes, &dec)
	return dec
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


