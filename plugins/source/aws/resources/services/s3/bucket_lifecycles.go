// Code generated by codegen; DO NOT EDIT.

package s3

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BucketLifecycles() *schema.Table {
	return &schema.Table{
		Name:      "aws_s3_bucket_lifecycles",
		Resolver:  fetchS3BucketLifecycles,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "bucket_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "abort_incomplete_multipart_upload",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AbortIncompleteMultipartUpload"),
			},
			{
				Name:     "expiration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Expiration"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "noncurrent_version_expiration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NoncurrentVersionExpiration"),
			},
			{
				Name:     "noncurrent_version_transitions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NoncurrentVersionTransitions"),
			},
			{
				Name:     "prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Prefix"),
			},
			{
				Name:     "transitions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Transitions"),
			},
		},
	}
}
