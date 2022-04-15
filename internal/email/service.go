package email

import "github.com/Mustafa0831/TSarka/pkg/util"

//Service ...
type Service struct {
	Repository Email
}

//NewService ...
func NewService() *Service {
	return &Service{
		Repository: NewEmailRepository(),
	}
}

//EmailService ...
type EmailService interface {
	FindEmailFromText(text string) (string, error)
	FindIinFromText(text string) (string, error)
}

//FindEmailFromText ..
func (s *Service) FindEmailFromText(text string) (string, error) {
	email := util.FindEmailFromText(text)
	return email, nil
}

//FindIinFromText ...
func (s *Service) FindIinFromText(text string) (string, error) {
	iin := util.FindIinFromText(text)
	return iin, nil
}
