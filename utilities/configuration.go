package utilities

import (
	"cimble/models"
	"fmt"
)

func FormatConfigurations(
	configurations []models.Configuration,
) (formattedConfigurations map[string]models.Configuration) {
	formattedConfigurations = make(map[string]models.Configuration)
	for _, configuration := range configurations {
		formattedConfigurations[configuration.Name] = models.Configuration{
			Info: configuration.Info,
		}
	}

	fmt.Printf("%+v\n", formattedConfigurations)
	return formattedConfigurations
}
