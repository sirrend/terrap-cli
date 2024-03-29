/*
Copyright © 2023 Sirrend
*/

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
	"github.com/sirrend/terrap-cli/internal/annotate"
	"github.com/sirrend/terrap-cli/internal/commons"
	"github.com/sirrend/terrap-cli/internal/files_handler"
	"github.com/sirrend/terrap-cli/internal/parser"
	"github.com/sirrend/terrap-cli/internal/receiver"
	"github.com/sirrend/terrap-cli/internal/scanning"
	"github.com/sirrend/terrap-cli/internal/state"
	"github.com/sirrend/terrap-cli/internal/utils"
	"github.com/sirrend/terrap-cli/internal/utils/cli"
	"github.com/sirrend/terrap-cli/internal/workspace"
	"github.com/spf13/cobra"
)

var (
	PRINTED                = false
	upgradeMessage         = ""
	notYetSupportedMessage = ""
)

// appliedRules is used to keep track of the rules applied in context
type appliedRules struct {
	ruleSet parser.RuleSet
	rules   []parser.Rule
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan your IaC code to find provider changes",

	Run: func(cmd *cobra.Command, args []string) {
		var workspace workspace.Workspace
		var asText []appliedRules
		asJson := map[string][]parser.Rule{}

		providersSet := cmd.Flag("fixed-providers").Changed
		if utils.IsInitialized(".") || providersSet {
			files, err := files_handler.FindResourcesPerFile(".")
			if err != nil {
				_, _ = commons.RED.Println(err)
			}

			// gather providers
			if !providersSet {
				err = state.Load("./.terrap.json", &workspace)
				if err != nil {
					_, _ = commons.RED.Println(err)
				}
			} else {
				workspace = cli.GetFixedProvidersFlag(*cmd)
			}

			// go over every provider in user's folder / user's declaration
			if len(workspace.Providers) > 0 {

				for provider, version := range workspace.Providers {
					rulebook, err := receiver.GetRules(provider, version.String())
					// validate rulebook downloaded
					if err != nil {
						if strings.Contains(err.Error(), utils.StripProviderPrefix(provider)) {
							notYetSupportedMessage = strings.Join([]string{notYetSupportedMessage, err.Error()}, ", ")
							continue
						}

						continue
					}

					flags := cli.ChangedComponentsFlags(*cmd)
					if !cmd.Flag("annotate").Changed {
						for file, fileResources := range files {
							if len(fileResources) == 0 {
								continue
							}

							for _, resource := range scanning.GetUniqueResources(fileResources) {
								if utils.IsItemInSlice(resource.Type, flags) {
									ruleset, err := resource.GetRuleset(rulebook, nil)
									if err != nil {
										_, _ = commons.RED.Println(err)
										os.Exit(1)
									}

									// fill json object with applied rules
									if cmd.Flag("json").Changed {
										if ruleset.Rules != nil {
											for _, rule := range ruleset.Rules {
												if apply, err := rule.DoesRuleApplyInContext(file, resource.Name, resource.Type); err == nil && apply {
													asJson[ruleset.ResourceName] = append(asJson[ruleset.ResourceName], rule)
													PRINTED = true
												}
											}
										}

										// combine ruleSets with applied rules
									} else {
										if ruleset.Rules != nil {
											var rules []parser.Rule
											for _, rule := range ruleset.Rules {
												if apply, err := rule.DoesRuleApplyInContext(file, resource.Name, resource.Type); err == nil && apply {
													rules = append(rules, rule)
													PRINTED = true
												}
											}

											if len(rules) > 0 {
												asText = append(asText, appliedRules{
													ruleSet: ruleset,
													rules:   rules,
												})
											}
										}
									}
								} else {
									PRINTED = true // to avoid wrong possible upgrade message
								}
							}

							// print json object
							if cmd.Flag("json").Changed {
								if len(asJson) != 0 {
									utils.PrettyPrintStruct(map[string]interface{}{file: asJson})
								}

								asJson = map[string][]parser.Rule{} // reset for next provider

							} else {
								if len(asText) > 0 {
									_, _ = commons.SIRREND.Println("* File:", utils.GetAbsPath(file))
								}

								for _, appliedRules := range asText {
									if len(appliedRules.rules) > 0 {
										appliedRules.ruleSet.PrettyPrint(appliedRules.rules)
									}
								}

								asText = []appliedRules{} // clean up for next iteration
							}
						}

						// check if upgrade is possible for the provider in context
						if !PRINTED {
							if rulebook.TargetVersion != "" {
								upgradeMessage += fmt.Sprintf("  %s  %s: ",
									emoji.UpArrow, utils.StripProviderPrefix(provider))
								upgradeMessage += commons.GREEN.Sprintf("%s -> %s\n", version, rulebook.TargetVersion)

								PRINTED = false // for next provider
							}
						}

					} else {
						for _, fileResources := range files {
							for _, resource := range fileResources {
								ruleset, err := resource.GetRuleset(rulebook, nil)
								if err != nil {
									_, _ = commons.RED.Println(err)
									os.Exit(1)
								}

								annotate.AddAnnotationByRuleSet(resource, ruleset)
							}
						}
					}
				}
			} else {
				_, _ = commons.YELLOW.Println("Couldn't find any providers under the provided context.")
			}

			// print safe upgrade message
			if !cmd.Flag("no-safe-upgrade-message").Changed && !cmd.Flag("no-messages").Changed {
				longestLine := 45

				if upgradeMessage != "" {
					for _, line := range strings.Split(upgradeMessage, "\n") {
						if len(line) > longestLine {
							longestLine = len(line)
						}
					}

					utils.PrintCharacterXTimes("-", longestLine)
					_, _ = commons.SIRREND.Println("The following providers are safe to upgrade: ")
					fmt.Println(upgradeMessage)
					utils.PrintCharacterXTimes("-", longestLine)
				}
			}

			// print not supported message
			if !cmd.Flag("no-not-supported-message").Changed && !cmd.Flag("no-messages").Changed {
				if notYetSupportedMessage != "" {
					message := strings.TrimLeft(notYetSupportedMessage, ", ")
					_, _ = commons.SIRREND.Print("The following providers and corresponding versions are not yet supported by terrap: ")
					fmt.Println(message, emoji.CryingFace.String())
					_, _ = commons.SIRREND.Print("Check again soon! ")
					fmt.Println("We're actively working on increasing our Providers support " + emoji.BuildingConstruction.String())
				}
			}

		} else {
			yellow := color.New(color.FgYellow)
			_, _ = yellow.Println("Hmm..seems like you didn't setup this folder yet.\nPlease execute < terrap init >.")
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// utility flags
	scanCmd.Flags().BoolP("json", "j", false, "Print scan output as json.")
	scanCmd.Flags().BoolP("annotate", "a", false, "Annotate the code itself.")
	scanCmd.Flags().BoolP("provider", "p", false, "Show only provider changes.")
	scanCmd.Flags().BoolP("data-sources", "d", false, "Show only data source changes.")
	scanCmd.Flags().BoolP("resources", "r", false, "Show only resources changes.")
	scanCmd.Flags().StringSliceP("fixed-providers", "f", []string{}, "A comma separated list of fixed providers written in the following format: `<provider>:<version>`.If this flag is used, all other in-context providers are ignored.")

	// extra output flags
	scanCmd.Flags().Bool("no-safe-upgrade-message", false, "Don't print which providers are safe to upgrade.")
	scanCmd.Flags().Bool("no-not-supported-message", false, "Don't print if providers are not supported.")
	scanCmd.Flags().BoolP("no-messages", "n", false, "Don't print any message other than pure command output.")
}
