package films

const filmsEndpointUrl string = "https://toolbox.palette-adv.spectrocloud.com:5002/films"

// Your API token. Needed to successfully authenticate when calling the films endpoint.
// This needs to be included in the Authorization header in the request you send to the films endpoint.
const apiToken string = "8c5996d5-fb89-46c9-8821-7063cfbc18b1"

type Film struct {
}

// Retrieves the data for all films by calling the https://toolbox.palette-adv.spectrocloud.com:5002/films endpoint.
func GetFilms() []Film {
	//TODO Implement
	return nil
}

/**
Retrieves the name of the best rated film that was directed by the director with the given name.
If there are no films directed the the given director, this method should return an empty string.
Note there will only be one film with the best rating.
*/
func (film *Film) getBestRatedFilm(directorName string) string {
	//TODO Implement
	return ""
}

/**
Retrieves the name of the director who has directed the most films in the CodeScreen Film service.
*/
func (film *Film) getDirectorWithMostFilms() string {
	//TODO Implement
	return ""
}

/**
Retrieves the average rating for the films directed by the given director, rounded to 1 decimal place.
If there are no films directed by the given director, this method should return 0.0.
*/
func (film *Film) getAverageRating(directorName string) float64 {
	//TODO Implement
	return 0.0
}

/**
Retrieves the shortest number of days between any two film releases directed by the given director.
If there are no films directed the the given director, this method should return 0.
If there is only one film directed by the given director, return 0.
Note that no director released more than one film on any given day.

For example, if the service returns the following 3 films,

{
    "name": "Batman Begins",
    "length": 140,
    "rating": 8.2,
    "releaseDate": "2006-06-16",
    "directorName": "Christopher Nolan"
},
{
    "name": "Interstellar",
    "length": 169,
    "rating": 8.6,
    "releaseDate": "2014-11-07",
    "directorName": "Christopher Nolan"
},
{
    "name": "Prestige",
    "length": 130,
    "rating": 8.5,
    "releaseDate": "2006-11-10",
    "directorName": "Christopher Nolan"
}

then this method should return 147 for Christopher Nolan, as Prestige was released 147 days after Batman Begins.
*/
func (film *Film) getShortestNumberOfDaysBetweenFilmReleases(directorName string) int {
	//TODO Implement
	return 0
}
