package playetcd

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func makeCli(t *testing.T) *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	require.NoError(t, err)
	return cli
}

func TestEtcd(t *testing.T) {
	cli := makeCli(t)
	defer cli.Close()

	putRes, err := cli.Put(cli.Ctx(), "foo", "bar1")
	require.NoError(t, err)
	require.NotNil(t, putRes)

	resp, err := cli.Get(cli.Ctx(), "foo")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "bar1", string(resp.Kvs[0].Value))
}

func TestEtcdMulti(t *testing.T) {
	cli := makeCli(t)
	defer cli.Close()

	for i := 0; i < 10; i++ {
		putRes, err := cli.Put(cli.Ctx(), fmt.Sprintf("multi%d", i), fmt.Sprintf("bar%d", i))
		require.NoError(t, err)
		require.NotNil(t, putRes)
	}

	resp, err := cli.Get(cli.Ctx(), "multi", clientv3.WithPrefix())
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Len(t, resp.Kvs, 10)
}

func TestWatch(t *testing.T) {
	cli := makeCli(t)
	defer cli.Close()

	watchChan := cli.Watch(cli.Ctx(), "foo", clientv3.WithPrefix())
	require.NotNil(t, watchChan)

}
