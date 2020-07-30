package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"testing"
)

func vnetClient(subscriptionID string) (*network.VirtualNetworksClient, error) {
	authorizer, err := DeploymentServicePrincipalAuthorizer()
	if err != nil {
		return nil, err
	}

	client := network.NewVirtualNetworksClient(subscriptionID)
	client.Authorizer = authorizer
	return &client, err
}

// VnetSubnetsListE - Return the subnets that exist wihin a given VNET
func VnetSubnetsListE(subscriptionID string, resourceGroupName string, vnetName string) ([]string, error) {

	client, err := vnetClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	vnet, err := client.Get(context.Background(), resourceGroupName, vnetName, "")
	if err != nil {
		return nil, err
	}

	subnets := make([]string, len(*vnet.VirtualNetworkPropertiesFormat.Subnets))
	for index, subnet := range *vnet.VirtualNetworkPropertiesFormat.Subnets {
		subnets[index] = *subnet.ID
	}

	return subnets, nil
}

// VnetSubnetsList - Like VnetSubnetsListE but fails in the case an error is returned
func VnetSubnetsList(t *testing.T, subscriptionID string, resourceGroupName string, vnetName string) []string {
	subnets, err := VnetSubnetsListE(subscriptionID, resourceGroupName, vnetName)
	if err != nil {
		t.Fatal(err)
	}
	return subnets
}

// VnetsListE - Return the vnets that exist within a given resource group
func VnetsListE(subscriptionID string, resourceGroupName string) (*[]string, error) {

	lsList, err := VnetsIDListE(subscriptionID, resourceGroupName)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0)
	for _, linkedService := range *lsList {
		results = append(results, *linkedService.ID)
		if err != nil {
			return nil, err
		}
	}

	return &results, nil
}

// VnetsIDListE - Returns the ID of the vnets in the resource group
func VnetsIDListE(subscriptionID string, resourceGroupName string) (*[]network.VirtualNetwork, error) {
	client, err := vnetClient(subscriptionID)
	if err != nil {
		return nil, err
	}

	iteratorVnet, err := client.ListComplete(context.Background(), resourceGroupName)
	if err != nil {
		return nil, err
	}

	lsList := make([]network.VirtualNetwork, 0)

	for iteratorVnet.NotDone() {
		lsList = append(lsList, iteratorVnet.Value())
		err = iteratorVnet.Next()
		if err != nil {
			return nil, err
		}
	}
	return &lsList, err

}

// VnetsList - Like VnetsListE but fails in the case an error is returned
func VnetsList(t *testing.T, subscriptionID string, resourceGroupName string) *[]string {
	vnets, err := VnetsListE(subscriptionID, resourceGroupName)
	if err != nil {
		t.Fatal(err)
	}
	return vnets
}
