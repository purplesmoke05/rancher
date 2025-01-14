package user

import (
	"context"

	"github.com/rancher/rancher/pkg/controllers/user/authz"
	"github.com/rancher/rancher/pkg/controllers/user/dnsrecord"
	"github.com/rancher/rancher/pkg/controllers/user/endpoints"
	"github.com/rancher/rancher/pkg/controllers/user/healthsyncer"
	"github.com/rancher/rancher/pkg/controllers/user/helm"
	"github.com/rancher/rancher/pkg/controllers/user/nodesyncer"
	"github.com/rancher/rancher/pkg/controllers/user/secret"
	"github.com/rancher/rancher/pkg/controllers/user/workloadservice"
	"github.com/rancher/types/config"
)

func Register(ctx context.Context, cluster *config.UserContext) error {
	nodesyncer.Register(cluster)
	healthsyncer.Register(ctx, cluster)
	authz.Register(cluster)
	secret.Register(cluster)
	helm.Register(cluster)

	userOnlyContext := cluster.UserOnlyContext()
	//workload.Register(ctx, userOnlyContext)
	dnsrecord.Register(ctx, userOnlyContext)
	workloadservice.Register(ctx, userOnlyContext)
	endpoints.Register(ctx, userOnlyContext)

	return nil
}
