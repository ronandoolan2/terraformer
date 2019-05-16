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
package gcp

import (
	"context"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var imagesAllowEmptyValues = []string{""}

var imagesAdditionalFields = map[string]string{}

type ImagesGenerator struct {
	GCPService
}

// Run on imagesList and create for each TerraformResource
func (g ImagesGenerator) createResources(ctx context.Context, imagesList *compute.ImagesListCall) []terraform_utils.Resource {
	resources := []terraform_utils.Resource{}
	if err := imagesList.Pages(ctx, func(page *compute.ImageList) error {
		for _, obj := range page.Items {
			resources = append(resources, terraform_utils.NewResource(
				obj.Name,
				obj.Name,
				"google_compute_image",
				"google",
				map[string]string{
					"name":    obj.Name,
					"project": g.GetArgs()["project"],
					"region":  g.GetArgs()["region"],
				},
				imagesAllowEmptyValues,
				imagesAdditionalFields,
			))
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return resources
}

// Generate TerraformResources from GCP API,
// from each images create 1 TerraformResource
// Need images name as ID for terraform resource
func (g *ImagesGenerator) InitResources() error {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.ComputeReadonlyScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	imagesList := computeService.Images.List(g.GetArgs()["project"])

	g.Resources = g.createResources(ctx, imagesList)
	g.PopulateIgnoreKeys()
	return nil

}