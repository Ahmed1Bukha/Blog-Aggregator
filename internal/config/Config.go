package config

import (
	"encoding/json"
	"fmt"
	"os"
)
const configFileName = ".gatorconfig.json"

type Config struct{
	DBURL string `json:"db_url"`
	CurrentUserName string`json:"current_user_name"`
}

func Read() (*Config,error){
	var config *Config

	direct,err:= getConfigFilePath()
	if err !=nil{
		return config,fmt.Errorf("error with getConfig func")
	}
	fullPath:= direct+"/"+configFileName
	
	file,err := os.Open(fullPath)

	if err !=nil { 
		return config,fmt.Errorf("error reading the file")
	}
	defer file.Close()

	decoder:= json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err !=nil {
		return config,fmt.Errorf("error decoding the file")
	}
	return config,nil

}

func (c *Config) SetUser(username string) error{
	c.CurrentUserName = username
	err := write(c)
	if err !=nil {
		return fmt.Errorf("error with write function")
	}
	return nil
}

func write (cfg *Config)(error){
	filepath,err:= getConfigFilePath()
	if err !=nil {
		return fmt.Errorf("error with getting local path")
	} 

	jsonData,err := json.Marshal(cfg)
	if err !=nil {
		return fmt.Errorf("error marshaling json data")
	}
	err = os.WriteFile(filepath+"/"+configFileName,jsonData,0644)
	if err !=nil {
		return fmt.Errorf("error with writing file")
	}
	err = os.WriteFile("/Users/bukha/.gatorconfig.json",jsonData,0644)
	if err !=nil {
		return fmt.Errorf("error with writing file")
	}

	return nil
}


func getConfigFilePath() (string, error){
	direct,err := os.Getwd()
	if err != nil {
		return "",fmt.Errorf("error with getting the directory") 
	}
	return direct,nil
}
