package urls

import (
	"context"

	"github.com/Toha22BSK/UrlShortener/gen/models"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations"
	"github.com/Toha22BSK/UrlShortener/gen/restapi/operations/short_url"
	"github.com/go-openapi/runtime/middleware"
)

type UrlService interface {
	createShortLink(ctx context.Context, Url string) (string, error)
}

func Configure(api *operations.BackendCoreAPI, urlService UrlService) {
	api.ShortURLCreateShortURLHandler = short_url.CreateShortURLHandlerFunc(
		func(params short_url.CreateShortURLParams) middleware.Responder {
			if params.Data.URL == "" {
				return short_url.NewCreateShortURLBadRequest().
					WithPayload(&models.ErrorV1{Message: "Joke did not del in database"})
			}

			shortUrl, err := urlService.createShortLink(
				params.HTTPRequest.Context(),
				params.Data.URL,
			)
			if err != nil {
				return short_url.NewCreateShortURLInternalServerError().
					WithPayload(&models.ErrorV1{Message: "Ooops, something went wrong"})
			}

			return short_url.NewCreateShortURLOK().WithPayload(&models.ShortURL{Short: shortUrl})
		},
	)

}
