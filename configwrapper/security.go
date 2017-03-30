package configwrapper

// A config struct for turning a security config source into authorization data
type AuthorizationConfig struct {
	Rules []struct {
		Id         string              `yaml:"Id"`
		Message    string              `yaml:"Message"`
		Operation  string              `yaml:"Operation"`
		Authorize  string              `yaml:"Authorize"`
		Aggregate  string              `yaml:"Aggregate"`
		Properties map[string][]string `yaml:"Property"`
	} `yaml:"Rules"`
}
