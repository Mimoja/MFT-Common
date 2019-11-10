package MFTCommon

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

const CurrentImportDataDefinition = "0.2+lastImportTimestamp"
const CurrentFlashDataDefinition = "0.1+initial"

func DataDefinitionUpgradeRequired(latestKnown string, current string) (bool, error) {

	latestKnownVersion, err := version.NewVersion(latestKnown)
	if err != nil {
		return true, fmt.Errorf("Could not parse latestKnownVersion")
	}

	currentVersion, err := version.NewVersion(current)
	if err != nil {
		return true, fmt.Errorf("Could not parse currentVersion")
	}

	return currentVersion.LessThan(latestKnownVersion), nil
}
