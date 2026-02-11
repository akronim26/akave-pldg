package utils 

import (
	"testing"
	"context"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/common"

)

func GetLogs(t *testing.T , start int, end int) {
	ctx := context.Background()
	// Example parameters for GetLogs
	request_id := 1
	max_retry := 3
	address := common.HexToAddress("0xbFAbD47bF901ca1341D128DDD06463AA476E970B")
	topics := [][]string{
		[]string{"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"},
	}
	logs, err := rpc.GetLogs(ctx, request_id, max_retry, start, end, address, topics)
	if err != nil {
		t.Fatalf("GetLogs failed: %v", err)
	}	
	fmt.Printf("Logs: %v\n", logs)
}

func TestMakeRequest(t *testing.T) {
	//eth_chainId
	ctx := context.Background()
	response, err := rpc.MakeRequest(ctx, "eth_chainId", 1, 3, nil)
	if err != nil {
		t.Fatalf("MakeRequest failed: %v", err)
	}

	fmt.Printf("Response: %s\n", string(response.Result))
}

func TestGetLogs(t *testing.T) {
	cases := []struct {
		name  string
		start int
		end   int
	}{
		{"Test Case 1", 10000, 12000},
		{"Test Case 2", 10000, 50000},
	}

	for _, c := range cases {
		before := time.Now()
		t.Run(c.name, func(t *testing.T) {
			GetLogs(t, c.start, c.end)
		})
		after := time.Now()
		fmt.Printf("%s took %v\n", c.name, after.Sub(before))
	}
}