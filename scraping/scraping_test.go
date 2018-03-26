package scraping

import (
	"strings"
	"testing"
)

func TestGetResourceUrl(t *testing.T) {
	result_aws := GetResourceUrl("aws_instance")
	if result_aws != "https://www.terraform.io/docs/providers/aws/r/instance.html" {
		t.Error(result_aws)
	}

	result_azure := GetResourceUrl("azurerm_virtual_machine")
	if result_azure != "https://www.terraform.io/docs/providers/azurerm/r/virtual_machine.html" {
		t.Error(result_azure)
	}

	result_grafana := GetResourceUrl("grafana_alert_notification")
	if result_grafana != "https://www.terraform.io/docs/providers/grafana/r/alert_notification.html" {
		t.Error(result_grafana)
	}
}

func TestScrapingDoc(t *testing.T) {
	result := ScrapingDoc("http://www.terraform.io/docs/providers/aws/r/instance.html")
	if !containsTfResource(result, "instance_initiated_shutdown_behavior") {
		t.Error("Terraform resource args is invalid")
	}

	for i, v := range result.Args {
		if v.Name == "ephemeral_block_device" && strings.Contains(v.Description, "Customize Ephemeral") {
			if v.Field_name == "block_devices" {
				break
			}
		}
		if i == len(result.Args) {
			t.Error("Terraform resource args is invalid")
		}
	}

	result = ScrapingDoc("http://www.terraform.io/docs/providers/aws/r/lambda_function.html")
	if !containsTfResource(result, "function_name") {
		t.Error("Terraform resource args is invalid")
	}

	for i, v := range result.Args {
		if v.Name == "s3_bucket" && strings.Contains(v.Description, "The S3 bucket location") {
			break
		}
		if i == len(result.Args) {
			t.Error("Terraform resource args is invalid")
		}
	}

}

func containsTfResource(tfr *TfResource, arg_key string) bool {
	for _, v := range tfr.Args {
		if v.Name == arg_key {
			return true
		}
	}
	return false
}