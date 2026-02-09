package helpers

// these interfaces serve as bridge between the seeding resources and the api resources and are used to convert the former into the latter

type HasID interface {
	GetID() int32
}

type IsNamed interface {
	HasID
	GetResParamsNamed() ResParamsNamed
}

type IsUnnamed interface {
	HasID
	GetResParamsUnnamed() ResParamsUnnamed
}

type IsLocationBased interface {
	HasID
	GetResParamsLocation() ResParamsLocation
}

type IsTypeBased interface {
	HasID
	GetResParamsTyped() ResParamsTyped
}

type ResParamsNamed struct {
	ID            int32
	Name          string
	Version       *int32
	Specification *string
}

type ResParamsUnnamed struct {
	ID	int32
}

type ResParamsLocation struct {
	AreaID			int32
	Location		string
	Sublocation		string
	Area			string
	Version			*int32
	Specification 	*string
}

type ResParamsTyped struct {
	ID            	int32
	Name          	string
	Description		string
}