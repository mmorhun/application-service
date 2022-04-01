module github.com/redhat-appstudio/application-service

go 1.16

require (
	github.com/brianvoe/gofakeit/v6 v6.9.0
	github.com/devfile/api/v2 v2.0.0-20211021164004-dabee4e633ed
	github.com/devfile/library v1.2.1-0.20211104222135-49d635cb492f
	github.com/devfile/registry-support/index/generator v0.0.0-20220222194908-7a90a4214f3e
	github.com/devfile/registry-support/registry-library v0.0.0-20220222194908-7a90a4214f3e
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-logr/logr v1.2.2
	github.com/google/go-cmp v0.5.7
	github.com/google/go-github/v41 v41.0.0
	github.com/migueleliasweb/go-github-mock v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/openshift/api v3.9.0+incompatible
	github.com/redhat-appstudio/build-service v0.0.0-20220330075120-4d8c279d5f71
	github.com/redhat-appstudio/service-provider-integration-scm-file-retriever v0.2.5
	github.com/redhat-developer/alizer/go v0.0.0-20220215154256-33df7feef4ae
	github.com/spf13/afero v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tektoncd/pipeline v0.33.0
	github.com/tektoncd/triggers v0.19.1
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/yaml v1.3.0

)

replace github.com/antlr/antlr4 => github.com/antlr/antlr4 v0.0.0-20211106181442-e4c1a74c66bd
