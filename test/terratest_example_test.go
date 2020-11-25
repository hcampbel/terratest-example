package test

import (
	"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformEC2Creation(t *testing.T) {

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "/home/hugh/Documents/Programming/Golang/github.com/hcampbel/terratest-example",
	})

	region := "us-west-2"
	tagName := "Name"
	tagValue := "Test Instance"

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "id")
	// deployedRegion := terraform.Output(t, terraformOptions, "region")

	instances, err := aws.GetTagsForEc2InstanceE(t, region, id)

	for k, v := range instances {

		if err != nil {
			log.Panic("Cannot find instance")
		}

		assert.Equal(t, tagName, k)
		assert.Equal(t, tagValue, v)
	}

}
