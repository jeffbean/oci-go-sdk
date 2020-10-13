// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v27/common"
)

// EntityTagSummary Summary of an entity tag.
type EntityTagSummary struct {

	// Unique tag key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The date and time the tag was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Name of the tag that matches the term name.
	Name *string `mandatory:"false" json:"name"`

	// URI to the tag instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" json:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" json:"termPath"`

	// Description of the related term.
	TermDescription *string `mandatory:"false" json:"termDescription"`

	// Unique id of the parent glossary of the term.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// State of the Tag.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The unique key of the parent entity.
	EntityKey *string `mandatory:"false" json:"entityKey"`
}

func (m EntityTagSummary) String() string {
	return common.PointerString(m)
}
