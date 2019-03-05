package config

// CognisesConfiguration options for cognises
type CognisesConfiguration struct {
	Providers         []string          `mapstructure:"providers,omitempty" yaml:"providers"`
	SSHGenerateConfig SSHGenerateConfig `mapstructure:"ssh_generator_config,omitempty" yaml:"ssh_generator_config"`
}

// SSHGenerateConfig options for ssh-config command
type SSHGenerateConfig struct {
	General      SSHGeneralSettings `mapstructure:"ssh_config_general,omitempty" yaml:"ssh_config_general"`
	PreferPublic bool               `mapstructure:"prefer_public,omitempty" yaml:"prefer_public"`
	StaticHosts  string             `mapstructure:"static_hosts,omitempty" yaml:"static_hosts"`
	PrivateKeys  SSHKeyPairSettings `mapstructure:"private_keys,omitempty" yaml:"private_keys"`
}

// SSHGeneralSettings general settings
type SSHGeneralSettings struct {
	Port         int    `mapstructure:"port,omitempty" yaml:"port"`
	ProxyCommand string `mapstructure:"proxy_command,omitempty" yaml:"proxy_command"`
	User         string `mapstructure:"user,omitempty" yaml:"user"`
}

// SSHKeyPairSettings options for keypairs
type SSHKeyPairSettings struct {
	KeyPath            string `mapstructure:"key_path,omitempty" yaml:"key_path"`
	UseInstanceKeypair bool   `mapstructure:"use_instance_keypair,omitempty" yaml:"use_instance_keypair"`
	CustomKeypair      string `mapstructure:"custom_keypair,omitempty" yaml:"custom_keypair"`
}

// DefaultConfig default cognises configuration
var DefaultConfig = CognisesConfiguration{
	SSHGenerateConfig: SSHGenerateConfig{
		PrivateKeys: SSHKeyPairSettings{
			KeyPath:            "~/.ssh/",
			UseInstanceKeypair: true,
		},
		PreferPublic: false,
		General: SSHGeneralSettings{
			Port: 22,
			User: "admin",
		},
	},
	Providers: []string{
		"aws",
	},
}
