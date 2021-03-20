package main

import (
    "github.com/go-chi/chi"
    "go.uber.org/zap"
    "net/http"
)

type api struct{
    logger *zap.Logger
    offerService *offerService
}

func newApi(logger *zap.Logger, service *offerService) *api{
    return &api{
        logger: logger,
        offerService: service,
    }
}

func (a *api) mount(router chi.Router) {
    router.Route("/api/v1/hotel", func(r chi.Router) {
        r.Route("/offers", func(r chi.Router) {
            r.Post("/", a.handlePostOffer)
        })
    })
}

type CreateOfferRequest struct{
    //todo: fill this with fields required by API spec, you can use sth like https://mholt.github.io/json-to-go/
}

func (a *api) handlePostOffer(w http.ResponseWriter,r *http.Request){
    //todo: we check headers and generally do http-related stuff here.
    // 1. Check headers (e.g. content-type, user-headers
    // 2. Unmarshall request
    // 3. Call offerService and choose appropriate response here
    // 4. Give appropriate response


    panic("implement me")
}



