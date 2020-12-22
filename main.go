package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClient("tcp://192.168.100.33:2376", "v1.12", nil, nil)
	defer cli.Close()
	log(err)
	listImage(cli)
	//id := createContainer(cli)
	//fmt.Printf("%s\n", id)
	//time.Sleep(time.Second * 1)
	//startContainer(id, cli)
	//time.Sleep(time.Second * 40)
	//stopContainer(id, cli)
	//time.Sleep(time.Second * 3)
	//id, err = removeContainer(id, cli)
	//if err == nil {
	//	fmt.Println("删除容器", id, "成功")
	//}
}
// 列出镜像
func listImage(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	log(err)

	for _, image := range images {
		fmt.Println(image)
	}
}
// 创建容器
//func createContainer(cli *client.Client) string {
//	exports := make(nat.PortSet, 10)
//	port, err := nat.NewPort("tcp", "80")
//	log(err)
//	exports[port] = struct{}{}
//	config := &container.Config{Image: "nginx", ExposedPorts: exports}
//
//	portBind := nat.PortBinding{HostPort: "80"}
//	portMap := make(nat.PortMap, 0)
//	tmp := make([]nat.PortBinding, 0, 1)
//	tmp = append(tmp, portBind)
//	portMap[port] = tmp
//	hostConfig := &container.HostConfig{PortBindings: portMap}
//
//	containerName := "hel"
//	body, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, containerName)
//	log(err)
//	fmt.Printf("ID: %s\n", body.ID)
//	return body.ID
//}
// 启动
func startContainer(containerID string, cli *client.Client) {
	err := cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	log(err)
	if err == nil {
		fmt.Println("容器", containerID, "启动成功")
	}
}
// 停止
func stopContainer(containerID string, cli *client.Client) {
	timeout := time.Second * 10
	err := cli.ContainerStop(context.Background(), containerID, &timeout)
	if err != nil {
		log(err)
	} else {
		fmt.Printf("容器%s已经被停止\n", containerID)
	}
}
// 删除
func removeContainer(containerID string, cli *client.Client) (string, error) {
	err := cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{})
	log(err)
	return containerID, err
}

func log(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
}