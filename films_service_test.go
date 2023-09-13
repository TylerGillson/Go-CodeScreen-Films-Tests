package films

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFilms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Films Service Test Suite")
}

var _ = Describe("Films API", func() {

	Context("Films API Service Tests", func() {
		films := Film{}

		It("Best rated film for director Ridley Scott", func() {
			Expect(films.getBestRatedFilm("Ridley Scott")).Should(Equal("Gladiator"))
		})

		It("Director with the most films", func() {
			Expect(films.getDirectorWithMostFilms()).Should(Equal("Ridley Scott"))
		})

		It("Average rating for films directed by Ridley Scott", func() {
			Expect(films.getAverageRating("Ridley Scott")).Should(Equal(6.9))
		})

		It("Shortest number of days between film releases by Ridley Scott", func() {
			Expect(films.getShortestNumberOfDaysBetweenFilmReleases("Ridley Scott")).Should(Equal(29))
		})

	})

})
