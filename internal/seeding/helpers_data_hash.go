package seeding

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func generateDataHash(h Hashable) string {
	return CombineFields(h.ToHashFields())
}

func generateJunctionHash(j Junction, desc string) string {
	return CombineFields(j.ToHashFieldsJ(desc))
}

func (l *Lookup) GetHashID(h Hashable) (int32, error) {
	id, ok := l.Hashes[generateDataHash(h)]
	if !ok {
		return 0, fmt.Errorf("no data hash available for %s", CombineFields(h.ToHashFields()))
	}

	return id, nil
}

func CombineFields(fields []any) string {
	var builder strings.Builder

	for i, field := range fields {
		if i > 0 {
			builder.WriteString("|")
		}

		if field == nil {
			builder.WriteString("NULL")
		} else {
			fmt.Fprint(&builder, field)
		}
	}

	return builder.String()
}

type hashSave struct {
	ID		int32	`json:"id"`
	Hash 	string	`json:"hash"`
}

func (h hashSave) GetID() int32 {
	return h.ID
}

func (h hashSave) ToKeyFields() []any {
	return []any{
		h.ID,
		h.Hash,
	}
}

func (h hashSave) Error() string {
	return fmt.Sprintf("hash save with id: %d, hash: %s", h.ID, h.Hash)
}

func (l *Lookup) saveHashMap() error {
	fileName := "hashes.json"
	lookupDir, err := h.GetAbsoluteFilepath("data_lookups")
	if err != nil {
		return err
	}
	destPath := filepath.Join(lookupDir, fileName)

	var s []hashSave

	for hash, id := range l.Hashes {
		save := hashSave{
			ID: 	id,
			Hash: 	hash,
		}

		s = append(s, save)
	}

	slices.SortStableFunc(s, func(a, b hashSave) int {
		if a.Hash < b.Hash {
			return -1
		}

		if a.Hash > b.Hash {
			return 1
		}

		return 0
	})

	jsonBytes, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(destPath, jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) assignHashes() error {
	path := filepath.Join("data_lookups", "hashes.json")

	var dataSlice []hashSave
	err := loadJsonFile(path, &dataSlice)
	if err != nil {
		return err
	}

	for _, save := range dataSlice {
		l.Hashes[save.Hash] = save.ID
	}

	return nil
}