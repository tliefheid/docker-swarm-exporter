package common

import "github.com/docker/docker/client"

func GetCLI() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli
}
