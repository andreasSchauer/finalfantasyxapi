package seeding

// these interfaces serve as bridge between the seeding resources and the api resources and are used to convert the former into the latter

type IsNamed interface {
	Lookupable
	GetResParamsNamed() ResParamsNamed
}

type ResParamsNamed struct {
	ID            int32
	Name          string
	Version       *int32
	Specification *string
}

type IsUnnamed interface {
	Lookupable
	GetResParamsUnnamed() ResParamsUnnamed
}

type ResParamsUnnamed struct {
	ID int32
}

type IsLocationBased interface {
	Lookupable
	GetResParamsLocation() ResParamsLocation
}

type ResParamsLocation struct {
	AreaID        int32
	Location      string
	Sublocation   string
	Area          string
	Version       *int32
	Specification *string
}

type IsTyped interface {
	Lookupable
	GetResParamsTyped() ResParamsTyped
}

type ResParamsTyped struct {
	ID            int32
	Name          string
	Version       *int32
	Specification *string
	Type          string
}

type IsQuest interface {
	Lookupable
	GetResParamsQuest() ResParamsQuest
}

type ResParamsQuest struct {
	ID        int32
	Sidequest *string
	Subquest  *string
	Type      string
}

type IsEnum interface {
	Lookupable
	GetResParamsEnum() ResParamsEnum
}

type ResParamsEnum struct {
	ID          int32
	Name        string
	Description string
}
