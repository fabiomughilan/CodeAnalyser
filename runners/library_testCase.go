package runners

import (
	languageSpecificPB "code-analyser/protos/pb/output/languageSpecific"
	versionsPB "code-analyser/protos/pb/versions"
)

type LibrariesParsingCase struct {
	Input  LibrariesParsingInput
	Output map[string]DependencyDetail
}

type LibrariesParsingInput struct {
	DependenciesList map[string]string
	LangYamlObject   *versionsPB.LanguageVersion
}

var LibrariesParsingCases = []LibrariesParsingCase{
	{
		Input: LibrariesParsingInput{
			DependenciesList: map[string]string{},
			LangYamlObject:   &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: LibrariesParsingInput{
			DependenciesList: map[string]string{
				"sdgdfg": "sfd454",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{},
	},
	{
		Input: LibrariesParsingInput{
			DependenciesList: map[string]string{
				"kafka-node": "5.6",
				"react":      "17.5",
				"mongoose":   "5.8",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"kafka-node": {
				Version: "v1.x",
				Command: "node plugin/js/libraries/kafka/v1_x/server.js",
			},
			"react": {
				Version: "v1.x",
				Command: "node plugin/js/libraries/react/v1_x/server.js",
			},
			"mongoose": {
				Version: "v1.x",
				Command: "node plugin/js/libraries/mongoose/v1_x/server.js",
			},
		},
	},
	{
		Input: LibrariesParsingInput{
			DependenciesList: map[string]string{
				"knex": "0.6",
			},
			LangYamlObject: &SupportedDependencies,
		},
		Output: map[string]DependencyDetail{
			"knex": {
				Version: "v1.x",
				Command: "node plugin/js/libraries/knex/v1_x/server.js",
			},
		},
	},
}

type LibraryRunnerCase struct {
	Input  LibraryRunnerInput
	Output []*languageSpecificPB.LibraryOutput
}

type LibraryRunnerInput struct {
	LibraryList    map[string]DependencyDetail
	RuntimeVersion string
	Root           string
}

var LibraryRunnerCases = []LibraryRunnerCase{
	{
		Input:  LibraryRunnerInput{},
		Output: nil,
	},
}
