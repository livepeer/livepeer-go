// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type IpfsFileInfo struct {
	// CID of the file on IPFS
	Cid string `json:"cid"`
	// URL with IPFS scheme for the file
	URL *string `json:"url,omitempty"`
	// URL to access file via HTTP through an IPFS gateway
	GatewayURL *string `json:"gatewayUrl,omitempty"`
}

func (o *IpfsFileInfo) GetCid() string {
	if o == nil {
		return ""
	}
	return o.Cid
}

func (o *IpfsFileInfo) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *IpfsFileInfo) GetGatewayURL() *string {
	if o == nil {
		return nil
	}
	return o.GatewayURL
}
