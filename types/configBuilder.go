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

func (b *configBuilder) BuildAuthConfig() (error, interfaces.AuthMethod) {
	switch b.service.Service.AuthMethod.Type {
	case globalConfig.AUTH_METHOD_BASIC:
		errorGettingBasicAuth, basicAuthConfig := b.buildBasicAuth()
		if errorGettingBasicAuth != nil {
			return errorGettingBasicAuth, nil
		}
		return nil, basicAuthConfig
	default:
		return NewCustomError("Invalid auth method"), nil
	}
}

func (b *configBuilder) GetConfigs() *settings.Config {
	return b.service
}

func (b *configBuilder) buildBasicAuth() (error, *basic.Config) {
	var basicAuthConfig basic.Config
	var filepath string
	if globalConfig.AUTH_METHODS_PATH[len(globalConfig.AUTH_METHODS_PATH)-1:] != "/" {
		filepath = globalConfig.AUTH_METHODS_PATH + "/" + globalConfig.AUTH_METHOD_BASIC + ".xml"
	} else {
		filepath = globalConfig.AUTH_METHODS_PATH + globalConfig.AUTH_METHOD_BASIC + ".xml"
	}

	errorGettingXml := util.ReadXmlFile[basic.Config](filepath, &basicAuthConfig)
	if errorGettingXml != nil {
		util.LoggerHandler().Error("Error getting xml file", "error", errorGettingXml.Error())
		return errorGettingXml, nil
	}
	errorValidatingBasicAuthConfig := util.GetValidator().Struct(basicAuthConfig)
	if errorValidatingBasicAuthConfig != nil {
		util.LoggerHandler().Error("Error validating basic auth config", "error", errorValidatingBasicAuthConfig.Error())
		return errorGettingXml, nil
	}

	return nil, &basicAuthConfig
}
