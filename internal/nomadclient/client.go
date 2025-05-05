package nomadclient

import (
    "github.com/hashicorp/nomad/api"
)

func NewClient() (*api.Client, error) {
    return api.NewClient(api.DefaultConfig())
}
