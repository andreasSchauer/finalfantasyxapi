package seeding

import "fmt"

type BaseStat struct {
	ID       int32
	StatID   int32
	StatName string `json:"name"`
	Value    int32  `json:"value"`
}

func (bs BaseStat) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bs),
		bs.StatID,
		bs.Value,
	}
}

func (bs BaseStat) GetID() int32 {
	return bs.ID
}

func (b *BaseStat) SetID(id int32) {
	b.ID = id
}

func (bs BaseStat) GetName() string {
	return bs.StatName
}

func (bs BaseStat) GetVersion() *int32 {
	return nil
}

func (bs BaseStat) GetVal() int32 {
	return bs.Value
}

func (bs BaseStat) Error() string {
	return fmt.Sprintf("base stat %s, value: %d", bs.StatName, bs.Value)
}
