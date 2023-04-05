package usecases

func (u *usecase) CheckPermission(username string) (string, error) {
	return u.repo.CheckPermission(username)
}
