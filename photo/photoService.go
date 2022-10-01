package photo

type Service interface {
	FindAll() ([]Photo, error)
	FindByID(ID int) (Photo, error)
	Create(photoRequest AddPhoto) (Photo, error)
	Update(ID int, photoRequest AddPhoto) (Photo, error)
	Delete(ID int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Photo, error) {
	photos, err := s.repository.FindAll()
	return photos, err
}

func (s *service) FindByID(ID int) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	return photo, err
}

func (s *service) Create(photoRequest AddPhoto) (Photo, error) {

	photo := Photo{
		Title:    photoRequest.Title,
		Caption:  photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
		UserID:   photoRequest.UserID,
	}

	newPhoto, err := s.repository.Create(photo)
	return newPhoto, err
}

func (s *service) Update(ID int, photoRequest AddPhoto) (Photo, error) {
	photo, err := s.repository.FindByID(ID)

	photo.Title = photoRequest.Title
	photo.Caption = photoRequest.Caption
	photo.PhotoUrl = photoRequest.PhotoUrl
	photo.UserID = photoRequest.UserID

	newPhoto, err := s.repository.Update(photo)
	return newPhoto, err
}

func (s *service) Delete(ID int) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	delPhoto, err := s.repository.Delete(photo)
	return delPhoto, err
}
