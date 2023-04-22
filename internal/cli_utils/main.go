package cli_utils

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/enescakir/emoji"
	"github.com/hashicorp/go-version"
	"github.com/olekukonko/tablewriter"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
	validate "golang.org/x/mod/semver"
	"os"
	"strings"
)

// SirrendLogoPrint
/*
@brief:
	SirrendLogoPrint prints the Sirrend logo
*/
func SirrendLogoPrint() {
	myFigure := figure.NewColorFigure("Sirrend", "", "purple", true)
	myFigure.Print()
	fmt.Println()
}

// flagsToAPIRepresentation
/*
@brief:
	flagsToAPIRepresentation change flag names to the corresponding API json field name
@params:
	flags - []string - the flags list to change
@returns:
	[]string - the flags as stated under the API response json
*/
func flagsToAPIRepresentation(flags []string) []string {
	for i, flag := range flags {
		if flag == "data-sources" {
			flags[i] = "data"
		} else if flag == "resources" {
			flags[i] = "resource"
		}
	}

	return flags
}

// ChangedComponentsFlags
/*
@brief:
	ChangedComponentsFlags checks which components flag were changed by the user.
	The passed cmd must have the flags: provider, data-sources and resources under it
@params:
	cmd - cobra.Command - the Command to look for flags under
@returns:
	[]string - the resources to filter the output according to
*/
func ChangedComponentsFlags(cmd cobra.Command) []string {
	flags := []string{"provider", "data-sources", "resources"}

	var componentsShow []string
	for _, flag := range flags {
		if cmd.Flag(flag).Changed {
			componentsShow = append(componentsShow, flag)
		}
	}

	if len(componentsShow) > 0 {
		return flagsToAPIRepresentation(componentsShow)
	}

	return flagsToAPIRepresentation(flags)
}

// GetFixedProvidersFlag
/*
@brief:
	GetFixedProvidersFlag returns a parsed workspace with the providers inserted under the fixed-providers flag
@params:
	cmd - cobra.Command - the Command to look for flags under
@returns:
	workspace.Workspace - the workspace which holds all the providers found
*/
func GetFixedProvidersFlag(cmd cobra.Command) workspace.Workspace {
	var ws workspace.Workspace

	ws.Providers = make(map[string]*version.Version)

	fp, err := cmd.Flags().GetStringSlice("fixed-providers")
	if err != nil {
		_, _ = commons.RED.Println(emoji.CrossMark, " Couldn't parse the received list of providers")
		os.Exit(1)
	}

	for _, p := range fp {
		if strings.Contains(p, ":") {
			providerValue := strings.Split(p, ":")
			if len(providerValue) > 1 {
				if validate.IsValid("v" + providerValue[1]) {
					v, err := version.NewVersion(providerValue[1])
					if err != nil {
						if providerValue[1] != "" {
							_, _ = commons.RED.Println(emoji.CrossMark, " The received "+providerValue[0]+" version is malformed: "+providerValue[1])
						} else {
							_, _ = commons.RED.Println(emoji.CrossMark, " Didn't receive the provider version of "+providerValue[0])
						}

						os.Exit(1)
					}

					ws.Providers[providerValue[0]] = v
				} else {
					_, _ = commons.RED.Println(emoji.CrossMark, " The received "+providerValue[0]+" provider version is malformed: "+providerValue[1])
					os.Exit(1)
				}
			} else {
				_, _ = commons.RED.Println(emoji.CrossMark, " Provider:Version format received is malformed: "+p)
				os.Exit(1)
			}
		} else {
			_, _ = commons.RED.Println(emoji.CrossMark, " Provider format received is malformed: "+p)
			os.Exit(1)
		}
	}

	return ws
}

// GetTable
/*
@brief:
	GetTable returns a new initialized table
@params:
	headers - []string - the list of headers to append to the new table
@returns:
	*tablewriter.Table - the new table
*/
func GetTable(headers []string) *tablewriter.Table {
	var headerColors []tablewriter.Colors
	columnAlignment := []int{0}
	columnColors := []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}

	table := tablewriter.NewWriter(os.Stdout) // init table
	table.SetHeader(headers)                  // add headers

	// set colors
	for i := 0; i < len(headers); i++ {
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.Bold, tablewriter.BgMagentaColor})
		columnColors = append(columnColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})
	}

	columnColors = columnColors[:len(columnColors)-1] // trim last element as the first one is set before the iteration

	table.SetHeaderColor(headerColors...)
	table.SetColumnColor(columnColors...)

	// set alignment
	for i := 0; i < len(headers)-1; i++ {
		columnAlignment = append(columnAlignment, 1)
	}
	table.SetColumnAlignment(columnAlignment)

	return table
}
