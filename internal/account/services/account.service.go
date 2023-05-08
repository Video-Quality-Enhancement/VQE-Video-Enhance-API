package services

import (
	"log"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/account/models"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/account/repositories"
)

type AccountService interface {
	ContinueWithGoogle(account *models.Account) error
	// TODO: add whatsapp number and discord id
}

type accountService struct {
	repository repositories.AccountRepository
}

func NewAccountService(repository repositories.AccountRepository) AccountService {
	return &accountService{repository}
}

func (service *accountService) ContinueWithGoogle(account *models.Account) error {

	// TODO: /services/gapi verifyTokenId

	err := service.repository.UpsertAccount(account)
	if err != nil {
		log.Println("Error continuing with google ", account)
		return err
	}

	log.Println("Continued with google with email: ", account.Email)
	return nil

}
