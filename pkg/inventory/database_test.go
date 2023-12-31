package inventory

import (
	"event-sourcing-go/pkg/uuid"
	"reflect"
	"testing"
)

func TestDatabase_UpsertInventoryItemRecord(t *testing.T) {
	db := NewDatabase()
	id := uuid.New()
	rec := &itemRecord{
		ID:           id,
		Name:         "hello",
		CurrentCount: 1,
		Version:      0,
	}

	db.UpsertInventoryItemRecord(rec)
	if db.records[id] != rec {
		t.Fatalf("whoops")
	}

	rec2 := &itemRecord{
		ID:           id,
		Name:         "hello again",
		CurrentCount: 2,
		Version:      1,
	}

	db.UpsertInventoryItemRecord(rec2)
	if db.records[id] != rec2 {
		t.Fatalf("whoops")
	}
}

func TestDatabase_GetInventoryItemRecord(t *testing.T) {
	db := NewDatabase()
	rec := &itemRecord{
		ID:           uuid.New(),
		Name:         "hello",
		CurrentCount: 1,
		Version:      0,
	}

	db.records[rec.ID] = rec
	if db.GetInventoryItemRecord(rec.ID) != rec {
		t.Fatalf("nope")
	}
}

func TestDatabase_GetInventoryItemRecords(t *testing.T) {
	db := NewDatabase()
	expectedRecs := []*itemRecord{
		{
			ID:           uuid.New(),
			Name:         "hello",
			CurrentCount: 1,
			Version:      0,
		},
		{
			ID:           uuid.New(),
			Name:         "hello again",
			CurrentCount: 2,
			Version:      1,
		},
	}

	for _, r := range expectedRecs {
		db.records[r.ID] = r
	}

	recs := db.GetInventoryItemRecords()
	if !reflect.DeepEqual(recs, expectedRecs) {
		t.Fatalf("nope")
	}
}

func TestDatabase_RemoveInventoryItemRecord(t *testing.T) {
	db := NewDatabase()
	rec := &itemRecord{
		ID:           uuid.New(),
		Name:         "hello",
		CurrentCount: 1,
		Version:      0,
	}

	db.records[rec.ID] = rec
	db.RemoveInventoryItemRecord(rec.ID)

	if len(db.records) != 0 {
		t.Fatalf("nope")
	}
}
