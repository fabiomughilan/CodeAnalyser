package main

import (
	"code-analyser/pluginClient"
	dockerFile "code-analyser/pluginClient/docker"
	"code-analyser/protos/pb/output/global"
	pb "code-analyser/protos/pb/plugin"
	"github.com/hashicorp/go-plugin"
	"os"
)

//DockerFile is a plugin
type DockerFile struct {
}

//TODO add file path instead of file name  for docker and docker compose
func (d DockerFile) DetectDockerFiles(path *pb.ServiceInputString) (*pb.ServiceOutputDockerFile, error) {
	dockerFileOutput := &pb.ServiceOutputDockerFile{
		Error: nil,
	}
	if _, err := os.Stat(path.Value + "/Dockerfile"); !os.IsNotExist(err) {
		dockerFileOutput.DockerFile = &global.DockerFileOutput{
			Used:     true,
			FilePath: path.Value + "/Dockerfile",
		}
	}
	return dockerFileOutput, nil
}

func (d DockerFile) DetectDockerComposeFiles(path *pb.ServiceInputString) (*pb.ServiceOutputDockerComposeFile, error) {
	dockerComposeOutput := &pb.ServiceOutputDockerComposeFile{
		Error:             nil,
	}
	if _, err := os.Stat(path.Value + "/docker-compose.yaml"); !os.IsNotExist(err) {
		dockerComposeOutput.DockerComposeFile = &global.DockerComposeFileOutput{
			FilePath: path.Value + "/docker-compose.yaml",
			Used: true,
		}
	}
	if _, err := os.Stat(path.Value + "/docker-compose.yaml"); !os.IsNotExist(err) {
		dockerComposeOutput.DockerComposeFile = &global.DockerComposeFileOutput{
			FilePath: path.Value + "/docker-compose.yaml",
			Used: true,
		}
	}
	return dockerComposeOutput, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginClient.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			pluginClient.PluginDispenserDockerFile: &dockerFile.GRPCPlugin{
				Impl: &DockerFile{},
			}},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
