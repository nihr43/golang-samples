package main

import (
	"fmt"
	"github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Instance struct {
	Name  string
	Image string
}

func (i *Instance) create(c incus.InstanceServer) error {
	req := api.InstancesPost{
		Name: i.Name,
		Source: api.InstanceSource{
			Type:     "image",
			Server:   "https://images.linuxcontainers.org",
			Protocol: "simplestreams",
			Alias:    i.Image,
		},
		Type: "container",
	}

	op, err := c.CreateInstance(req)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}
	return nil
}

func listInstances(c incus.InstanceServer) {
	instances, err := c.GetInstanceNames(api.InstanceTypeAny)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range instances {
		fmt.Println(n)
	}
}

func main() {
	c, err := incus.ConnectIncusUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}
	listInstances(c)

	yamlFile, err := os.ReadFile("incus-config.yaml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var instances []Instance

	err = yaml.Unmarshal(yamlFile, &instances)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range instances {
		err = i.create(c)
		if err != nil {
			log.Fatal(err)
		}
	}
}
