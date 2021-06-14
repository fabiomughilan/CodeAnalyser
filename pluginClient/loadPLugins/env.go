package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type EnvPlugin struct {
	Methods *interfaces.Env
	Client  *plugin.Client
}

func (receiver *EnvPlugin) Load(yamlFile *pbUtils.Details) {
	methods, client := pluginClient.CreateEnvClient(utils.CallPluginCommand(yamlFile.Command))
	receiver.Client = client
	receiver.Methods = &methods
}
