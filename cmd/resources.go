package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/mateo1647/kk/resources"
	"github.com/mateo1647/kk/util"

	"github.com/spf13/cobra"
)

var (
	serviceCmd = &cobra.Command{
		Use:     "service",
		Aliases: []string{"services", "svc"},
		Short:   "Service list with pod details",
		Long:    `shows pod details along with service details`,
		Run: func(cmd *cobra.Command, args []string) {
			if namespace == "" {
				namespace = util.GetDefaultNamespace()
			}

			var keyword string

			if len(args) >= 1 && args[0] != "" {
				keyword = util.TrimQuoteAndSpace(args[0])
			}

			serviceResults := resources.GetServicesandPods(searchOptions, keyword)

			templates := &promptui.SelectTemplates{
				Active:   "{{ .Service.Name | underline | yellow }}",
				Inactive: "{{ .Service.Name }}",
				Details: `
--------- Pods ----------
{{ .Headerline }}{{ range $i, $pod := .PodResponse }}
{{ .StatusLine }}{{end}}`,
			}

			prompt := promptui.Select{
				Label:     "SERVICE NAME",
				Items:     serviceResults,
				Templates: templates,
				Size:      20,
			}

			i, _, err := prompt.Run()

			if err != nil {
				return
			}
			output := util.RawK8sOutput(namespace, context, labels, "get", "service", serviceResults[i].Service.Name, "-oyaml")
			for _, line := range output {
				fmt.Println(line)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(serviceCmd)
}
