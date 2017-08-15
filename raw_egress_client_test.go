package loggregator_test

import (
	"context"

	loggregator "code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"code.cloudfoundry.org/go-loggregator/testhelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RawEgressClient", func() {
	var (
		client *loggregator.RawEgressClient
		server *testhelpers.TestEgressServer
	)

	BeforeEach(func() {
		var err error
		server, err = testhelpers.NewTestEgressServer(
			fixture("server.crt"),
			fixture("server.key"),
			fixture("CA.crt"),
		)
		Expect(err).NotTo(HaveOccurred())

		err = server.Start(rxCallbackStub)
		Expect(err).NotTo(HaveOccurred())

		tlsConfig, err := loggregator.NewIngressTLSConfig(
			fixture("CA.crt"),
			fixture("client.crt"),
			fixture("client.key"),
		)
		Expect(err).NotTo(HaveOccurred())

		client, _, err = loggregator.NewEgressClient(
			server.Addr(),
			tlsConfig,
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		server.Stop()
	})

	It("creates a streaming client receiver", func() {
		_, err := client.Receiver(
			context.Background(),
			&loggregator_v2.EgressRequest{},
		)

		Expect(err).NotTo(HaveOccurred())
	})

	It("creates a streaming batch client receiver", func() {
		_, err := client.BatchReceiver(
			context.Background(),
			&loggregator_v2.EgressBatchRequest{},
		)

		Expect(err).NotTo(HaveOccurred())
	})
})

func rxCallbackStub(
	*loggregator_v2.EgressRequest,
	loggregator_v2.Egress_ReceiverServer,
) error {

	return nil
}