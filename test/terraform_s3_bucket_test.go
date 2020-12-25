package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformS3Example(t *testing.T) {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../s3_example",
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
    test_bucket := terraform.Output(t, terraformOptions, "test_bucket")
	assert.Equal(t, "terraform-santiago-terratest", test_bucket)


	test1 := terraform.Output(t, terraformOptions, "test1")
	assert.Equal(t, "test1.txt", test1)

	test2 := terraform.Output(t, terraformOptions, "test2")
	assert.Equal(t, "test2.txt", test2)

}


