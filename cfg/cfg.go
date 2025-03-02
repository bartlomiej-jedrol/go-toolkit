package cfg

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type BloodResults struct {
	AirtableBaseID  string `yaml:"airtable_base_id"`
	AirtableTableID string `yaml:"airtable_table_id"`
}

type Functions struct {
	BloodResults `yaml:"blood_results"`
}

type DBAddress struct {
}

type LocalPaths struct {
	SecondBrainPath  string `yaml:"second_brain_path"`
	UploadPath       string `yaml:"upload_path"`
	GoogleDriveCreds string `yaml:"google_drive_creds"`
}

type GoogleDriveFolders struct {
	SecondBrainBackups string `yaml:"second_brain_backups"`
}

type Service struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	// Common fields for all services

	// Optional fields
	LocalPaths         `yaml:"local_paths"`
	GoogleDriveFolders `yaml:"google_drive_folders"`
	DBAddress          `yaml:"db_address"`
	Functions          `yaml:"functions"`
	S3Bucket           string `yaml:"s3_bucket"`
	LambdaTmpPath      string `yaml:"lambda_tmp_path"`
}

type Config struct {
	Email    string    `yaml:"email"`
	Services []Service `yaml:"services"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Parse(cfgPath string) error {
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Printf("failed to read instance config: %v", err)
		return err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Printf("failed to unmarshal instance config: %v", err)
		return err
	}

	return nil
}
