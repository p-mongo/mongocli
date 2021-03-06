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
	"fmt"

	"github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/file"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type opsManagerAutomationUpdateOpts struct {
	globalOpts
	filename string
	fs       afero.Fs
	store    store.AutomationUpdater
}

func (opts *opsManagerAutomationUpdateOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerAutomationUpdateOpts) Run() error {
	newConfig := new(opsmngr.AutomationConfig)
	err := file.Load(opts.fs, opts.filename, newConfig)
	if err != nil {
		return err
	}

	if err = opts.store.UpdateAutomationConfig(opts.ProjectID(), newConfig); err != nil {
		return err
	}

	fmt.Print(deploymentStatus(config.OpsManagerURL(), opts.ProjectID()))

	return nil
}

// mongocli om cluster(s) update --projectId projectId --file myfile.yaml
func OpsManagerAutomationUpdateBuilder() *cobra.Command {
	opts := &opsManagerAutomationUpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:    "update",
		Short:  description.UpdateOMCluster,
		Hidden: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flags.File, flags.FileShort, "", "Filename to use")

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.File)

	return cmd
}
