package deploy

import "fmt"

func validateComponents(components []string) error {
	for _, c := range components {
		if !in(c, DeployOrder) {
			return fmt.Errorf("[ %s ] is not a valid component", c)
		}
	}
	return nil
}

func validateElasticComponents(components []string) error {
	for _, c := range components {
		if !in(c, ElasticDeployOrder) {
			return fmt.Errorf("[ %s ] is not a valid component", c)
		}
	}
	return nil
}

func ValidateComponets(components []string, elastic bool) error {
	if in("all", components)
	if elastic {
		return validateElasticComponents(components)
	}
	return validateComponents(components)
}

func in(a string, b []string) bool {
	for _, i :=  range b {
		if i == a {
			return true
		}
	}
	return false
}
