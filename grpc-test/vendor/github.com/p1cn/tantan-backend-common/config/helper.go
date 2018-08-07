package config 
import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"github.com/xeipuuv/gojsonschema"
)

const (
	CommonConfig = iota
	DatabaseConfig
	RpcConfig
	HttpConfig
	MqConfig
	DclConfig
	CacheConfig
	DependenciesConfig
	EventLogConfig
	ServiceConfig
)

var (
	configFiles = map[int]string{
		CommonConfig:       "common.json",
		DatabaseConfig:     "db.json",
		RpcConfig:          "rpc.json",
		HttpConfig:         "http.json",
		MqConfig:           "mq.json",
		DclConfig:          "dcl.json",
		CacheConfig:        "cache.json",
		DependenciesConfig: "dependencies.json",
		EventLogConfig:     "eventlog.json",
		ServiceConfig:      "service.json",
	}
)

func ParseConfig(rootPath string, cfgs map[int]interface{}) error {

	for k, v := range cfgs {
		fp := filepath.Join(rootPath, string(filepath.Separator), configFiles[k])
		err := UnmarshalJsonConfig(fp, v)

		if err != nil {
			err := errors.New(fmt.Sprintf("%s : %v", fp, err))
			return err
		}
	}

	return nil
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ValidateConfig(rootPath string, cfgs map[int]interface{}) error {

	for k, _ := range cfgs {
		validator := filepath.Join(rootPath, string(filepath.Separator), "validate_"+configFiles[k])
		fp := filepath.Join(rootPath, string(filepath.Separator), configFiles[k])

		validator,_ = filepath.Abs(validator)
		fp,_ = filepath.Abs(fp)


        if exist,_:=PathExists(validator); !exist {
			 panic(fmt.Sprintf("config file[%v] not found", fp))
		}

        if exist,_:=PathExists(fp); !exist {
			 panic(fmt.Sprintf("config file[%v] not found", fp))
		}

		schemaLoader := gojsonschema.NewReferenceLoader("file://"+validator)
		documentLoader := gojsonschema.NewReferenceLoader("file://"+fp)

		result ,err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			 panic(fmt.Sprintf("config file[%v] found error[%s], detail:[%v]", fp, err,result))
		}
        if !result.Valid() {
            var totalErr string = " jsonschema found err:\n"
			for _, err := range result.Errors() {
				totalErr += fmt.Sprintf("- %s\n",  err)
			}
            panic(totalErr)
		}
	}

	return nil
}
