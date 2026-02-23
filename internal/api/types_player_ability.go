package api


import (

)

type PlayerAbility struct {
	ID						int32
	Name					string
	Version					*int32
	Specification			*string
	Description				*string
	Effect					string
	Rank					*int32
	AppearsInHelpBar		bool
	CanCopycat				bool
	CanUseOutsideBattle		bool
	MpCost					*int32
	Category				string
	AeonLearnItem			*ItemAmount
	RelatedStats			[]NamedAPIResource
	LearnedBy				[]NamedAPIResource
	StandardGridCharacter	*NamedAPIResource
	ExpertGridCharacter		*NamedAPIResource
	Topmenu					*NamedAPIResource
	Submenu					*NamedAPIResource
	OpenSubmenu				*NamedAPIResource
	Cursor					*NamedAPIResource
}


type BattleInteraction struct {
	Target				NamedAPIResource
	BasedOnPhysAttack	bool
	Range				*int32
	ShatterRate			*int32
	Darkable			bool
	Silenceable			bool
	Reflectable			bool
	HitAmount			int32
	SpecialAction		*string
}