package urls

import (
	"context"
	"time"

	"github.com/Toha22BSK/UrlShortener/gen/models"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations/analytics"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations/short_url"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type UrlService interface {
	createShortLink(ctx context.Context, Url string) (string, error)
	createShortLinkExpire(ctx context.Context, Url string, date_expire time.Time) (string, error)
	getUrl(ctx context.Context, shortUrl string) (string, error)
	writeLog(ctx context.Context, shortUrl string, msg string) error
	getAnalitics(ctx context.Context, shortUrl string) (int64, error)
}

func Configure(api *operations.BackendCoreAPI, urlService UrlService) {
	api.ShortURLCreateShortURLHandler = short_url.CreateShortURLHandlerFunc(
		func(params short_url.CreateShortURLParams) middleware.Responder {
			if params.Data.URL == "" {
				return short_url.NewCreateShortURLBadRequest().
					WithPayload(&models.ErrorV1{Message: "Not correct params, url is empty"})
			}
			shortUrl := ""
			err := error(nil)
			minutes := params.Data.Minutes + params.Data.Hours*60 + params.Data.Days*24*60
			if minutes > 0 {
				shortUrl, err = urlService.createShortLinkExpire(
					params.HTTPRequest.Context(),
					params.Data.URL,
					time.Now().UTC().Add(time.Minute*time.Duration(minutes)),
				)
				if err != nil {
					return short_url.NewCreateShortURLInternalServerError().
						WithPayload(&models.ErrorV1{Message: "Ooops, something went wrong"})
				}
			} else {

				shortUrl, err = urlService.createShortLink(
					params.HTTPRequest.Context(),
					params.Data.URL,
				)
				if err != nil {
					return short_url.NewCreateShortURLInternalServerError().
						WithPayload(&models.ErrorV1{Message: "Ooops, something went wrong"})
				}
			}

			return short_url.NewCreateShortURLOK().WithPayload(&models.ShortURL{Short: shortUrl})
		},
	)
	api.ShortURLGetShortURLHandler = short_url.GetShortURLHandlerFunc(
		func(params short_url.GetShortURLParams) middleware.Responder {
			if params.ShortURL == "" {
				return short_url.NewCreateShortURLBadRequest().
					WithPayload(&models.ErrorV1{Message: "Not correct params, short url is empty"})
			}

			url, err := urlService.getUrl(
				params.HTTPRequest.Context(),
				params.ShortURL,
			)
			if err != nil {
				return short_url.NewCreateShortURLInternalServerError().
					WithPayload(&models.ErrorV1{Message: "Ooops, something went wrong"})
			}
			err = urlService.writeLog(
				params.HTTPRequest.Context(),
				params.ShortURL,
				"Redirect to "+url,
			)
			if err != nil {
				logrus.WithError(err).Fatal("failed to write log")
			}
			return short_url.NewGetShortURLFound().WithLocation(url)
		},
	)
	api.AnalyticsGetAnalyticsHandler = analytics.GetAnalyticsHandlerFunc(
		func(params analytics.GetAnalyticsParams) middleware.Responder {
			if params.ShortURL == "" {
				return short_url.NewCreateShortURLBadRequest().
					WithPayload(&models.ErrorV1{Message: "Not correct params, short url is empty"})
			}

			count, err := urlService.getAnalitics(
				params.HTTPRequest.Context(),
				params.ShortURL,
			)
			if err != nil {
				return analytics.NewGetAnalyticsInternalServerError().
					WithPayload(&models.ErrorV1{Message: "Ooops, something went wrong"})
			}
			if count == 0 {
				return analytics.NewGetAnalyticsNotFound().
					WithPayload("Nothing redirect with this short url")
			}

			return analytics.NewGetAnalyticsOK().WithPayload(&models.Analytics{Redirects: count})
		},
	)
}
