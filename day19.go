package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type replacement struct {
	from string
	to   string
}

func main() {
	// Parse replacements.
	var replacements []replacement
	reReplacement := regexp.MustCompile(`^(\S+)\s=>\s(\S+)$`)
	scanner := bufio.NewScanner(strings.NewReader(inputReplacement))
	for scanner.Scan() {
		rc := reReplacement.FindStringSubmatch(scanner.Text())
		replacements = append(replacements, replacement{
			from: rc[1],
			to:   rc[2],
		})
	}

	// Iterate on all the different molecules that could be generated.
	results := make(map[string]struct{})
	for _, repl := range replacements {
		for idx := range input {
			if len(repl.from)+idx < len(input) &&
				repl.from == input[idx:idx+len(repl.from)] {

				results[strings.Join([]string{
					input[:idx],
					repl.to,
					input[idx+len(repl.from):]}, "")] = struct{}{}
			}
		}
	}
	fmt.Println("Total number of different molecules:", len(results))
}

// const input = `HOHOHO`
const input = `CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr`

/*const inputReplacement = `H => HO
H => OH
O => HH`*/
const inputReplacement = `Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg`
