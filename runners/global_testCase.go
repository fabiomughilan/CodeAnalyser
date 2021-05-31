package runners

import (
	"code-analyser/protos/pb/output/global"
	versionsPB "code-analyser/protos/pb/versions"
	commonUtils "code-analyser/utils"
	"code-analyser/utils/testing"
	"golang.org/x/net/context"
)

type DockerCase struct {
	Input  GlobalRunnerInput
	Output testing.DockerCaseOutput
}

type GlobalRunnerInput struct {
	Ctx          context.Context
	Path         string
	GlobalPlugin *versionsPB.GlobalPlugin
}

type ProcFileCase struct {
	Input  GlobalRunnerInput
	Output *global.ProcFileOutput
}

type MakeFileCase struct {
	Input  GlobalRunnerInput
	Output *global.MakefileOutput
}

var DockerCases = []DockerCase{
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/detectDockerFile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{
			DockerFile: &global.DockerFileOutput{
				Used:     true,
				FilePath: commonUtils.ProjectPath() + "/testingRepos/detectDockerFile/repo1/Dockerfile",
			},
			DockerComposeFile: &global.DockerComposeFileOutput{
				Used:     true,
				FilePath: commonUtils.ProjectPath() + "/testingRepos/detectDockerFile/repo1/docker-compose.yml",
			},
		},
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/detectDockerFile/repo2",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{
			DockerFile: nil,
			DockerComposeFile: &global.DockerComposeFileOutput{
				Used:     true,
				FilePath: commonUtils.ProjectPath() + "/testingRepos/detectDockerFile/repo2/docker-compose.yaml",
			},
		},
	},
	{
		Input: GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: testing.DockerCaseOutput{},
	},
}

var ProcFileCases = []ProcFileCase{
	{
		Input:  GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: nil,
	},
	{
		Input:  GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/detectProcfile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: &global.ProcFileOutput{
			Used:     true,
			FilePath: commonUtils.ProjectPath() + "/testingRepos/detectProcfile/repo1/Procfile",
			Commands: map[string]*global.Command{
				"web": {
					Command: "bundle",
					Args:    []string{"exec","rackup"},
				},
				"worker": {
					Command: "rake",
					Args:    []string{"resque:work"},
				},
			},
		},
	},
}

var MakeFileCases = []MakeFileCase{
	{
		Input:  GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/emptyRepo",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: nil,
	},
	{
		Input:  GlobalRunnerInput{
			Ctx:          nil,
			Path:         commonUtils.ProjectPath() + "/testingRepos/detectMakefile/repo1",
			GlobalPlugin: &GlobalPluginPath,
		},
		Output: &global.MakefileOutput{
			Used: true,
			FilePath: commonUtils.ProjectPath() + "/testingRepos/detectMakefile/repo1/Makefile",
		},
	},
}