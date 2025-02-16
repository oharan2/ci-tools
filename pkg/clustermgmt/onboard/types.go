package onboard

import (
	"fmt"
	"path/filepath"

	"github.com/openshift/ci-tools/pkg/api"
)

const (
	BuildFarm                  = "build-farm"
	BuildUFarm                 = "build_farm"
	CI                         = "ci"
	CIOperator                 = "ci-operator"
	ClusterDisplay             = "cluster-display"
	ConfigUpdater              = "config-updater"
	GithubLdapUserGroupCreator = "github-ldap-user-group-creator"
	Master                     = "master"
	PodScaler                  = "pod-scaler"
	PromotedImageGovernor      = "promoted-image-governor"
)

func ServiceAccountKubeconfigPath(serviceAccount, clusterName string) string {
	return ServiceAccountFile(serviceAccount, clusterName, "config")
}

func ServiceAccountFile(serviceAccount, clusterName, fileType string) string {
	return fmt.Sprintf("sa.%s.%s.%s", serviceAccount, clusterName, fileType)
}

func ServiceAccountTokenFile(serviceAccount, clusterName string) string {
	return ServiceAccountFile(serviceAccount, clusterName, "token.txt")
}

func RepoMetadata() *api.Metadata {
	return &api.Metadata{
		Org:    "openshift",
		Repo:   "release",
		Branch: "master",
	}
}

func BuildFarmDirFor(releaseRepo, clusterName string) string {
	return filepath.Join(releaseRepo, "clusters", "build-clusters", clusterName)
}
