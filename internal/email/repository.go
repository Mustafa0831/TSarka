package email

type emailRepository struct {
}

//NewEmailRepository ...
func NewEmailRepository() *emailRepository {
	return &emailRepository{}
}

type Email interface {
}
