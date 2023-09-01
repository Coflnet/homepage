package usecase

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

const configFile = "./config.yaml"

type Config struct {
	projects []Project
}

func NewConfig() (*Config, error) {
	c := Config{}

	err := c.initConfig()
	if err != nil {
		slog.Error("failed to read config file", "err", err)
		return nil, err
	}

	return &c, nil
}

func (c *Config) initConfig() error {

	viper.SetConfigFile(configFile)
	slog.Debug(fmt.Sprintf("set config file to: %s", configFile))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	slog.Info("successfully read config file")

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	// dump all config
	slog.Debug(fmt.Sprintf("config is: %v", viper.Get("projects")))

	projectsMap := viper.Get("projects")

	for i, project := range projectsMap.([]interface{}) {

		// convert project to a map
		projectMap := project.(map[string]interface{})
		desc := projectMap["description"]
		title := projectMap["title"]
		imagePath := projectMap["imagepath"]
        imageCaption := projectMap["imagecaption"]

		slog.Debug(fmt.Sprintf("project %d is %v", i, projectMap), "title", title, "desc", desc, "imagePath", imagePath, "imageCaption", imageCaption)

		c.projects = append(c.projects, Project{
			Title:       title.(string),
			Description: desc.(string),
			ImagePath:   imagePath.(string),
            ImageCaption: imageCaption.(string),
		})

		slog.Info(fmt.Sprintf("project %d is %v; %v", i, project, desc))
	}

	return nil
}

func (c *Config) ListProjects() []Project {
	slog.Info(fmt.Sprintf("loaded %d projects", len(c.projects)))
	return c.projects
}
