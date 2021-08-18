package models

type Movie struct{
	ID int `json:"id"`,
	Title string  `json:"title"`, 
	Description string `json:"description"`, 
	Year int `json:"year"`,
	ReleaseDate time.time `json:"release_date"`,
	Runtime int `json:"runtime"`, 
	Rating int `json:"rating"`, 
	MPAARating string `json:"mpaa_rating"`,
	CreatedAt time.time `json:"created_at"`, 
	UpdatedAt time.time `json:"updated_at"`, 
	MovieGenre []MovieGenre `json:"-"`,

}

type Genre struct{
	ID int `json:"id"`,
	GenreName string `json:"genre_name"`,
	CreatedAt time.time `json:"created_at"`, 
	UpdatedAt time.time `json:"updated_at"`, 
}

type MovieGenre struct{
	ID int `json:"id"`,
	MovieId int `json:"movie_id"`,
	CreatedAt time.time `json:"created_at"`, 
	UpdatedAt time.time `json:"updated_at"`, 

}