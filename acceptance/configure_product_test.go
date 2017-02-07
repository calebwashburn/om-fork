package acceptance

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const propertiesJSON = `{
	".properties.something": {"value": "configure-me"},
	".a-job.job-property": {"value": {"identity": "username", "password": "example-new-password"} },
	".top-level-property": { "value": [ { "guid": "some-guid", "name": "max", "my-secret": {"secret": "headroom"} } ] }
}`

const productNetworkJSON = `{
  "singleton_availability_zone": {"name": "az-one"},
  "other_availability_zones": [{"name": "az-two" }, {"name": "az-three"}],
  "network": {"name": "network-one"}
}`

const resourceConfigJSON = `
{
  "some-job": {
	  "instances": 1,
		"persistent_disk": { "size_mb": "20480" },
    "instance_type": { "id": "m1.medium" }
  }
}`

var _ = Describe("configure-product command", func() {
	var (
		server                  *httptest.Server
		productPropertiesMethod string
		productPropertiesBody   []byte
		productNetworkMethod    string
		productNetworkBody      []byte
		resourceConfigMethod    string
		resourceConfigBody      []byte
	)

	BeforeEach(func() {
		server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			switch req.URL.Path {
			case "/uaa/oauth/token":
				w.Write([]byte(`{
				"access_token": "some-opsman-token",
				"token_type": "bearer",
				"expires_in": 3600
			}`))
			case "/api/v0/staged/products":
				w.Write([]byte(`[
					{
						"installation_name": "some-product-guid",
						"guid": "some-product-guid",
						"type": "cf"
					},
					{
						"installation_name": "p-bosh-installation-name",
						"guid": "p-bosh-guid",
						"type": "p-bosh"
					}
				]`))
			case "/api/v0/staged/products/some-product-guid/jobs":
				w.Write([]byte(`{
					"jobs": [
					  {
							"name": "not-the-job",
							"guid": "bad-guid"
						},
					  {
							"name": "some-job",
							"guid": "the-right-guid"
						}
					]
				}`))
			case "/api/v0/staged/products/some-product-guid/properties":
				var err error
				productPropertiesMethod = req.Method
				productPropertiesBody, err = ioutil.ReadAll(req.Body)
				Expect(err).NotTo(HaveOccurred())

				w.Write([]byte(`{}`))
			case "/api/v0/staged/products/some-product-guid/networks_and_azs":
				var err error
				productNetworkMethod = req.Method
				productNetworkBody, err = ioutil.ReadAll(req.Body)
				Expect(err).NotTo(HaveOccurred())

				w.Write([]byte(`{}`))
			case "/api/v0/staged/products/some-product-guid/jobs/the-right-guid/resource_config":
				var err error
				resourceConfigMethod = req.Method
				resourceConfigBody, err = ioutil.ReadAll(req.Body)
				Expect(err).NotTo(HaveOccurred())

				w.Write([]byte(`{}`))
			default:
				auth := req.Header.Get("Authorization")
				if auth != "Bearer some-opsman-token" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				out, err := httputil.DumpRequest(req, true)
				Expect(err).NotTo(HaveOccurred())
				Fail(fmt.Sprintf("unexpected request: %s", out))
			}
		}))
	})

	It("successfully configures any product", func() {
		command := exec.Command(pathToMain,
			"--target", server.URL,
			"--username", "some-username",
			"--password", "some-password",
			"--skip-ssl-validation",
			"configure-product",
			"--product-name", "cf",
			"--product-properties", propertiesJSON,
			"--product-network", productNetworkJSON,
			"--product-resources", resourceConfigJSON,
		)

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		Expect(session.Out).To(gbytes.Say("setting properties"))
		Expect(session.Out).To(gbytes.Say("finished setting properties"))

		Expect(productPropertiesMethod).To(Equal("PUT"))
		Expect(productPropertiesBody).To(MatchJSON(fmt.Sprintf(`{"properties": %s}`, propertiesJSON)))

		Expect(productNetworkMethod).To(Equal("PUT"))
		Expect(productNetworkBody).To(MatchJSON(fmt.Sprintf(`{"networks_and_azs": %s}`, productNetworkJSON)))

		Expect(resourceConfigMethod).To(Equal("PUT"))
		Expect(resourceConfigBody).To(MatchJSON(`{
        "instances": 1,
        "persistent_disk": {
          "size_mb": "20480"
        },
        "instance_type": {
          "id": "m1.medium"
        },
        "internet_connected": false,
        "elb_names": null
      }`))
	})
})
