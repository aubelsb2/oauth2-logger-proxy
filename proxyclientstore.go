package oauth2_logger_proxy

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"log"
)

type proxyClientStore struct {
	clientStore oauth2.ClientStore
}

func ProxyClientStore(cs oauth2.ClientStore) oauth2.ClientStore {
	return &proxyClientStore{clientStore: cs}
}

func (pcs proxyClientStore) GetByID(ctx context.Context, id string) (ci oauth2.ClientInfo, err error) {
	log.Printf("GetByID(%#v)", id)
	ci, err = pcs.clientStore.GetByID(ctx, id)
	log.Printf("GetByID() = %v, %v", ci, err)
	return
}
