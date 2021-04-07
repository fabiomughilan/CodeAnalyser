package runners

import (
	"code-analyser/helpers"
	"code-analyser/pluginClient"
	"code-analyser/pluginClient/pb"
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/utils"
	"os/exec"
)

func ParseLibraryFromDependencies(dependenciesList map[string]string, langYamlObject *versionsPB.LanguageVersion) map[string]DependencyDetail {
	library := map[string]DependencyDetail{}
	for key, supportedLibrary := range langYamlObject.Libraries {
		if versionUsed, ok := dependenciesList[key]; ok {
			for versionName, v := range supportedLibrary.Version {
				if helpers.SeverValidateFromArray(v.Semver, versionUsed) {
					library[key] = DependencyDetail{
						Version: versionName,
						Command: v.Plugincommand,
					}
				}
			}
		}
	}
	return library
}

func LibraryRunner(libraryList map[string]DependencyDetail, runtimeVersion, root string) []*languageSpecificPB.LibraryOutput {
	var libraryOutputs []*languageSpecificPB.LibraryOutput
	for libraryUsed, libraryDetails := range libraryList {
		isUsed := LibraryDetectorRunner(libraryUsed, libraryDetails, runtimeVersion, root)
		if isUsed != nil {
			libraryOutputs = append(libraryOutputs, isUsed)
		}
	}
	return libraryOutputs
}

//FrameworkDetectorRunner will find and run version detector & returns protos.FrameworkOutput to
func LibraryDetectorRunner(name string, libraryDetails DependencyDetail, runtimeVersion, root string) *languageSpecificPB.LibraryOutput {
	libraryResponse, client := pluginClient.LibraryPluginCall(exec.Command("sh", "-c", libraryDetails.Command))
	defer client.Kill()
	isUsed, err := libraryResponse.IsUsed(&pb.ServiceInput{
		RuntimeVersion: runtimeVersion,
		Root:           root,
	})
	if err != nil {
		utils.Logger(err)
		return nil
	}
	if isUsed.Value == true {
		detection, err := libraryResponse.Detect(&pb.ServiceInput{
			RuntimeVersion: runtimeVersion,
			Root:           root,
		})
		if err != nil {
			utils.Logger(err)
			return nil
		}
		if detection.Value == true {
			percentUsed, err := libraryResponse.PercentOfUsed(&pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           root,
			})
			if err != nil {
				utils.Logger(err)
				return nil
			}
			return &languageSpecificPB.LibraryOutput{
				Type:           languageSpecificPB.LibraryOutput_Type(detection.Type),
				Used:           isUsed.Value,
				Name:           name,
				Version:        libraryDetails.Version,
				PercentageUsed: percentUsed.Value,
			}
		}
	}
	return nil
}
