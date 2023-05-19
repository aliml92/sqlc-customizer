package config


type Config struct {
	Modify	Modify	`mapstructure:"modify"`
}

type Modify struct {
	Models	[]Model	`mapstructure:"model"`
}

type Model struct {
	Name        	string	`mapstructure:"name"`
	Source 			string	`mapstructure:"source"`
	Destination 	string	`mapstructure:"destination"`
	Package 		string	`mapstructure:"package"`
	OldPackage		string	`mapstructure:"old_package"`
	PackagePath		string	`mapstructure:"package_path"`
	JSONOmitempty	bool	`mapstructure:"json_omitempty"`
}

