package main

type DaasConfig struct {
	DbHost      string `yaml:"dbhost"`
	DbPort      int    `yaml:"dbport"`
	DbUser      string `yaml:"dbuser"`
	DbPassword  string `yaml:"dbpassword"`
	DbName      string `yaml:"dbname"`
	DbParams    string `yaml:"dbparams"`
	ServicePort int    `yaml:"serviceport"`
}

type SqlField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type DaasService struct {
	Name      string              `yaml:"name"`
	Url       string              `yaml:"url"`
	Sql       string              `yaml:"sql"`
	Type      string              `yaml:"type"`
	Fields    map[string]SqlField `yaml:"fields"`
	Arguments map[string]SqlField `yaml:"arguments"`
}

type DaasConfigFile struct {
	Db       DaasConfig             `yaml:"Db"`
	Services map[string]DaasService `yaml:"Services"`
}
