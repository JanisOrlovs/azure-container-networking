// Copyright 2017 Microsoft. All rights reserved.
// MIT License

package networkcontainers

import (
	"errors"
	"fmt"
	"net"

	"github.com/Azure/azure-container-networking/cns"
	"github.com/Azure/azure-container-networking/log"
)

const (
	// AzureWebApps type will be used to setup network containers for webapps.
	AzureWebApps = "AzureWebApps"

	// ACI type will be used to setup network containers for ACI.
	ACI = "ACI"
)

// NetworkContainers can be used to perform operations on network containers.
type NetworkContainers struct {
	logpath string
}

func interfaceExists(iFaceName string) (bool, error) {

	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		log.Printf("\n\n*******\nInterface description\n%+v\n", iface)
		log.Printf("%v\n", iface.Flags)
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			log.Printf("%v\n", addr)
		}
		log.Printf("\n\n")

	}
	/*
		ift, _ := interfaceTable(0)
		for _, ifi := range ift {
			log.Printf("\n*******\nInterface description from Table\n%+v\n\n", ifi)
		}
	*/
	_, err := net.InterfaceByName(iFaceName)

	if err != nil {
		errMsg := fmt.Sprintf("[Azure CNS] Unable to get interface by name %v", iFaceName)
		log.Printf(errMsg)
		return false, errors.New(errMsg)
	}

	return true, nil
}

// Create creates a network container.
func (cn *NetworkContainers) Create(createNetworkContainerRequest cns.CreateNetworkContainerRequest) error {
	log.Printf("[Azure CNS] NetworkContainers.Create called")
	err := createOrUpdateInterface(createNetworkContainerRequest)
	log.Printf("[Azure CNS] NetworkContainers.Create finished.")
	return err
}

// Update updates a network container.
func (cn *NetworkContainers) Update(createNetworkContainerRequest cns.CreateNetworkContainerRequest) error {
	log.Printf("[Azure CNS] NetworkContainers.Update called")
	err := createOrUpdateInterface(createNetworkContainerRequest)
	log.Printf("[Azure CNS] NetworkContainers.Update finished.")
	return err
}

// Delete deletes a network container.
func (cn *NetworkContainers) Delete(networkContainerID string) error {
	log.Printf("[Azure CNS] NetworkContainers.Update called")
	err := deleteInterface(networkContainerID)
	log.Printf("[Azure CNS] NetworkContainers.Update finished.")
	return err
}
