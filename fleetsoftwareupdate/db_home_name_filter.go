// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbHomeNameFilter Database home name to include in the discovery.
// '*' Wildcard is allowed for 'startsWith' or 'endsWith' filtering.
type DbHomeNameFilter struct {

	// List of Database home names to include in the discovery.
	Names []string `mandatory:"true" json:"names"`

	// INCLUDE or EXCLUDE the filter results in the discovery for DB targets.
	// Supported for 'FSUCOLLECTION' RESOURCE_ID filter only.
	Mode DbFleetDiscoveryFilterModeEnum `mandatory:"false" json:"mode,omitempty"`
}

// GetMode returns Mode
func (m DbHomeNameFilter) GetMode() DbFleetDiscoveryFilterModeEnum {
	return m.Mode
}

func (m DbHomeNameFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbHomeNameFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbFleetDiscoveryFilterModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDbFleetDiscoveryFilterModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbHomeNameFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbHomeNameFilter DbHomeNameFilter
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDbHomeNameFilter
	}{
		"DB_HOME_NAME",
		(MarshalTypeDbHomeNameFilter)(m),
	}

	return json.Marshal(&s)
}
