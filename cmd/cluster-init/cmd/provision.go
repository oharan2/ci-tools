package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/openshift/ci-tools/pkg/clustermgmt"
	"github.com/openshift/ci-tools/pkg/clustermgmt/provision/aws"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type options struct {
	clusterInstall string
}

var (
	opts = options{}
)

func NewProvision(ctx context.Context, log *logrus.Entry) *cobra.Command {
	cmd := cobra.Command{
		Use:   "provision",
		Short: "Commands to provision the infrastructure on a cloud provider",
	}
	cmd.PersistentFlags().StringVar(&opts.clusterInstall, "cluster-install", "", "Path to cluster-install.yaml")
	cmd.MarkPersistentFlagRequired("cluster-install")
	cmd.AddCommand(newProvisionAWS(ctx, log))
	return &cmd
}

func newProvisionAWS(ctx context.Context, log *logrus.Entry) *cobra.Command {
	cmd := cobra.Command{
		Use:   "aws",
		Short: "Provision assets on AWS",
		Long: `Provision the required infrastructure on AWS.
An AWS profile must be properly set for these subcommands to work properly.
How to use a named profile: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-using-profiles
For more information regarding env. variables: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html#envvars-set`,
	}
	cmd.AddCommand(newAWSCreateStacks(ctx, log))
	return &cmd
}

func newAWSCreateStacks(ctx context.Context, log *logrus.Entry) *cobra.Command {
	cmd := cobra.Command{
		Use:   "create-stacks",
		Short: "Create cloud formation stacks",
		Long:  `Create cloud formation stacks `,
		RunE: func(cmd *cobra.Command, args []string) error {
			step := aws.NewCreateAWSStacksStep(log,
				func() (*clustermgmt.ClusterInstall, error) {
					log.WithField("path", opts.clusterInstall).Info("Loading cluster install")
					return clustermgmt.LoadClusterInstall(opts.clusterInstall)
				},
				func() (aws.CloudFormationClient, error) {
					log.Info("Loading AWS config")
					awsconfig, err := config.LoadDefaultConfig(ctx)
					if err != nil {
						return nil, fmt.Errorf("load aws config: %v", err)
					}
					return cloudformation.NewFromConfig(awsconfig), nil
				},
				nil, nil,
			)
			if err := step.Run(ctx); err != nil {
				return fmt.Errorf("%s: %v", step.Name(), err)
			}
			return nil
		},
	}
	return &cmd
}
