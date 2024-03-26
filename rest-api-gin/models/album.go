package models

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var Albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Year: 1957},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1954},
	{ID: 4, Title: "Mingus Ah Um", Artist: "Charles Mingus", Year: 1959},
	{ID: 5, Title: "A Love Supreme", Artist: "John Coltrane", Year: 1965},
}
