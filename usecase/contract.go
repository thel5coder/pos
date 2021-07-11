package usecase

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"majoo-test/domain/view_models"
	"majoo-test/pkg/jwe"
	"majoo-test/pkg/jwt"
	redisPkg "majoo-test/pkg/redis"
)

type Contract struct {
	ReqID         string
	UserID        string
	MerchantID    string
	RoleID        int
	App           *fiber.App
	DB            *sql.DB
	TX            *sql.Tx
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	Validate      *validator.Validate
	Translator    ut.Translator
	Redis         redisPkg.RedisClient
}

const (
	//default limit for pagination
	defaultLimit = 10

	//max limit for pagination
	maxLimit = 50

	//default order by
	defaultOrderBy = "id"

	//default sort
	defaultSort = "asc"

	//default last page for pagination
	defaultLastPage = 0
)

func (uc Contract) SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
