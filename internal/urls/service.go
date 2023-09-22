package urls

import (
	"context"
	"strings"

	"github.com/Toha22BSK/UrlShortener/gen/database"
)

type apiConfig struct {
	DB *database.Queries
}

func NewApiConfig(db *database.Queries) *apiConfig {
	return &apiConfig{
		DB: db,
	}
}

func (apiCfg *apiConfig) createShortLink(ctx context.Context, Url string) (string, error) {
	short_url, err := apiCfg.DB.CreateUrl(ctx, Url)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			short_url, err = apiCfg.DB.GetUrlTokenByUrl(ctx, Url)
			if err != nil {
				return "", err
			}
			return short_url, nil
		}
		return "", err
	}
	return short_url, nil
}
