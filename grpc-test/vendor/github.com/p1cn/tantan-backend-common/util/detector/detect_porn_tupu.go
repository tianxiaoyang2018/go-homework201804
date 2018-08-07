package detector

import (
	"fmt"

	rcn "github.com/tuputech/tupu-go-sdk/recognition"

	slog "github.com/p1cn/tantan-backend-common/log"
)

const (
	PORN_TASK_ID = "54bcfc6c329af61034f7c2fc"
	AD_TASK_ID   = "5588dba4c7ee53a04b5fad7d"

	PORN_LABEL   = 0
	SEXY_LABEL   = 1
	NORMAL_LABEL = 2
)

type TupuClient struct {
	Handler  *rcn.Handler
	SecretId string
}

var label = map[string]string{
	"0": "porn",
	"1": "sexy",
	"2": "normal",
}

func NewTupuClient(privateKeyFile string, secretId string) *TupuClient {

	handler, e := rcn.NewHandler(privateKeyFile)
	if e != nil {
		slog.Err("Failed: %v\n", e)
	}

	return &TupuClient{
		Handler:  handler,
		SecretId: secretId,
	}
}

func (self *TupuClient) Detect(images1 []string) (resultMap map[string]int) {
	return formatResult(self.Handler.PerformWithURL(self.SecretId, images1, nil))
}

func (self *TupuClient) IsPornMoment(images1 []string) bool {
	//slog.Debug("%v", images1)
	return self.isPornUrls(formatResult(self.Handler.PerformWithURL(self.SecretId, images1, nil)))
}

func (self *TupuClient) isPornUrls(resultMap map[string]int) bool {

	if nil == resultMap {
		return false
	}

	for _, label := range resultMap {
		if label == PORN_LABEL {
			return true
		}
	}

	return false
}

func formatResult(result string, statusCode int, e error) map[string]int {
	if e != nil {
		slog.Err("Failed: %v\n", e)
		return make(map[string]int)
	}

	resultMap := map[string]int{}

	//slog.Debug("%v", result)

	r := rcn.ParseResult(result)

	if r != nil {

		v, exist := r.Tasks[PORN_TASK_ID]

		if exist {
			tmp := v.(map[string]interface{})
			files := tmp["fileList"].([]interface{})
			for _, file := range files {
				field := file.(map[string]interface{})
				resultMap[fmt.Sprintf("%v", field["name"])] = int(field["label"].(float64))
			}
		}
	}

	//slog.Debug("%v", resultMap)

	return resultMap
}

func printResult(result string, statusCode int, e error) {
	if e != nil {
		fmt.Printf("Failed: %v\n", e)
		return
	}

	r := rcn.ParseResult(result)
	if r != nil {
		v, exist := r.Tasks["54bcfc6c329af61034f7c2fc"]

		if exist {
			tmp := v.(map[string]interface{})
			files := tmp["fileList"].([]interface{})
			for _, file := range files {
				field := file.(map[string]interface{})
				//data, _ := json.Marshal(field)
				fmt.Printf("%v,%v\n", field["name"], label[fmt.Sprintf("%v", field["label"])])
			}
		}
	}
}
