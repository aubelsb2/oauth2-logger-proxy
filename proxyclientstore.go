package oauth2_logger_proxy

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"log"
)

type ProxyClientStore struct {
	clientStore oauth2.ClientStore
}

func (pcs ProxyClientStore) GetByID(ctx context.Context, id string) (ci oauth2.ClientInfo, err error) {
	log.Printf("GetByID(%#v)", id)
	ci, err = pcs.clientStore.GetByID(ctx, id)
	log.Printf("GetByID() = %v, %v", ci, err)
	return
}
