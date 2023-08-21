package ec2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_ChangeStatus(t *testing.T) {
	var err error
	ctx := context.TODO()
	client, err := New(ctx, "ap-southeast-1")
	assert.NoError(t, err)
	err = client.ChangeStatus(ctx, 1, "phuong", "test")
	assert.NoError(t, err)

}
