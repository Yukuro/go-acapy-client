package acapy

import (
	"fmt"
	"strconv"
)

type CreateInvitationResponse struct {
	InvitationURL string `json:"invitation_url,omitempty"`
	ConnectionID  string `json:"connection_id,omitempty"`
	Invitation    struct {
		ImageURL        string   `json:"imageUrl,omitempty"`
		Label           string   `json:"label,omitempty"`
		ServiceEndpoint string   `json:"serviceEndpoint,omitempty"`
		RecipientKeys   []string `json:"recipientKeys,omitempty"`
		RoutingKeys     []string `json:"routingKeys,omitempty"`
		ID              string   `json:"@id,omitempty"`
		DID             string   `json:"did,omitempty"`
		Type            string   `json:"@type,omitempty"`
	} `json:"invitation,omitempty"`
}

func (c *Client) CreateInvitation(alias string, autoAccept bool, multiUse bool, public bool) (CreateInvitationResponse, error) {
	var createInvitationResponse CreateInvitationResponse
	err := c.post(c.AcapyURL+"/connections/create-invitation", map[string]string{
		"alias":       alias,
		"auto_accept": strconv.FormatBool(autoAccept),
		"multi_use":   strconv.FormatBool(multiUse),
		"public":      strconv.FormatBool(public),
	}, nil, &createInvitationResponse)
	return createInvitationResponse, err
}

type ReceiveInvitationResponse struct {
	InboundConnectionID string `json:"inbound_connection_id,omitempty"`
	InvitationKey       string `json:"invitation_key,omitempty"`
	MyDid               string `json:"my_did,omitempty"`
	TheirDid            string `json:"their_did,omitempty"`
	TheirRole           string `json:"their_role,omitempty"`
	RequestID           string `json:"request_id,omitempty"`
	State               string `json:"state,omitempty"`
	ConnectionID        string `json:"connection_id,omitempty"`
	Alias               string `json:"alias,omitempty"`
	InvitationMode      string `json:"invitation_mode,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	Accept              string `json:"accept,omitempty"`
	Initiator           string `json:"initiator,omitempty"`
	ErrorMsg            string `json:"error_msg,omitempty"`
	TheirLabel          string `json:"their_label,omitempty"`
	RoutingState        string `json:"routing_state,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
}

func (c *Client) ReceiveInvitation(invitation Invitation) (ReceiveInvitationResponse, error) {
	var receiveInvitationResponse ReceiveInvitationResponse

	err := c.post(c.AcapyURL+"/connections/receive-invitation", map[string]string{
		"alias":       invitation.Label,
		"auto_accept": strconv.FormatBool(false),
	}, invitation, &receiveInvitationResponse)
	return receiveInvitationResponse, err
}

type Invitation struct {
	ImageURL        string   `json:"imageUrl,omitempty"`
	Label           string   `json:"label,omitempty"`
	ServiceEndpoint string   `json:"serviceEndpoint,omitempty"`
	RecipientKeys   []string `json:"recipientKeys,omitempty"`
	RoutingKeys     []string `json:"routingKeys,omitempty"`
	ID              string   `json:"@id,omitempty"`
	DID             string   `json:"did,omitempty"`
}

func (c *Client) AcceptInvitation(connectionID string) (Connection, error) {
	var connection Connection
	err := c.post(fmt.Sprintf("%s/connections/%s/accept-invitation", c.AcapyURL, connectionID), nil, nil, &connection)
	return connection, err
}

func (c *Client) AcceptRequest(connectionID string) (Connection, error) {
	var connection Connection
	err := c.post(fmt.Sprintf("%s/connections/%s/accept-request", c.AcapyURL, connectionID), nil, nil, &connection)
	return connection, err
}

type Connection struct {
	Accept              string `json:"accept"`
	Alias               string `json:"alias"`
	ConnectionID        string `json:"connection_id"`
	CreatedAt           string `json:"created_at"`
	ErrorMsg            string `json:"error_msg"`
	InboundConnectionID string `json:"inbound_connection_id"`
	Initiator           string `json:"initiator"`
	InvitationKey       string `json:"invitation_key"`
	InvitationMode      string `json:"invitation_mode"`
	MyDid               string `json:"my_did"`
	RequestID           string `json:"request_id"`
	RoutingState        string `json:"routing_state"`
	State               string `json:"state"`
	TheirDid            string `json:"their_did"`
	TheirLabel          string `json:"their_label"`
	TheirRole           string `json:"their_role"`
	UpdatedAt           string `json:"updated_at"`
}

// QueryConnectionsParams model
//
// Parameters for querying connections
//
type QueryConnectionsParams struct {

	// Alias of connection invitation
	Alias string `json:"alias,omitempty"`

	// Initiator is Connection invitation initiator
	Initiator string `json:"initiator,omitempty"`

	// Invitation key
	InvitationKey string `json:"invitation_key,omitempty"`

	// MyDID is DID of the agent
	MyDID string `json:"my_did,omitempty"`

	// State of the connection invitation
	State string `json:"state"`

	// TheirDID is other party's DID
	TheirDID string `json:"their_did,omitempty"`

	// TheirRole is other party's role
	TheirRole string `json:"their_role,omitempty"`
}

func (c *Client) QueryConnections(params QueryConnectionsParams) ([]Connection, error) {
	var connections = struct {
		Result []Connection `json:"results"`
	}{}

	var queryParams = map[string]string{
		"alias":            params.Alias,
		"initiator":        params.Initiator,
		"invitation_key":   params.InvitationKey,
		"my_did":           params.MyDID,
		"connection_state": params.State,
		"their_did":        params.TheirDID,
		"their_role":       params.TheirRole,
	}
	err := c.get(c.AcapyURL+"/connections", queryParams, &connections)
	return connections.Result, err
}

func (c *Client) GetConnection(connectionID string) (Connection, error) {
	var connection Connection
	err := c.get(fmt.Sprintf("%s/connections/%s", c.AcapyURL, connectionID), nil, &connection)
	return connection, err
}

func (c *Client) RemoveConnection(connectionID string) error {
	return c.post(fmt.Sprintf("%s/connections/%s", c.AcapyURL, connectionID), nil, nil, nil)
}

type Thread struct {
	ThreadID string `json:"thread_id"`
}

func (c *Client) SendPing(connectionID string) (Thread, error) {
	ping := struct {
		Comment string `json:"comment"`
	}{
		Comment: "ping",
	}
	var thread Thread
	err := c.post(fmt.Sprintf("%s/connections/%s/send-ping", c.AcapyURL, connectionID), nil, ping, &thread)
	return thread, err
}
