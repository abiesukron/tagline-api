package web

type BeritaUpdateRequest struct {
	Id        int    `json:"id" validate:"required"`
	Judul     string `json:"judul" validate:"required,min=3,max=255"`
	Slug      string `json:"slug" validate:"required,min=3,max=255"`
	Excerpt   string `json:"excerpt" validate:"required,min=10"`
	Isi       string `json:"isi" validate:"required,min=20"`
	Tag       string `json:"tag" validate:"required"`
	Gambar    string `json:"gambar" validate:"required,min=3"`
	Kategori  string `json:"kategori" validate:"required,min=3"`
	Publikasi string `json:"publikasi" validate:"required"`
	User      string `json:"user" validate:"required"`
}
