package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type BuildDirectoryPlugin struct {
	Methods *interfaces.BuildDirectory
	Client  *plugin.Client
}

func (receiver *BuildDirectoryPlugin) Load(pluginPath string) {
	methods, client := pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(pluginPath))
	receiver.Client = client
	receiver.Methods = &methods
}
