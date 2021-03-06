// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package convert

import (
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
)

// Public constants
const (
	AdminDB = "admin"
)

// Private constants
const (
	roleSep             = "@"
	automationAgentName = "mms-automation"
	monitoringAgentName = "mms-monitoring-agent"
	backupAgentName     = "mms-backup-agent"
	defaultUserDatabase = "admin"
)

// BuildAtlasRoles converts the roles inside the array of string in an array of mongodbatlas.Role structs
// r contains roles in the format roleName@dbName
func BuildAtlasRoles(r []string) []atlas.Role {
	roles := make([]atlas.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := defaultUserDatabase
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}

// BuildOMRoles converts the roles inside the array of string in an array of opsmngr.Role structs
// r contains roles in the format roleName@dbName
func BuildOMRoles(r []string) []*om.Role {
	roles := make([]*om.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := defaultUserDatabase
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = &om.Role{
			Role:     roleName,
			Database: databaseName,
		}
	}
	return roles
}

func newBackupUser(password string) *om.MongoDBUser {
	// roles for Monitoring Agent
	// https://docs.opsmanager.mongodb.com/current/reference/required-access-monitoring-agent/
	return &om.MongoDBUser{
		Username:                   backupAgentName,
		Database:                   defaultUserDatabase,
		AuthenticationRestrictions: []string{},
		Mechanisms:                 []string{},
		InitPassword:               password,
		Roles: []*om.Role{
			{
				Database: AdminDB,
				Role:     "clusterAdmin",
			},
			{
				Database: AdminDB,
				Role:     "readAnyDatabase",
			},
			{
				Database: AdminDB,
				Role:     "userAdminAnyDatabase",
			},
			{
				Database: AdminDB,
				Role:     "readWrite",
			},
			{
				Database: AdminDB,
				Role:     "readWrite",
			},
		},
	}
}
func newMonitoringUser(password string) *om.MongoDBUser {
	// roles for Monitoring Agent
	// https://docs.opsmanager.mongodb.com/current/reference/required-access-monitoring-agent/
	return &om.MongoDBUser{
		Username:                   monitoringAgentName,
		Database:                   defaultUserDatabase,
		InitPassword:               password,
		AuthenticationRestrictions: []string{},
		Mechanisms:                 []string{},
		Roles: []*om.Role{
			{
				Database: AdminDB,
				Role:     "clusterMonitor",
			},
		},
	}
}
