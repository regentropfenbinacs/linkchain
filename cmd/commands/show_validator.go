package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/lianxiangcloud/linkchain/libs/ser"
	"github.com/lianxiangcloud/linkchain/types"
)

// ShowValidatorCmd adds capabilities for showing the validator info.
var ShowValidatorCmd = &cobra.Command{
	Use:   "show_validator",
	Short: "Show this node's validator info",
	Run:   showValidator,
}

func showValidator(cmd *cobra.Command, args []string) {
	privValidator := types.LoadOrGenFilePV(config.PrivValidatorFile())
	pubKeyJSONBytes, _ := ser.MarshalJSON(privValidator.GetPubKey())
	fmt.Println(string(pubKeyJSONBytes))
}
