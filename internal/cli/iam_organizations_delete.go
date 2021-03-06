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

package cli

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type iamOrganizationsDeleteOpts struct {
	*deleteOpts
	store store.OrganizationDeleter
}

func (opts *iamOrganizationsDeleteOpts) init() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *iamOrganizationsDeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteOrganization)
}

// mongocli iam organization(s) delete [id] [--orgId orgId]
func IAMOrganizationsDeleteBuilder() *cobra.Command {
	opts := &iamOrganizationsDeleteOpts{
		deleteOpts: &deleteOpts{
			failMessage:    "Organization not deleted",
			successMessage: "Organization '%s' deleted\n",
		},
	}
	cmd := &cobra.Command{
		Use:   "delete [ID]",
		Short: description.DeleteOrganization,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			opts.entry = args[0]
			return opts.Confirm()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().BoolVar(&opts.confirm, flags.Force, false, usage.Force)

	return cmd
}
