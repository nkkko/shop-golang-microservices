package configurations

import (
	"context"
	"github.com/labstack/echo/v4"
	v1 "github.com/meysamhadeli/shop-golang-microservices/internal/services/identity-service/identity/features/registering_user/endpoints/v1"
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/identity-service/shared/contracts"
)

func (c *usersModuleConfigurator) configEndpoints(ctx context.Context, group *echo.Group) {

	userEndpointBase := &contracts.IdentityEndpointBase[contracts.InfrastructureConfiguration]{
		ProductsGroup: group,
		Configuration: *c.InfrastructureConfiguration,
	}

	registerUserEndpoint := v1.NewCreteUserEndpoint(userEndpointBase)
	registerUserEndpoint.MapRoute()
}