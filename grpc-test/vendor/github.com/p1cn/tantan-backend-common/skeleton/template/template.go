package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"text/template"
)

func GenerateTemplates(rootPath string, param Model) error {
	//rootPath = "."

	appDestPrefix := fmt.Sprintf("%s/%s/%s/%s", rootPath, param.RepoName, "app", param.AppName)
	cmdDestPrefix := fmt.Sprintf("%s/%s/%s/%s", rootPath, param.RepoName, "cmd", param.AppName)
	rootDestPrefix := fmt.Sprintf("%s/%s", rootPath, param.RepoName)
	// tarBuffer, err := os.Create(fmt.Sprintf("%s/%s.gz", rootPath, param.RepoName))
	// defer tarBuffer.Close()

	// gw := gzip.NewWriter(tarBuffer)

	// defer gw.Close()

	// tarWriter := tar.NewWriter(gw)
	// defer tarWriter.Close()

	appTplPath := rootPath + "/tpl/app"
	cmdTplPath := rootPath + "/tpl/cmd"
	rootTplPath := rootPath + "/tpl"

	// makefile
	err := generateTemplate(rootTplPath+"/Makefile.tpl", rootDestPrefix+"/Makefile", &param)
	if err != nil {
		log.Fatal(err)
	}

	// config
	err = generateTemplate(appTplPath+"/config/config.tpl", appDestPrefix+"/config/config.go", &param)
	if err != nil {
		log.Fatal(err)
	}

	// events
	if len(param.Events) > 0 {
		// event
		err = generateTemplate(appTplPath+"/event/event.tpl", appDestPrefix+"/event/event_handler.go", &param)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/event/event_test.tpl", appDestPrefix+"/event/event_handler_test.go", &param)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/event/event_mock.tpl", appDestPrefix+"/event/event_handler_mock.go", &param)
		if err != nil {
			log.Fatal(err)
		}

		// event handler
		err = generateTemplate(appTplPath+"/handler/event_handler.tpl", appDestPrefix+"/handler/event_handler.go", &param)
		if err != nil {
			log.Fatal(err)
		}
		err = generateTemplate(appTplPath+"/handler/event_handler_mock.tpl", appDestPrefix+"/handler/event_handler_mock.go", &param)
		if err != nil {
			log.Fatal(err)
		}
		err = generateTemplate(appTplPath+"/handler/event_handler_test.tpl", appDestPrefix+"/handler/event_handler_test.go", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	// model
	for _, model := range param.Models {
		tmp := param
		tmp.Data = model
		err = generateTemplate(appTplPath+"/model/model.tpl", appDestPrefix+"/model/"+model.PackageName+"/model.go", &tmp)
		if err != nil {

			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/model/model_test.tpl", appDestPrefix+"/model/"+model.PackageName+"/model_test.go", &tmp)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/model/model_mock.tpl", appDestPrefix+"/model/"+model.PackageName+"/model_mock.go", &tmp)
		if err != nil {
			log.Fatal(err)
		}

		if param.DclCommiter {
			err = generateTemplate(appTplPath+"/model/dcl_commiter.tpl", appDestPrefix+"/model/"+model.PackageName+"/dcl_commiter.go", &tmp)
			if err != nil {
				log.Fatal(err)
			}

			err = generateTemplate(appTplPath+"/model/dcl_commiter_test.tpl", appDestPrefix+"/model/"+model.PackageName+"/dcl_commiter_test.go", &tmp)
			if err != nil {
				log.Fatal(err)
			}

			err = generateTemplate(appTplPath+"/model/dcl_commiter_mock.tpl", appDestPrefix+"/model/"+model.PackageName+"/dcl_commiter_mock.go", &tmp)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = generateTemplate(appTplPath+"/model/db.tpl", appDestPrefix+"/model/"+model.PackageName+"/db.go", &tmp)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/model/db_test.tpl", appDestPrefix+"/model/"+model.PackageName+"/db_test.go", &tmp)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/model/db_mock.tpl", appDestPrefix+"/model/"+model.PackageName+"/db_mock.go", &tmp)
		if err != nil {
			log.Fatal(err)
		}
	}

	// RPC
	if param.Config.RPC {
		err = generateTemplate(appTplPath+"/rpcserver/server.tpl", appDestPrefix+"/rpcserver/server.go", &param)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/rpcserver/server_test.tpl", appDestPrefix+"/rpcserver/server_test.go", &param)
		if err != nil {
			log.Fatal(err)
		}
		err = generateTemplate(appTplPath+"/rpcserver/server_mock.tpl", appDestPrefix+"/rpcserver/server_mock.go", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	// http
	if param.Config.HTTP {
		err = generateTemplate(appTplPath+"/httpserver/server.tpl", appDestPrefix+"/httpserver/server.go", &param)
		if err != nil {
			log.Fatal(err)
		}

		err = generateTemplate(appTplPath+"/httpserver/server_mock.tpl", appDestPrefix+"/httpserver/server_mock.go", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	if param.Config.Dependencies {
		err = generateTemplate(appTplPath+"/rpcclient/client.tpl", appDestPrefix+"/rpcclient/client.go", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	// service
	err = generateTemplate(appTplPath+"/service/service.tpl", appDestPrefix+"/service/service.go", &param)
	if err != nil {
		log.Fatal(err)
	}
	err = generateTemplate(appTplPath+"/service/service_test.tpl", appDestPrefix+"/service/service_test.go", &param)
	if err != nil {
		log.Fatal(err)
	}

	// configuration files
	err = generateTemplate(appTplPath+"/conf/common.tpl", appDestPrefix+"/conf/common.json", &param)
	if err != nil {
		log.Fatal(err)
	}

	err = generateTemplate(appTplPath+"/conf/service.tpl", appDestPrefix+"/conf/service.json", &param)
	if err != nil {
		log.Fatal(err)
	}

	if param.Config.Cache {
		err = generateTemplate(appTplPath+"/conf/cache.tpl", appDestPrefix+"/conf/cache.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.DB {
		err = generateTemplate(appTplPath+"/conf/db.tpl", appDestPrefix+"/conf/db.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.DCL {
		err = generateTemplate(appTplPath+"/conf/dcl.tpl", appDestPrefix+"/conf/dcl.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.MQ {
		err = generateTemplate(appTplPath+"/conf/mq.tpl", appDestPrefix+"/conf/mq.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.RPC {
		err = generateTemplate(appTplPath+"/conf/rpc.tpl", appDestPrefix+"/conf/rpc.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.HTTP {
		err = generateTemplate(appTplPath+"/conf/http.tpl", appDestPrefix+"/conf/http.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	if param.Config.Dependencies {
		err = generateTemplate(appTplPath+"/conf/dependencies.tpl", appDestPrefix+"/conf/dependencies.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}
	if param.Config.EventLog {
		err = generateTemplate(appTplPath+"/conf/eventlog.tpl", appDestPrefix+"/conf/eventlog.json", &param)
		if err != nil {
			log.Fatal(err)
		}
	}

	// main
	err = generateTemplate(cmdTplPath+"/main.tpl", cmdDestPrefix+"/main.go", &param)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func newTemplate(path string, name string) (*template.Template, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	tpl, err := template.New(name).Parse(string(data))
	if err != nil {
		log.Fatal(err, path, name)
	}
	return tpl, nil
}

func executeTemplate(infile string, fileName string, wr io.Writer, param *Model) error {
	tpl, err := newTemplate(infile, fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(wr, &param)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// func tarFile(fileName string, buffer *bytes.Buffer *tar.Writer) error {
// 	h := new(tar.Header)
// 	h.Name = fileName
// 	h.Size = int64(buffer.Len())
// 	h.Mode = 0666

// 	err := tarWriter.WriteHeader(h)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = io.Copy(tarWriter, buffer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return nil
// }

// func generateTemplate(tplFile, tarFileName string, param *Model *tar.Writer) error {
// 	buffer := bytes.NewBuffer([]byte{})

// 	strs := strings.Split(tarFileName, "/")
// 	var name string
// 	if len(strs) > 0 {
// 		name = strs[len(strs)-1]
// 	}
// 	fmt.Println(name, tarFileName)
// 	err := executeTemplate(tplFile, name, buffer, param)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return tarFile(tarFileName, buffer)
// }

func generateTemplate(tplFile, tarFileName string, param *Model) error {
	buffer := bytes.NewBuffer([]byte{})

	strs := strings.Split(tarFileName, "/")
	var name string
	if len(strs) > 0 {
		name = strs[len(strs)-1]
	}

	err := executeTemplate(tplFile, name, buffer, param)
	if err != nil {
		log.Fatal(err)
	}

	fileDir, _ := filepath.Abs(filepath.Dir(tarFileName))

	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(tarFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(buffer.String())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
