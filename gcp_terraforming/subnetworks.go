// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.
package gcp_terraforming

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var subnetworksAllowEmptyValues = []string{""}

var subnetworksAdditionalFields = map[string]string{}

type SubnetworksGenerator struct {
	GCPService
}

// Run on subnetworksList and create for each TerraformResource
func (g SubnetworksGenerator) createResources(subnetworksList *compute.SubnetworksListCall, ctx context.Context) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := subnetworksList.Pages(ctx, func(page *compute.SubnetworkList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_subnetwork",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
				},
				subnetworksAllowEmptyValues,
				subnetworksAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each subnetworks create 1 TerraformResource
// Need subnetworks name as ID for terraform resource
func (g *SubnetworksGenerator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	subnetworksList := computeService.Subnetworks.List(g.GetArgs()["project"], g.GetArgs()["region"])

	g.Resources = g.createResources(subnetworksList, ctx)
	g.PopulateIgnoreKeys()
	return nil

}
