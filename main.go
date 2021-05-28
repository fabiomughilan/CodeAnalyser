package main

import (
	"code-analyser/analyser"
	"code-analyser/pluginClient"
	decisionmakerPB "code-analyser/protos/pb"
	pb "code-analyser/protos/pb/plugin"
	versionsPB "code-analyser/protos/pb/versions"
	"code-analyser/runners"
	"log"
	"math"
	"os/exec"
	"sync"

	"golang.org/x/net/context"
)

func main() {
	//path := os.Args[1]
	//log.Println("Initialized Scrapping ")
	//log.Println("Scrapping on "+path)
	//decisionMakerInput := Scrape("/home/deqode/Downloads/testingfolder")
	//log.Println(decisionMakerInput.LanguageSpecificDetection)
	//res, err := json.MarshalIndent(decisionMakerInput.LanguageSpecificDetection[0], "", "  ")
	//if err != nil {
	//	log.Println("error:", err)
	//}
	//fmt.Print(string(res))
	//
	res, client := pluginClient.ProcFilePluginCall(exec.Command("sh", "-c", "go run /home/deqode/Desktop/project/go/code-analyser/plugin/globalDetectors/procFile/main.go"))
	result, _ := res.Detect(&pb.ServiceInputString{
		Value: "/home/deqode/Downloads/basicRepo",
	})
	log.Println(result)
	if client.Exited() {
		client.Kill()
	}
}

//Scrape it scrape language, framework, orm etc .....
func Scrape(path string) *decisionmakerPB.DecisionMakerInput {
	languages, _, _ := analyser.Analyse(path)
	supportedLanguages, _ := SupportedLanguagedParser()
	decisionMakerInput := &decisionmakerPB.DecisionMakerInput{
		LanguageSpecificDetection: []*decisionmakerPB.LanguageSpecificDetections{},
		GloabalDetections:         &decisionmakerPB.GlobalDetections{},
	}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	var ctx context.Context = nil
	var mxLang = 0.0
	for _, lang := range languages {
		mxLang = math.Max(mxLang, lang.Percent)
	}
	for _, language := range languages {
		if language.Percent == mxLang {
			wg.Add(1)
			language := language
			go func() {
				defer wg.Done()
				var languagePath string
				for _, supportedLanguage := range supportedLanguages.Languages {
					if supportedLanguage.Name == language.Name {
						languagePath = supportedLanguage.Path
						break
					}
				}
				if languagePath != "" {
					pluginDetails := ParsePluginYamlFile(languagePath)
					globalPlugins := ParseGlobalPluginYaml("./plugin/globalDetectors")
					runtimeVersion := runners.DetectRuntime(ctx, path, pluginDetails)
					runners.RunPreDetectCommand(ctx, &pb.ServiceInput{
						RuntimeVersion: runtimeVersion,
						Root:           path,
					}, pluginDetails)
					if runtimeVersion != "" {
						allDependencies := runners.GetParsedDependencies(ctx, runtimeVersion, path, pluginDetails)
						languageSpecificDetections := decisionmakerPB.LanguageSpecificDetections{
							Name:           language.Name,
							RuntimeVersion: runtimeVersion,
						}
						gloabalDetections := decisionmakerPB.GlobalDetections{}
						commands := decisionmakerPB.Commands{}
						RunAllDetectors(ctx, &languageSpecificDetections, allDependencies, pluginDetails, runtimeVersion, path, &gloabalDetections, globalPlugins, &commands)
						mutex.Lock()
						decisionMakerInput.LanguageSpecificDetection = append(decisionMakerInput.LanguageSpecificDetection, &languageSpecificDetections)
						decisionMakerInput.GloabalDetections = &gloabalDetections
						decisionMakerInput.Commands = &commands
						mutex.Unlock()
						//for _, v := range decisionMakerInput.LanguageSpecificDetection {
						//	log.Println("Language identified as ", v.Name, "& version", v.RuntimeVersion)
						//	for _, f := range v.Framework {
						//		log.Println("Framework identified as ", f.Name, f.Version)
						//	}
						//	for _, f := range v.Db.Databases {
						//		log.Println("Database identified as ", f.Name, f.Version, f.Port)
						//	}
						//	log.Println(v.Commands)
						//	log.Println(v.Env)
						//	log.Println(v.TestCases)
						//	log.Println(v.Libraries)
						//	log.Println(v.Orm)
						//}
					}

				}
			}()
		}
	}
	wg.Wait()
	return decisionMakerInput
}

//RunAllDetectors it runs all detectors of dependencies ex. orm,framework etc ....
func RunAllDetectors(ctx context.Context, languageSpecificDetections *decisionmakerPB.LanguageSpecificDetections, allDependencies map[string]map[string]runners.DependencyDetail, pluginDetails *versionsPB.LanguageVersion, runtimeVersion string, path string, globalDetection *decisionmakerPB.GlobalDetections, globalPlugin *versionsPB.GlobalPlugin, commands *decisionmakerPB.Commands) {
	var wg sync.WaitGroup
	wg.Add(12)
	var mutex = &sync.Mutex{}
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.DockerFile, globalDetection.DockerComposeFile = runners.DetectDockerAndComposeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.Commands != "" {
			mutex.Lock()
			languageSpecificDetections.Commands = runners.GetCommands(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Orm = runners.OrmRunner(allDependencies[runners.ORM], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		commands.SeedCommands, commands.BuildCommands, commands.MigrationCommands, commands.StartUpCommands, commands.AdHocScriptsOutput = runners.DetectAndRunCommands(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.ProcFile = runners.DetectAndRunProcFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		globalDetection.Makefile = runners.DetectAndRunMakeFile(ctx, path, globalPlugin)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Db = runners.DbRunner(allDependencies[runners.DB], runtimeVersion, path)
		mutex.Unlock()

	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Framework = runners.FrameworkRunner(allDependencies[runners.Framework], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.DetectEnvCommand != "" {
			mutex.Lock()
			languageSpecificDetections.Env = runners.EnvDetectAndRunner(pluginDetails, runtimeVersion, path)
			mutex.Unlock()
		}

	}()
	go func() {
		defer wg.Done()
		if pluginDetails.StaticAssetsCommand != "" {
			mutex.Lock()
			languageSpecificDetections.StaticAssets = runners.RunStaticAssetsCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		languageSpecificDetections.Libraries = runners.LibraryRunner(allDependencies[runners.Library], runtimeVersion, path)
		mutex.Unlock()
	}()
	go func() {
		if pluginDetails.BuildDirectoryCommand != "" {
			mutex.Lock()
			languageSpecificDetections.BuildDirectory = runners.DetectAndRunBuildDirectory(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if pluginDetails.DetectTestCasesCommand != "" {
			mutex.Lock()
			languageSpecificDetections.TestCases = runners.DetectTestCasesCommand(ctx, &pb.ServiceInput{
				RuntimeVersion: runtimeVersion,
				Root:           path,
			}, pluginDetails)
			mutex.Unlock()
		}
	}()
	wg.Wait()

}
