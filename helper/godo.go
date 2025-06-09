package helper

import (
	"fmt"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// CreateDroplet creates a droplet using GODO
func CreateDroplet(doToken, dropletName, region, size, image string) (*godo.Droplet, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: doToken})
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	dropletRequest := &godo.DropletCreateRequest{
		Name:   dropletName,
		Region: region,
		Size:   size,
		Image:  godo.DropletCreateImage{Slug: image},
	}

	droplet, _, err := client.Droplets.Create(oauthClient.Context(), dropletRequest)
	if err != nil {
		return nil, fmt.Errorf("error creating droplet: %v", err)
	}

	return droplet, nil
}
