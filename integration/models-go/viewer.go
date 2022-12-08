package models

import "github.com/beyondan/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
