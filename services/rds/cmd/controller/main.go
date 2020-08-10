// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package main

import (
	"os"

	ackrt "github.com/aws/aws-controllers-k8s/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrlrt "sigs.k8s.io/controller-runtime"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"

	svcresource "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource"
	svctypes "github.com/aws/aws-controllers-k8s/services/rds/apis/v1alpha1"

	 _ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_cluster_parameter_group"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/option_group"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/event_subscr_ip_tion"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_parameter_group"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/global_cluster"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_snapshot"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_cluster_snapshot"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_subnet_group"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_instance"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_instance_read_replica"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_proxy"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_cluster"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_security_group"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/db_cluster_endpoint"
	_ "github.com/aws/aws-controllers-k8s/services/rds/pkg/resource/custom_availability_zone"
	
)

var (
	awsServiceAPIGroup = "rds.services.k8s.aws"
	awsServiceAlias    = "rds"
	scheme             = runtime.NewScheme()
	setupLog           = ctrlrt.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = svctypes.AddToScheme(scheme)
}

func main() {
	var ackCfg ackrt.Config
	ackCfg.BindFlags()
	flag.Parse()
	ackCfg.SetupLogger()

	if err := ackCfg.Validate(); err != nil {
		setupLog.Error(
			err, "Unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	mgr, err := ctrlrt.NewManager(ctrlrt.GetConfigOrDie(), ctrlrt.Options{
		Scheme:             scheme,
		Port:               ackCfg.BindPort,
		MetricsBindAddress: ackCfg.MetricsAddr,
		LeaderElection:     ackCfg.EnableLeaderElection,
		LeaderElectionID:   awsServiceAPIGroup,
	})
	if err != nil {
		setupLog.Error(
			err, "unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	stopChan := ctrlrt.SetupSignalHandler()

	setupLog.Info(
		"initializing service controller",
		"aws.service", awsServiceAlias,
	)
	sc := ackrt.NewServiceController(
		awsServiceAlias, awsServiceAPIGroup,
	).WithLogger(
		ctrlrt.Log,
	).WithResourceManagerFactories(
		svcresource.GetManagerFactories(),
	)
	if err = sc.BindControllerManager(mgr, ackCfg); err != nil {
		setupLog.Error(
			err, "unable bind to controller manager to service controller",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	setupLog.Info(
		"starting manager",
		"aws.service", awsServiceAlias,
	)
	if err := mgr.Start(stopChan); err != nil {
		setupLog.Error(
			err, "unable to start controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}
}
