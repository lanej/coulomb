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
		adapter = RackAdapter{
			Application: func(env map[string]interface{}) (int, map[string]string, string, error) {
				return 200, map[string]string{}, "", nil
			},
		}

		client = Client{
			URL:     "http://example.org",
			Adapter: adapter,
		}
	})

	Describe("RackAdapter", func() {
		It("servers a GET request", func() {
			response := client.Get("/get")

			Expect(response.Success()).To(BeTrue())
			Expect(response.Env.Method).To(Equal("GET"))
			Expect(response.Env.URL.Path).To(Equal("/get"))
		})
	})
})
