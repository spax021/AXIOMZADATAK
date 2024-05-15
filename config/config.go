package config

import (
	"bufio"
	"os"
	"strings"
)

type Properties struct {
	UserBaseUrl       string
	OrderBaseUrl      string
	Username          string
	Password          string
	IncorrectPassword string
}

func ReadPropertiesFile() (*Properties, error) {
	file, err := os.Open("../test_resources/application.properties")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := &Properties{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			switch key {
			case "userBaseUrl":
				properties.UserBaseUrl = value
			case "orderBaseUrl":
				properties.OrderBaseUrl = value
			case "username":
				properties.Username = value
			case "password":
				properties.Password = value
			case "incorectPassword":
				properties.IncorrectPassword = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}
