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
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerMeasurementsProcessOpts struct {
	globalOpts
	measurementsOpts
	hostID string
	store  store.HostMeasurementLister
}

func (opts *opsManagerMeasurementsProcessOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerMeasurementsProcessOpts) Run() error {
	listOpts := opts.newProcessMeasurementListOptions()
	result, err := opts.store.HostMeasurements(opts.ProjectID(), opts.hostID, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli om|cm measurements process(es) hostId [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId]
func OpsManagerMeasurementsProcessBuilder() *cobra.Command {
	opts := &opsManagerMeasurementsProcessOpts{}
	cmd := &cobra.Command{
		Use:     "process",
		Short:   description.ProcessMeasurements,
		Aliases: []string{"processes"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.hostID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.granularity, flags.Granularity, "", usage.Granularity)
	cmd.Flags().StringVar(&opts.period, flags.Period, "", usage.Period)
	cmd.Flags().StringVar(&opts.start, flags.Start, "", usage.MeasurementStart)
	cmd.Flags().StringVar(&opts.end, flags.End, "", usage.MeasurementEnd)
	cmd.Flags().StringSliceVar(&opts.measurementType, flags.Type, nil, usage.MeasurementType)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
