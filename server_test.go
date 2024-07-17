package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("포트 수신에 실패했습니다: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	eg.Go(func() error {
		s := NewServer(l, mux)
		return s.Run(ctx)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// 어떤 포트 번호로 리슨하고 있는지 확인
	t.Logf("요청을 시도합니다: %q", url)
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("요청 실패: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("응답 본문을 읽는 데 실패했습니다: %v", err)
	}
	// HTTP 서버의 반환 값을 검증합니다
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("원하는 값은 %q, 그러나 받은 값은 %q", want, got)
	}
	// run 함수에 종료 통지를 보냅니다.
	cancel()
	// run 함수의 반환 값을 검증합니다
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
