package playetcd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcd(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	require.NoError(t, err)
	defer cli.Close()

	putRes, err := cli.Put(cli.Ctx(), "foo", "bar")
	require.NoError(t, err)
	require.NotNil(t, putRes)

	resp, err := cli.Get(cli.Ctx(), "foo")
	require.NoError(t, err)
	require.NotNil(t, resp)
}
