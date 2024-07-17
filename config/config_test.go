package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	// Given: 환경 변수를 설정하고 기대 값을 정의합니다.
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))
	wantEnv := "dev"

	// When: 설정을 생성합니다.
	got, err := NewConfig()

	// Then: 결과를 검증합니다.
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but got %d", wantPort, got.Port)
	}
	if got.Env != wantEnv {
		t.Errorf("want %s, but got %s", wantEnv, got.Env)
	}
}
