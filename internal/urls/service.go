package urls

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Toha22BSK/UrlShortener/gen/database"
	"github.com/sirupsen/logrus"
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
			short_url, err = apiCfg.DB.SetActiveTrue(ctx, Url)
			if err != nil {
				return "", err
			}
			return short_url, nil
		}
		return "", err
	}
	return short_url, nil
}

func (apiCfg *apiConfig) createShortLinkExpire(
	ctx context.Context,
	Url string,
	date_expire time.Time,
) (string, error) {
	args := database.CreateUrlWithExpireParams{
		Url:       Url,
		ExpiresAt: date_expire,
	}
	argsExpire := database.SetActiveTrueExpireParams{
		Url:       Url,
		ExpiresAt: date_expire,
	}
	short_url, err := apiCfg.DB.CreateUrlWithExpire(ctx, args)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			short_url, err = apiCfg.DB.SetActiveTrueExpire(ctx, argsExpire)
			if err != nil {
				return "", err
			}
			return short_url, nil
		}
		return "", err
	}
	return short_url, nil
}

func (apiCfg *apiConfig) getUrl(ctx context.Context, short_url string) (string, error) {
	errorExp := apiCfg.DB.ChangeActiveStatus(ctx)
	if errorExp != nil {
		logrus.WithError(errorExp).Fatal("Error with change active status")
	}
	url, err := apiCfg.DB.GetUrlByToken(ctx, short_url)
	if err != nil {
		return "", err
	}
	if !url.IsActive {
		return "", errors.New("Url is not active")
	}
	return url.Url, nil
}

func (apiCfg *apiConfig) writeLog(ctx context.Context, short_url string, msg string) error {
	arg := database.WriteLogParams{
		TokenUrl: short_url,
		LogMsg:   msg,
	}
	err := apiCfg.DB.WriteLog(ctx, arg)
	if err != nil {
		return err
	}
	return nil
}

func (apiCfg *apiConfig) getAnalitics(ctx context.Context, short_url string) (int64, error) {
	count, err := apiCfg.DB.GetRedirectCount(ctx, short_url)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (apiCfg *apiConfig) deleteShortUrl(ctx context.Context, short_url string) error {
	err := apiCfg.DB.DeleteUrlByToken(ctx, short_url)
	if err != nil {
		return err
	}
	return nil
}
