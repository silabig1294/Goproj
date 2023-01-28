package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies = []Movie{
	{//Movie
	ImdbID:      "tt4154796",
	Title:       "Avenger: Endgame",
	Year:		 2019,
	Rating:      8.9,
	IsSuperHero: true,
	},
}

func getAllMoviesHandler(c echo.Context) error {
	year1 := c.QueryParam("year")
	fmt.Println("year:",year1)
	if year1 == ""{
		return c.JSON(http.StatusOK,movies)
	}
	year2,err := strconv.Atoi(year1)
	if err != nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	ms := []Movie{}
	for _,m := range movies{
		if m.Year == year2 {
			ms = append(ms,m)
		}
	}
	return c.JSON(http.StatusOK,ms)
	// return c.JSON(200,[]string{"movie1","movie2"})
}

func getMoviesByIdHandler(c echo.Context) error {
	id := c.Param("id")
	for _,m := range movies{
		if m.ImdbID == id {
			return c.JSON(http.StatusOK,movies)
		}
	}

	return c.JSON(http.StatusNotFound,map[string]string{"message":"not found"})
}

func createMovieHandler(c echo.Context) error {
	m := &Movie{}
	//Bind io.readall unmarshall If we use library it will help you.
	if err := c.Bind(m);err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	movies = append(movies, *m)
	return c.JSON(http.StatusCreated,m)
}

func deleteMoviesByIdHandler(c echo.Context) error{
	
	return c.JSON(http.StatusOK,"success")
}


func main(){
	e := echo.New() // new struct
	e.GET("/movies",getAllMoviesHandler)
	e.GET("/movies/:id",getMoviesByIdHandler)
	e.POST("/movies",createMovieHandler)
	e.DELETE("/movies",deleteMoviesByIdHandler)
	port := "8080"
	log.Println("starting --- Port:",port)
	err := e.Start(":"+port)
	log.Fatal(err) // kill process. It cannot start server
}

