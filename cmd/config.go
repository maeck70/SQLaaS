package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	configPaths = []string{
		"/var/sqlaas",
		"/etc",
		".",
		"..",
		"../configs",
		"../../configs",
	}
	settings *DaasConfigFile
)

func loadConfig(filename string) {
	settings = NewConfig(filename)
}

func NewConfig(filename string) *DaasConfigFile {
	b := DaasConfigFile{}
	b.configReader("")
	t := DaasConfigFile{}
	filepath := fileFinder(filename, configPaths)
	t.configReader(filepath)
	t.fillInBlanks(&b)
	return &t
}

// Load and Marshal yaml file into Config struct
func (t *DaasConfigFile) configReader(file string) {
	var err error
	var yamlFile []byte

	if file != "" {
		yamlFile, err = os.ReadFile(file)
		FailOnError(err, "Unable to read config yaml file "+file)
	} else {
		yamlFile = []byte(configfile_base)
	}

	err = yaml.Unmarshal(yamlFile, &t)
	FailOnError(err, "Unable to unmarshal config yaml file "+file)
}

func (t *DaasConfigFile) fillInBlanks(b *DaasConfigFile) {
	// Database Defaults
	if t.Db.DbPort == 0 {
		t.Db.DbPort = b.Db.DbPort
	}
	if t.Db.DbParams == "" {
		t.Db.DbParams = b.Db.DbParams
	}
	if t.Db.ServicePort == 0 {
		t.Db.ServicePort = b.Db.ServicePort
	}

	// Service Defaults
	for sk := range t.Services {
		service, exists := t.Services[sk]
		if !exists {
			service = DaasService{}
		}
		if service.Name == "" {
			service.Name = sk
		}
		if service.Type == "" {
			service.Type = "GET"
		}
		// Service Fields Defaults
		for fk := range service.Fields {
			field, exists := service.Fields[fk]
			if !exists {
				field = SqlField{}
			}
			if field.Name == "" {
				field.Name = fk
			}
			if field.Type == "" {
				field.Type = "string"
			}
			// Update the field back to service.Fields
			service.Fields[fk] = field
		}
		// Service Arguments Defaults
		for pk := range service.Arguments {
			parameter, exists := service.Arguments[pk]
			if !exists {
				parameter = SqlField{}
			}
			if parameter.Name == "" {
				parameter.Name = pk
			}
			if parameter.Type == "" {
				parameter.Type = "string"
			}
			// Update the parameter back to service.Arguments
			service.Arguments[pk] = parameter
		}
		// Update the service back to t.Services
		t.Services[sk] = service
	}
}

// Scan the paths provided for the file provided, return first found
func fileFinder(file string, paths []string) string {
	// Note, cannot use twistygo.Log as this has not been initialized yet
	var foundFile string
	for _, path := range paths {
		foundFile = fmt.Sprintf("%s/%s", path, file)
		if ok, _ := os.Stat(foundFile); ok != nil {
			fmt.Println("Loading config file " + foundFile)
			return foundFile
		}
	}
	fmt.Printf("File %s does not exist.", file)
	return foundFile
}
