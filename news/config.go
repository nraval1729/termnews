package news

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

type config struct {
	ApiKey           string   `yaml:"apiKey"`
	Sources          []string `yaml:"sources"`
	Categories       []string `yaml:"categories"`
	Country          string   `yaml:"country"`
	RefreshFrequency int      `yaml:"refresh_frequency"`
}

const defaultRefreshFrequency = 15

func getConfig() (*config, error) {
	var config config
	err := config.getConfig(getConfigPath())
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *config) getConfig(path string) error {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("Unmarshal: %v", err)
	}

	err = c.validateConfig()
	return err
}

func (c *config) validateConfig() error {
	if c.ApiKey == "" {
		return fmt.Errorf("config missing API key")
	}

	if c.Sources == nil && c.Categories == nil && c.Country == "" {
		return fmt.Errorf("config needs at least one of [sources, (categories and country)]")
	}

	if c.Sources != nil && (c.Categories != nil || c.Country != "") {
		return fmt.Errorf("if config has sources present then it cannot have either categories or country present and vice versa\n" +
			"read more: https://newsapi.org/docs/endpoints/top-headlines")
	}

	if c.RefreshFrequency == 0 {
		c.RefreshFrequency = defaultRefreshFrequency
	}

	return nil
}

func getConfigPath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Printf("user.Current err #%v", err)
	}
	return path.Join(currentUser.HomeDir, ".termnews", "config.yml")
}
