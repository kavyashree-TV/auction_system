package infrastructure

import (
	"../utils"
	"github.com/tkanos/gonfig"
)

//ConfigService : Utility layer to incorporate all the helper function
type ConfigService struct{}

var _logger = utils.GetLogger()

//Configuration : Database configuration
type Configuration struct {
	MongoDBHost     string `json:"mongodbhost"`
	MongoDBDatabase string `json:"mongodbdatabase"`
	MongoDBUsername string `json:"mongodbusername"`
	MongoDBPassword string `json:"mongodbpassword"`
}

//GetConfiguration : Returns configuration for the database
func (cs *ConfigService) GetConfiguration() Configuration {
	var configuration = Configuration{}
	var err = gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		_logger.Error("Error access configuration file: " + err.Error())
	}
	return configuration
}
