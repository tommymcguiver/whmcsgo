package whmcsgo

import (
	"encoding/json"
)

/*
GetClientGroups

# Obtain an array of client groups

# WHMCS API docs

https://developers.whmcs.com/api-reference/getclientgroups/
*/
func (s *AccountsService) GetClientGroups(parms map[string]string) (*ClientGroup, *Response, error) {
	clientGroup := new(ClientGroup)
	resp, err := do(s.client, Params{parms: parms, u: "GetClientGroups"}, clientGroup)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &clientGroup)

	return clientGroup, resp, err
}
