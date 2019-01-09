package osquery_types

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	FirehoseAWSAccessKeyID                   string   `json:"firehose_aws_access_key_id"`
	FirehoseAWSSecretAccessKey               string   `json:"firehose_aws_secret_access_key"`
	FirehoseStreamName                       string   `json:"firehose_stream_name"`
	DistributedQueryLogger                   []string `json:"distributed_query_logger"`
	DistributedQueryLoggerS3BucketName       string   `json:"distributed_query_logger_s3_bucket_name"`
	DistributedQueryLoggerFirehoseStreamName string   `json:"distributed_query_logger_firehose_stream_name"`
	DistributedQueryLoggerFilesytemPath      string   `json:"distributed_query_logger_filesytem_path"`
	AutoApproveNodes                         string   `json:"auto_approve_nodes"`
	Domain                                   string   `json:"domain"`
	Email                                    string   `json:"email"`
	S3ConfigBucket                           string   `json:"s3_config_bucket"`
	UseLEStaging                             string   `json:"use_le_staging"`
}

func GetServerConfig(fn string) (*ServerConfig, error) {

	config := ServerConfig{}
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
