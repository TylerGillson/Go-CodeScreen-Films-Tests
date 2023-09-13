package films

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Films API", func() {

	Context("Films API Service Hidden Tests", func() {
		films := Film{}

		It("Best rated film for director Christopher Nolan", func() {
			Expect(films.getBestRatedFilm("Christopher Nolan")).Should(Equal("The Dark Knight"))
		})

		It("Average rating for films directed by Christopher Nolan", func() {
			Expect(films.getAverageRating("Christopher Nolan")).Should(Equal(8.3))
		})

		It("Shortest number of days between film releases by Christopher Nolan", func() {
			Expect(films.getShortestNumberOfDaysBetweenFilmReleases("Christopher Nolan")).Should(Equal(147))
		})

		It("Best rated film for unknown director", func() {
			Expect(films.getBestRatedFilm("Unknown Director")).Should(Equal(""))
		})

		It("Average rating for films directed by unknown director", func() {
			Expect(films.getAverageRating("Unknown Director")).Should(Equal(0.0))
		})

		It("Shortest number of days between film releases by unknown director", func() {
			Expect(films.getShortestNumberOfDaysBetweenFilmReleases("Unknown Director")).Should(Equal(0))
		})

	})

})
