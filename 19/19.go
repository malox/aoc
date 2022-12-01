package main

import (
	"fmt"
	"helper"
	"regexp"
	"strconv"
	"strings"
)

type ruleList []string

func formatRules(irules ruleList) map[int]ruleList {
	parsed := map[int]ruleList{}
	for _, irule := range irules {
		splitted := strings.Split(irule, ": ")
		ruleid, _ := strconv.Atoi(splitted[0])
		parsed[ruleid] = strings.Split(splitted[1], " | ")
		// for _, pp := range parsed[ruleid] {
		// 	fmt.Println("  - parsed key ", ruleid, " - rules ", pp)
		// }
	}
	return parsed
}

// -----------------------------------------------------------------

func parse(irules, ivalues []string) {
	// fmt.Println("irules : ", irules)
	// fmt.Println("ivalues : ", ivalues)

	tmpRules := formatRules(irules)
	// fmt.Println("tmpRules : ", tmpRules)
	// for id, rules := range tmpRules {
	// 	fmt.Println("tmprule ", id, " - len ", len(rules))
	// 	for _, rule := range rules {
	// 		fmt.Println("        - tmpvalue ", rule)
	// 	}
	// }

	fmt.Printf("%s\n\nBuilding rules\n\n", strings.Repeat("-", 23))

	rules := map[int]ruleList{}
	for len(rules) != len(tmpRules) {
		for tmpruleid, tmprules := range tmpRules {
			// fmt.Printf("%s\n\n", strings.Repeat("   -", 5))
			_, alreadyexists := rules[tmpruleid]
			if alreadyexists {
				continue
			}
			if len(tmprules) == 1 && strings.Contains(tmprules[0], "\"") {
				rules[tmpruleid] = ruleList{strings.ReplaceAll(tmprules[0], "\"", "")}
				// fmt.Println("adding rule ", tmpruleid, " with value ", rules[tmpruleid])
				continue
			}
			addRule := true
			rulesToAdd := ruleList{}
			for _, rule := range tmprules {
				tmpRulesToAdd := ruleList{""}
				for _, rulestrref := range strings.Split(rule, " ") {
					ruleref, _ := strconv.Atoi(rulestrref)
					currRules, rulesExists := rules[ruleref]
					if rulesExists {
						// fmt.Println("currRules ", len(currRules), " = ", currRules)
						newrules := ruleList{}
						for _, currRule := range currRules {
							// fmt.Println("currRule ", len(currRule), " = ", currRule)
							for _, prevrule := range tmpRulesToAdd {
								newrules = append(newrules, prevrule+currRule)
							}
						}
						// fmt.Println("candidate rule ", tmpruleid, " with ", len(newrules), " value  ", newrules)
						tmpRulesToAdd = newrules
					} else if ruleref == tmpruleid { // looop!
						newrules := ruleList{}
						// fmt.Println("currRule ", len(currRule), " = ", currRule)
						for _, prevrule := range tmpRulesToAdd {
							newrules = append(newrules, prevrule+".*")
						}
						// fmt.Println("candidate rule ", tmpruleid, " with ", len(newrules), " value  ", newrules)
						tmpRulesToAdd = newrules
					} else {
						addRule = false
						break
					}
					if !addRule {
						break
					}
				}
				rulesToAdd = append(rulesToAdd, tmpRulesToAdd...)
			}
			if addRule {
				// for _, ttt := range rulesToAdd {
				// 	fmt.Println("adding rule ", tmpruleid, " with value ", ttt)
				// }
				rules[tmpruleid] = rulesToAdd
			}
		}
	}

	// fmt.Printf("%s\n\n", strings.Repeat("-", 23))

	// for id, rules := range rules {
	// 	fmt.Println("rule ", id, " - len ", len(rules))
	// 	for _, rule := range rules {
	// 		fmt.Println("        - value ", rule)
	// 	}
	// }

	fmt.Printf("%s\n\nParsing values\n", strings.Repeat("-", 23))

	fmt.Printf("%s\n\nParsing values\n", strings.Repeat("-", 23))
	fmt.Println(rules[8])
	fmt.Println(rules[11])
	rulezero := rules[0]
	matches := 0
	for _, val := range ivalues {
		for _, rule := range rulezero {
			if strings.Contains(rule, ".*") {
				re := regexp.MustCompile(rule)
				if re.MatchString(val) {
					matches++
					break
				}
			} else if val == rule {
				matches++
				break
			}
		}
	}

	fmt.Printf("\n matches %d \n\n", matches)
}

// -----------------------------------------------------------------

func main() {
	rules, values := helper.FileAsDoubleStringList()
	parse(rules, values)
}
