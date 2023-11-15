package main

// var configfile = DaasConfigFile{
// 	Db: DaasConfig{
// 		DbHost:      "10.0.0.180",
// 		DbPort:      3306,
// 		DbUser:      "twisty_dex",
// 		DbPassword:  "wonderwall",
// 		DbName:      "twistygo_dex",
// 		DbParams:    "?parseTime=true",
// 		ServicePort: 8080,
// 	},
// 	Services: map[string]DaasService{
// 		"activetoken": {
// 			Name: "activetoken",
// 			Url:  "/activetoken/{token}",
// 			Sql:  "SELECT token, type, last_seen_dttm FROM blockchain_activetokens WHERE token = ?",
// 			Type: "GET",
// 			Fields: map[string]SqlField{
// 				"token":          {Name: "token", Type: "string"},
// 				"type":           {Name: "type", Type: "int"},
// 				"last_seen_dttm": {Name: "last_seen_dttm", Type: "datetime"},
// 			},
// 			Arguments: map[string]SqlField{
// 				"token": {Name: "token", Type: "string"},
// 			},
// 		},
// 		"activetokens": {
// 			Name: "activetokens",
// 			Url:  "/activetokens",
// 			Sql:  "SELECT token, type, last_seen_dttm FROM blockchain_activetokens",
// 			Type: "GET",
// 			Fields: map[string]SqlField{
// 				"token":          {Name: "token", Type: "string"},
// 				"type":           {Name: "type", Type: "int"},
// 				"last_seen_dttm": {Name: "last_seen_dttm", Type: "datetime"},
// 			},
// 		},
// 		"pricingjob": {
// 			Name: "pricingjob",
// 			Url:  "/pricingjob",
// 			Sql:  "select job_id, status, source, route, last_update, next_scheduled, run_interval from blockchain_pricingjob_go",
// 			Type: "GET",
// 			Fields: map[string]SqlField{
// 				"job_id":         {Name: "job_id", Type: "int"},
// 				"status":         {Name: "status", Type: "string"},
// 				"source":         {Name: "source", Type: "string"},
// 				"route":          {Name: "route", Type: "string"},
// 				"last_update":    {Name: "last_update", Type: "datetime"},
// 				"next_scheduled": {Name: "next_scheduled", Type: "datetime"},
// 				"run_interval":   {Name: "run_interval", Type: "int"},
// 			},
// 		},
// 	},
// }

var configfile_base = `
Db:
  dbport:      3306
  dbparams:    "?parseTime=true"
  serviceport: 8080
`
