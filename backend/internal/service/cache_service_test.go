package service

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
)

func TestCacheService_Get(t *testing.T) {
	db, mock := redismock.NewClientMock()
	svc := NewCacheService(db)
	ctx := context.Background()

	mock.ExpectGet("test_key").SetVal("test_value")

	val, err := svc.Get(ctx, "test_key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if val != "test_value" {
		t.Errorf("expected test_value, got %s", val)
	}
}

func TestCacheService_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()
	svc := NewCacheService(db)
	ctx := context.Background()

	mock.ExpectSet("test_key", "test_value", 30*time.Minute).SetVal("OK")

	err := svc.Set(ctx, "test_key", "test_value", 30*time.Minute)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet expectations: %v", err)
	}
}
