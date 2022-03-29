// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// SqlTuningAdvisorTaskSummaryFindingCounts The number of findings in the SQL Tuning Advisor summary report.
type SqlTuningAdvisorTaskSummaryFindingCounts struct {

	// The number of distinct SQL statements with recommended SQL profiles.
	RecommendedSqlProfile *int `mandatory:"true" json:"recommendedSqlProfile"`

	// The number of distinct SQL statements with implemented SQL profiles.
	ImplementedSqlProfile *int `mandatory:"true" json:"implementedSqlProfile"`

	// The number of distinct SQL statements with index recommendations.
	Index *int `mandatory:"true" json:"index"`

	// The number of distinct SQL statements with restructured SQL recommendations.
	Restructure *int `mandatory:"true" json:"restructure"`

	// The number of distinct SQL statements with stale or missing optimizer statistics recommendations.
	Statistics *int `mandatory:"true" json:"statistics"`

	// The number of distinct SQL statements with alternative plan recommendations.
	AlternatePlan *int `mandatory:"true" json:"alternatePlan"`
}

func (m SqlTuningAdvisorTaskSummaryFindingCounts) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningAdvisorTaskSummaryFindingCounts) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
