package client

import (
	"net/http"

	"github.com/textileio/lotus-client/api/apistruct"

	"github.com/textileio/lotus-client/api"
	"github.com/textileio/lotus-client/lib/jsonrpc"
)

// NewCommonRPC creates a new http jsonrpc client.
func NewCommonRPC(addr string, requestHeader http.Header) (api.Common, jsonrpc.ClientCloser, error) {
	var res apistruct.CommonStruct
	closer, err := jsonrpc.NewMergeClient(addr, "Filecoin",
		[]interface{}{
			&res.Internal,
		}, requestHeader)

	return &res, closer, err
}

// NewFullNodeRPC creates a new http jsonrpc client.
func NewFullNodeRPC(addr string, requestHeader http.Header) (api.FullNode, jsonrpc.ClientCloser, error) {
	var res apistruct.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(addr, "Filecoin",
		[]interface{}{
			&res.CommonStruct.Internal,
			&res.Internal,
		}, requestHeader)

	return &res, closer, err
}
