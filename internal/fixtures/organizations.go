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

package fixtures

import (
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
)

var Organization = &om.Organization{
	ID:   "5a0a1e7e0f2912c554080adc",
	Name: "Organization 0",
	Links: []*atlas.Link{
		{
			Href: "https://cloud.mongodb.com/api/public/v1.0/orgs/5a0a1e7e0f2912c554080adc",
			Rel:  "self",
		},
	},
}

func Organizations() *om.Organizations {
	return &om.Organizations{
		Links: []*atlas.Link{
			{
				Href: "https://cloud.mongodb.com/api/atlas/v1.0/orgs",
				Rel:  "self",
			},
		},
		Results: []*om.Organization{
			Organization,
		},
		TotalCount: 1,
	}
}
