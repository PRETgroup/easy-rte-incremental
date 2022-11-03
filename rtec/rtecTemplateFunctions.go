package rtec

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/PRETgroup/easy-rte/rtedef"
)

//CECCTransition is used with getCECCTransitionCondition to return results to the template
type CECCTransition struct {
	IfCond    string
	AssEvents []string
}

//getCECCTransitionCondition returns the C "if" condition to use in state machine next state logic and associated events
// returns "full condition", "associated events"
func getCECCTransitionCondition(function rtedef.EnforcedFunction, trans string) CECCTransition {
	var events []string

	re1 := regexp.MustCompile("([<>=!]+)")          //for capturing operators
	re2 := regexp.MustCompile("([a-zA-Z0-9_<>=]+)") //for capturing variable and event names and operators
	isNum := regexp.MustCompile("^[0-9.]+$")

	retVal := trans

	//rename AND and OR
	retVal = strings.Replace(retVal, "AND", "&&", -1)
	retVal = strings.Replace(retVal, "OR", "||", -1)

	//re1: add whitespace around operators
	retVal = re1.ReplaceAllStringFunc(retVal, func(in string) string {
		if in != "!" {
			return " " + in + " "
		}
		return " !"
	})

	//re2: add "me->" where appropriate
	retVal = re2.ReplaceAllStringFunc(retVal, func(in string) string {
		if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ContainsAny(in, "!><=") || strings.ToLower(in) == "true" || strings.ToLower(in) == "false" {
			//no need to make changes, these aren't variables or events
			return in
		}

		if isNum.MatchString(in) {
			//no need to make changes, it is a numerical value of some sort
			return in
		}

		//check to see if it is input data
		if function.InputVars != nil {
			for _, Var := range function.InputVars {
				if Var.Name == in {
					return "inputs->" + in
				}
			}
		}

		//check to see if it is output data
		if function.OutputVars != nil {
			for _, Var := range function.OutputVars {
				if Var.Name == in {
					return "outputs->" + in
				}
			}
		}

		//check to see if it is a policy internal var
		for i := 0; i < len(function.Policies); i++ {
			for _, Var := range function.Policies[i].InternalVars {
				if Var.Name == in || Var.Name+"_i" == in {
					if Var.Constant {
						return "CONST_" + function.Policies[i].Name + "_" + in
					}
					return "me->" + in
				}
			}
		}

		//else, return it (no idea what else to do!) - it might be a function call or strange text constant
		return in
	})

	//tidy the whitespace
	retVal = strings.Replace(retVal, "  ", " ", -1)

	return CECCTransition{IfCond: retVal, AssEvents: events}
}

//getPolicyEnfInfo will get a PEnforcer for a given policy
func getPolicyEnfInfo(function rtedef.EnforcedFunction, policyIndex int) *rtedef.PEnforcer {
	enfPol, err := rtedef.MakePEnforcer(function.InterfaceList, function.Policies[policyIndex])
	if err != nil {
		return nil
	}
	//Uncomment these two lines to generate the intermediate enforcer JSON file
	//bytes, _ := json.MarshalIndent(enfPol, "", "\t")
	//ioutil.WriteFile(function.Name+".json", bytes, 0644)
	return enfPol
}

//getAllPolicyEnfInfo will get a PEnforcer for a given policy
func getAllPolicyEnfInfo(function rtedef.EnforcedFunction) []rtedef.PEnforcer {
	pols := make([]rtedef.PEnforcer, 0, len(function.Policies))
	for i := 0; i < len(function.Policies); i++ {
		enfPol, err := rtedef.MakePEnforcer(function.InterfaceList, function.Policies[i])
		if err != nil {
			return nil
		}
		pols = append(pols, *enfPol)
	}

	//Uncomment these two lines to generate the intermediate enforcer JSON file
	//bytes, _ := json.MarshalIndent(enfPol, "", "\t")
	//ioutil.WriteFile(function.Name+".json", bytes, 0644)
	return pols
}

func sub(a, b int) int {
	return a - b
}

func times(a, b int) int {
	return a * b
}

func twoToThePower(a int) int {
	return int(math.Pow(2, float64(a)))
}

func getBinaryCombinations(numberBits int) string {
	var combinations string = ""
	for i := 0; i < twoToThePower(numberBits); i++ {
		thisLine := "\n\t{" + fmt.Sprintf("%."+strconv.Itoa(numberBits)+"b", i) + "},"
		thisLine = strings.Replace(thisLine, "0", "0,", -1)
		thisLine = strings.Replace(thisLine, "1", "1,", -1)
		thisLine = strings.Replace(thisLine, "1,}", "1}", -1)
		thisLine = strings.Replace(thisLine, "0,}", "0}", -1)
		combinations = combinations + thisLine
	}

	return combinations
}

func getBADString(count int) string {
	var combinations string = "{BAD"
	for i := 1; i < count; i++ {
		combinations = combinations + ", BAD"
	}

	return combinations + "};"
}

func getAcceptableOptions(recoveryExpression string, outputInterface []rtedef.Variable, blockName string, interfaceDirection string) string {
	var splitString []string = strings.Split(recoveryExpression, "=")
	var recoverOutput string = strings.Replace(splitString[0], " ", "", -1)
	var recoverValue int = -3
	var recoverBit int
	recoverValue, _ = strconv.Atoi(strings.Replace(splitString[1], " ", "", -1))

	for i, v := range outputInterface {
		if v.Name == recoverOutput {
			recoverBit = len(outputInterface) - 1 - i
			fmt.Println("Recover " + v.Name + " (bit " + strconv.Itoa(recoverBit) + ") by setting it to" + splitString[1])
		}
	}

	var splitBinaryCombinations []string = strings.Split(strings.Replace(getBinaryCombinations(len(outputInterface)), "\t", "", -1), "\n")
	fmt.Print("\tAll options \t\t")
	fmt.Print(splitBinaryCombinations)
	fmt.Print("\n")
	var acceptableOutputs []string
	for i, possibleOutput := range splitBinaryCombinations {
		if i > 0 { // Skip first
			// fmt.Println(possibleOutput)
			var possibleOutputTweak string = strings.Replace(possibleOutput, "{", "", -1)
			possibleOutputTweak = strings.Replace(possibleOutputTweak, "},", "", -1)
			var values []string = strings.Split(possibleOutputTweak, ",")
			// fmt.Println(values)

			var possibleOutputRecoveryValue int
			possibleOutputRecoveryValue, _ = strconv.Atoi(values[len(values)-recoverBit-1])
			// fmt.Println("I think " + recoverOutput + " is in bit " + strconv.Itoa(recoverBit) + " and has a value of " + strconv.Itoa(possibleOutputRecoveryValue))
			if possibleOutputRecoveryValue == recoverValue {
				// Keep
				acceptableOutputs = append(acceptableOutputs, strings.Replace(possibleOutput, "},", "}", -1))
			} // else discard
		}
	}
	fmt.Print("\tAcceptable options \t")
	fmt.Print(acceptableOutputs)
	fmt.Print("\n")

	return "const uint8_t numAccept = " + strconv.Itoa(len(acceptableOutputs)) + ";\n\t\t\t\tconst " + interfaceDirection + "_" + blockName + "_t acceptableOptions[" + strconv.Itoa(len(acceptableOutputs)) + "] = {" + strings.Join(acceptableOutputs, ", ") + "};"
}
