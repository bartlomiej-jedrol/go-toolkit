package cfg

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
	DBAddress          `yaml:"db_address`
}

type Config struct {
	Email    string    `yaml:"email"`
	Services []Service `yaml:"services"`
}
