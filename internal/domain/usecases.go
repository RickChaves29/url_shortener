package domain

import "github.com/RickChaves29/url_shortener/utils"

type IRepository interface {
	Create(originUrl, hashUrl string) error
	FindByHashUrl(hashUrl string) (UrlEntity, error)
}
type usecase struct {
	repository IRepository
}

func NewUsecase(r IRepository) *usecase {
	return &usecase{
		repository: r,
	}
}

func (uc usecase) CreateNewUrl(originUrl string) (string, error) {
	hashCodeUrl := utils.GenarateCode(6)
	err := uc.repository.Create(originUrl, hashCodeUrl)
	if err != nil {
		return "", err
	}
	entity, err := uc.repository.FindByHashUrl(hashCodeUrl)
	if err != nil {
		return "", err
	}
	return entity.HashUrl, nil
}

func (uc usecase) GetOriginUrlFromRedirect(hashUrl string) (string, error) {
	entity, err := uc.repository.FindByHashUrl(hashUrl)
	if err != nil {
		return "", err
	}
	return entity.OriginUrl, nil
}
