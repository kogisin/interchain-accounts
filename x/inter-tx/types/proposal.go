package types

import (
	"fmt"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeRegisterInterchainAccount string = "RegisterInterchainAccount"
)

// Implements Proposal Interface
var _ gov.Content = &RegisterInterchainAccountProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeRegisterInterchainAccount)
	gov.RegisterProposalTypeCodec(&RegisterInterchainAccountProposal{}, "cosmos-sdk/RegisterInterchainAccountProposal")
}

// NewRegisterInterchainAccountProposal creates a new RegisterInterchainAccountProposal object.
func NewRegisterInterchainAccountProposal(
	title string,
	description string,
	owner string,
	connectionId string,
	counterpartyConnectionId string,
) *RegisterInterchainAccountProposal {
	return &RegisterInterchainAccountProposal{
		Title:                    title,
		Description:              description,
		Owner:                    owner,
		ConnectionId:             connectionId,
		CounterpartyConnectionId: counterpartyConnectionId,
	}
}

func (p *RegisterInterchainAccountProposal) GetTitle() string { return p.Title }

func (p *RegisterInterchainAccountProposal) GetDescription() string { return p.Description }

func (p *RegisterInterchainAccountProposal) ProposalRoute() string { return RouterKey }

func (p *RegisterInterchainAccountProposal) ProposalType() string {
	return ProposalTypeRegisterInterchainAccount
}

func (p *RegisterInterchainAccountProposal) ValidateBasic() error {
	// TODO: only whitelisted validators can submit proposal
	return gov.ValidateAbstract(p)
}

func (p RegisterInterchainAccountProposal) String() string {
	return fmt.Sprintf(`Register Interchain Account Proposal:
  Title:              		%s
  Description:        		%s
  Owner:         	  		%s
  ConnectionId: 	  		%s
  CounterpartyConnectionId: %s
`, p.Title, p.Description, p.Owner, p.ConnectionId, p.CounterpartyConnectionId)
}
