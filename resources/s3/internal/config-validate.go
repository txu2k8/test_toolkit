package internal

import (
	"fmt"
	"strings"
)

// Check if version of the config is valid
func validateConfigVersion(config *configV10) (bool, string) {
	if config.Version != globalMCConfigVersion {
		return false, fmt.Sprintf("Config version '%s' does not match mc config version '%s', please update your binary.\n",
			config.Version, globalMCConfigVersion)
	}
	return true, ""
}

// Verifies the config file of the MinIO Client
func validateConfigFile(config *configV10) (bool, []string) {
	ok, err := validateConfigVersion(config)
	validationSuccessful := true
	var errors []string
	if !ok {
		validationSuccessful = false
		errors = append(errors, err)
	}
	aliases := config.Aliases
	for _, aliasConfig := range aliases {
		aliasConfigHealthOk, aliasErrors := validateConfigHost(aliasConfig)
		if !aliasConfigHealthOk {
			validationSuccessful = false
			errors = append(errors, aliasErrors...)
		}
	}
	return validationSuccessful, errors
}

func validateConfigHost(host aliasConfigV10) (bool, []string) {
	validationSuccessful := true
	var hostErrors []string
	if !isValidAPI(strings.ToLower(host.API)) {
		validationSuccessful = false
		hostErrors = append(hostErrors, errInvalidAPISignature(host.API, host.URL).ToGoError().Error())
	}
	if !isValidHostURL(host.URL) {
		validationSuccessful = false
		hostErrors = append(hostErrors, errInvalidURL(host.URL).ToGoError().Error())
	}
	return validationSuccessful, hostErrors
}
