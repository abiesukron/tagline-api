package web

type BeritaResponse struct {
	Id        int    `json:"id"`
	Judul     string `json:"judul"`
	Slug      string `json:"slug"`
	Excerpt   string `json:"excerpt"`
	Isi       string `json:"isi"`
	Tag       string `json:"tag"`
	Gambar    string `json:"gambar"`
	Kategori  string `json:"kategori"`
	Dibaca    string `json:"dibaca"`
	Publikasi string `json:"publikasi"`
	User      string `json:"user"`
	Dibuat    string `json:"dibuat"`
}
