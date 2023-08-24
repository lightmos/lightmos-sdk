package client

import (
	govclient "github.com/lightmos/lightmos-sdk/x/gov/client"
	"github.com/lightmos/lightmos-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
