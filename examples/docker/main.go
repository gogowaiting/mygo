package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		return
	}

	// 使用正确的类型来获取容器列表
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Println("Failed to list containers:", err)
		return
	}

	for _, container := range containers {
		fmt.Printf("\t%s\n", container.ID)
	}
}
