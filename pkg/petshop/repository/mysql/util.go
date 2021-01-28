package mysql

import (
	"strings"

	"github.com/funa1g/microservice-example/pkg/petshop/domain"
	"github.com/rs/zerolog/log"
)

func optimizeError(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "Error 1062") {
		return domain.ErrDuplicateEntry
	}
	log.Error().Msg(err.Error())
	return domain.ErrServerInternal
}
