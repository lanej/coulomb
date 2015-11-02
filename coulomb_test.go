package coulumb_test

import (
	. "github.com/lanej/coulomb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coulomb", func() {
	var (
		client  Client
		adapter Adapter
	)

	BeforeEach(func() {
		adapter = RackAdapter{}

		client = Client{
			Url:     "http://example.org",
			Adapter: adapter,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})
	})
})
