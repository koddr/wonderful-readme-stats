package main

import (
	"strconv"
)

// Config represents the configuration of the application.
type Config struct {
	GithubToken string
	Repository  *repository
	Server      *server
	Avatar      *avatar
	OutputImage *outputImage
}

// repository represents the GitHub repository of the application.
type repository struct {
	Owner, Name string
}

// server represents the server configuration of the application.
type server struct {
	Port, ReadTimeout, WriteTimeout int
}

// avatar represents the avatar configuration of the application.
type avatar struct {
	Shape                                  string
	Size, HorizontalMargin, VerticalMargin int
	RoundedRadius                          float64
}

// outputImage represents the output image configuration of the application.
type outputImage struct {
	MaxPerRow, MaxRows, UpdateInterval int
}

// validateEnvVariables initializes and validates the configuration from environment variables.
//
// It creates a new instance of the Config struct and populates it with values from environment variables.
// The function parses various environment variables and assigns them to the corresponding fields in the Config struct.
// It returns the populated Config struct and a nil error if the parsing is successful.
// If any parsing error occurs, it returns a nil Config struct and the corresponding error.
func validateEnvVariables() (*Config, error) {
	// Create a new instance of the Config struct.
	c := &Config{
		GithubToken: helpGetEnv("GITHUB_TOKEN", ""),
		Repository: &repository{
			Owner: helpGetEnv("REPOSITORY_OWNER", "koddr"),
			Name:  helpGetEnv("REPOSITORY_NAME", "wonderful-readme-stats"),
		},
		Server: &server{},
		Avatar: &avatar{
			Shape: helpGetEnv("AVATAR_SHAPE", "rounded"),
		},
		OutputImage: &outputImage{},
	}

	var err error

	// Parse the SERVER_PORT environment variable and assign it to c.Server.Port.
	c.Server.Port, err = strconv.Atoi(helpGetEnv("SERVER_PORT", "8080"))
	if err != nil {
		return nil, err
	}

	// Parse the SERVER_READ_TIMEOUT environment variable and assign it to c.Server.ReadTimeout.
	c.Server.ReadTimeout, err = strconv.Atoi(helpGetEnv("SERVER_READ_TIMEOUT", "5"))
	if err != nil {
		return nil, err
	}

	// Parse the SERVER_WRITE_TIMEOUT environment variable and assign it to c.Server.WriteTimeout.
	c.Server.WriteTimeout, err = strconv.Atoi(helpGetEnv("SERVER_WRITE_TIMEOUT", "10"))
	if err != nil {
		return nil, err
	}

	// Parse the AVATAR_SIZE environment variable and assign it to c.Avatar.Size.
	c.Avatar.Size, err = strconv.Atoi(helpGetEnv("AVATAR_SIZE", "64"))
	if err != nil {
		return nil, err
	}

	// Parse the AVATAR_HORIZONTAL_MARGIN environment variable and assign it to c.Avatar.HorizontalMargin.
	c.Avatar.HorizontalMargin, err = strconv.Atoi(helpGetEnv("AVATAR_HORIZONTAL_MARGIN", "12"))
	if err != nil {
		return nil, err
	}

	// Parse the AVATAR_VERTICAL_MARGIN environment variable and assign it to c.Avatar.VerticalMargin.
	c.Avatar.VerticalMargin, err = strconv.Atoi(helpGetEnv("AVATAR_VERTICAL_MARGIN", "12"))
	if err != nil {
		return nil, err
	}

	// Parse the AVATAR_ROUNDED_RADIUS environment variable and assign it to c.Avatar.RoundedRadius.
	c.Avatar.RoundedRadius, err = strconv.ParseFloat(helpGetEnv("AVATAR_ROUNDED_RADIUS", "16.0"), 64)
	if err != nil {
		return nil, err
	}

	// Parse the OUTPUT_IMAGE_MAX_PER_ROW environment variable and assign it to c.OutputImage.MaxPerRow.
	c.OutputImage.MaxPerRow, err = strconv.Atoi(helpGetEnv("OUTPUT_IMAGE_MAX_PER_ROW", "16"))
	if err != nil {
		return nil, err
	}

	// Parse the OUTPUT_IMAGE_MAX_ROWS environment variable and assign it to c.OutputImage.MaxRows.
	c.OutputImage.MaxRows, err = strconv.Atoi(helpGetEnv("OUTPUT_IMAGE_MAX_ROWS", "2"))
	if err != nil {
		return nil, err
	}

	// Parse the OUTPUT_IMAGE_UPDATE_INTERVAL environment variable and assign it to c.OutputImage.UpdateInterval.
	c.OutputImage.UpdateInterval, err = strconv.Atoi(helpGetEnv("OUTPUT_IMAGE_UPDATE_INTERVAL", "3600"))
	if err != nil {
		return nil, err
	}

	// Return the populated Config struct and nil error, indicating success.
	return c, nil
}
