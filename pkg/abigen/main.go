package abigen

import (
	"io"
	"os"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/vistastaking/custom-staking-indexer/pkg/bind"
)

type ABIGenOptions struct {
	Abi    string
	Pkg    string
	Type   string
	Output string
}

func Abigen(options ABIGenOptions) error {
	// If the entire solidity code was specified, build and bind based on that
	var (
		abis    []string
		bins    []string
		types   []string
		sigs    []map[string]string
		libs    = make(map[string]string)
		aliases = make(map[string]string)
	)

	if options.Abi != "" {
		// Load up the ABI, optional bytecode and type name from the parameters
		var (
			abi []byte
			err error
		)

		input := options.Abi

		if input == "-" {
			abi, err = io.ReadAll(os.Stdin)
		} else {
			abi, err = os.ReadFile(input)
		}

		if err != nil {
			utils.Fatalf("Failed to read input ABI: %v", err)
		}

		var bin []byte
		bins = append(bins, string(bin))

		abis = append(abis, string(abi))
		types = append(types, options.Type)
	}

	// Generate the contract binding
	code, err := bind.Bind(types, abis, bins, sigs, options.Pkg, 0, libs, aliases)
	if err != nil {
		utils.Fatalf("Failed to generate ABI binding: %v", err)
	}

	if err := os.WriteFile(options.Output, []byte(code), 0600); err != nil {
		utils.Fatalf("Failed to write ABI binding: %v", err)
	}

	return nil
}
