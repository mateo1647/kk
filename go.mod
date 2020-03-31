// This is a generated file. Do not edit directly.

module github.com/mateo1647/kk

go 1.13

require (
	github.com/canopytax/ckube v0.4.4
	github.com/fatih/color v1.9.0
	github.com/guessi/kubectl-grep v1.2.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/manifoldco/promptui v0.7.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.0
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	k8s.io/api v0.16.8
	k8s.io/apimachinery v0.16.8
	k8s.io/cli-runtime v0.0.0-20190918162238-f783a3654da8
	k8s.io/client-go v0.16.8
	sigs.k8s.io/kustomize v2.0.3+incompatible
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190918162238-f783a3654da8
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)
