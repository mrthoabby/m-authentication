package types

import (
	"com.github/mrthoabby/m-authentication/globalConfig"
	"com.github/mrthoabby/m-authentication/interfaces"
	"com.github/mrthoabby/m-authentication/types/basic"
	"com.github/mrthoabby/m-authentication/types/settings"
	"com.github/mrthoabby/m-authentication/util"
)

type configBuilder struct {
	service *settings.Config
}

func NewConfigBuilder() *configBuilder {
	var serviceConfig settings.Config
	errorGettingXml := util.ReadXmlFile[settings.Config](globalConfig.CONFIG_SERVICE_PATH, &serviceConfig)
	if errorGettingXml != nil {
		util.LoggerHandler().Error("Error getting xml file", "error", errorGettingXml.Error())
		return nil
	}

	errorValidatingServiceConfig := util.GetValidator().Struct(serviceConfig)
	if errorValidatingServiceConfig != nil {
		util.LoggerHandler().Error("Error validating service config", "error", errorValidatingServiceConfig.Error())
		return nil
	}

	return &configBuilder{
		service: &serviceConfig,
	}
}

func (b *configBuilder) BuildAuthConfig() interfaces.AuthMethod {
	switch b.service.Service.AuthMethod.Type {
	case globalConfig.AUTH_METHOD_BASIC:
		return b.buildBasicAuth()
	default:
		return nil
	}
}

func (b *configBuilder) GetConfigs() *settings.Config {
	return b.service
}

func (b *configBuilder) buildBasicAuth() *basic.Config {
	var basicAuthConfig basic.Config
	filepath := globalConfig.AUTH_METHOD_BASIC + globalConfig.AUTH_METHOD_BASIC + ".xml"

	errorGettingXml := util.ReadXmlFile[basic.Config](filepath, &basicAuthConfig)
	if errorGettingXml != nil {
		util.LoggerHandler().Error("Error getting xml file", "error", errorGettingXml.Error())
		return nil
	}
	errorValidatingBasicAuthConfig := util.GetValidator().Struct(basicAuthConfig)
	if errorValidatingBasicAuthConfig != nil {
		util.LoggerHandler().Error("Error validating basic auth config", "error", errorValidatingBasicAuthConfig.Error())
		return nil
	}

	return &basicAuthConfig
}
