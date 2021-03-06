package worker_drain

import (
	"fmt"
	. "tests/test_helpers"

	director "github.com/cloudfoundry/bosh-cli/director"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = WorkerDrainDescribe("Worker drain scenarios", func() {

	var (
		deployment          director.Deployment
		countRunningWorkers func() int
		kubectl             *KubectlRunner
		drainTypesSpec      = PathFromRoot("specs/drain-types.yml")
		storageClassSpec    string
	)

	BeforeEach(func() {
		var err error
		director := NewDirector(testconfig.Bosh)
		deployment, err = director.FindDeployment(testconfig.Bosh.Deployment)
		Expect(err).NotTo(HaveOccurred())
		countRunningWorkers = CountDeploymentVmsOfType(deployment, WorkerVmType, VmRunningState)

		kubectl = NewKubectlRunner(testconfig.Kubernetes.PathToKubeConfig)
		kubectl.CreateNamespace()

		Expect(countRunningWorkers()).To(Equal(3))
		Expect(AllBoshWorkersHaveJoinedK8s(deployment, kubectl)).To(BeTrue())
		storageClassSpec = PathFromRoot(fmt.Sprintf("specs/storage-class-%s.yml", testconfig.Iaas))
	})

	AfterEach(func() {
		kubectl.RunKubectlCommand("delete", "-f", drainTypesSpec)
		kubectl.RunKubectlCommand("delete", "namespace", kubectl.Namespace())
		kubectl.RunKubectlCommand("delete", "-f", storageClassSpec)
	})

	Specify("Workers are able to drain", func() {
		By("Deploying all of the drain types")
		Eventually(kubectl.RunKubectlCommand("create", "-f", storageClassSpec), "60s").Should(gexec.Exit(0))
		Eventually(kubectl.RunKubectlCommand("create", "-f", drainTypesSpec), "30s", "5s").Should(gexec.Exit(0))
		Eventually(kubectl.RunKubectlCommand("rollout", "status", "daemonset/fluentd-elasticsearch", "-w"), "900s").Should(gexec.Exit(0))

		By("Recreating all workers successfully")
		dir := NewDirector(testconfig.Bosh)
		deployment, err := dir.FindDeployment(testconfig.Bosh.Deployment)
		Expect(err).NotTo(HaveOccurred())
		err = deployment.Recreate(director.NewAllOrInstanceGroupOrInstanceSlug("worker", ""), director.RecreateOpts{Canaries: "0", MaxInFlight: "100%"})
		Expect(err).NotTo(HaveOccurred())
	})

})
