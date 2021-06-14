package loadPLugins

import (
	"code-analyser/languageDetectors/interfaces"
	"code-analyser/pluginClient"
	pbUtils "code-analyser/protos/pb/output/utils"
	"code-analyser/utils"
	"github.com/hashicorp/go-plugin"
)

type BuildDirectoryPlugin struct {
	Methods interfaces.BuildDirectory
	Client  *plugin.Client
}

func (plugin *BuildDirectoryPlugin) Load(yamlFile *pbUtils.Details) {
	plugin.Methods, plugin.Client = pluginClient.CreateBuildDirectoryClient(utils.CallPluginCommand(yamlFile.Command))
}
