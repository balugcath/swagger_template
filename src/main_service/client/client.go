package client

import (
	"context"

	"github.com/go-openapi/strfmt"

	"swagger_template/src/main_service/swagger/client_users"
	"swagger_template/src/main_service/swagger/client_users/main_service"
)

// Client ...
type Client int

// Delete ...
func (Client) Delete(ctx context.Context, id string) error {
	param := strfmt.UUID(id)

	_, err := client_users.New(client_users.Config{}).
		MainService.DeleteUser(
		ctx,
		main_service.NewDeleteUserParams().
			WithBody(main_service.DeleteUserBody{UserID: &param}))
	return err
}
