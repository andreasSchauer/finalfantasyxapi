package helpers

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


func NullNullifyArmored(s *string) database.NullNullifyArmored {
    if s == nil {
        return database.NullNullifyArmored{}
    }
    
    return database.NullNullifyArmored{
        NullifyArmored: database.NullifyArmored(*s),
        Valid: true,
    }
}

func ConvertNullNullifyArmored(ne database.NullNullifyArmored) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.NullifyArmored)
    return &val
}

func GetNullNullifyArmored(e *database.NullifyArmored) database.NullNullifyArmored {
    if e == nil {
        return database.NullNullifyArmored{}
    }

    return database.NullNullifyArmored{
        NullifyArmored: *e,
        Valid:  true,
    }
}


func NullCreationsUnlockedCategory(s *string) database.NullCreationsUnlockedCategory {
    if s == nil {
        return database.NullCreationsUnlockedCategory{}
    }
    
    return database.NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: database.CreationsUnlockedCategory(*s),
        Valid: true,
    }
}

func ConvertNullCreationsUnlockedCategory(ne database.NullCreationsUnlockedCategory) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CreationsUnlockedCategory)
    return &val
}

func GetNullCreationsUnlockedCategory(e *database.CreationsUnlockedCategory) database.NullCreationsUnlockedCategory {
    if e == nil {
        return database.NullCreationsUnlockedCategory{}
    }

    return database.NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: *e,
        Valid:  true,
    }
}


func NullEquipType(s *string) database.NullEquipType {
    if s == nil {
        return database.NullEquipType{}
    }
    
    return database.NullEquipType{
        EquipType: database.EquipType(*s),
        Valid: true,
    }
}

func ConvertNullEquipType(ne database.NullEquipType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.EquipType)
    return &val
}

func GetNullEquipType(e *database.EquipType) database.NullEquipType {
    if e == nil {
        return database.NullEquipType{}
    }

    return database.NullEquipType{
        EquipType: *e,
        Valid:  true,
    }
}


func NullAaActivationCondition(s *string) database.NullAaActivationCondition {
    if s == nil {
        return database.NullAaActivationCondition{}
    }
    
    return database.NullAaActivationCondition{
        AaActivationCondition: database.AaActivationCondition(*s),
        Valid: true,
    }
}

func ConvertNullAaActivationCondition(ne database.NullAaActivationCondition) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.AaActivationCondition)
    return &val
}

func GetNullAaActivationCondition(e *database.AaActivationCondition) database.NullAaActivationCondition {
    if e == nil {
        return database.NullAaActivationCondition{}
    }

    return database.NullAaActivationCondition{
        AaActivationCondition: *e,
        Valid:  true,
    }
}


func NullCounterType(s *string) database.NullCounterType {
    if s == nil {
        return database.NullCounterType{}
    }
    
    return database.NullCounterType{
        CounterType: database.CounterType(*s),
        Valid: true,
    }
}

func ConvertNullCounterType(ne database.NullCounterType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CounterType)
    return &val
}

func GetNullCounterType(e *database.CounterType) database.NullCounterType {
    if e == nil {
        return database.NullCounterType{}
    }

    return database.NullCounterType{
        CounterType: *e,
        Valid:  true,
    }
}


func NullMaCreationArea(s *string) database.NullMaCreationArea {
    if s == nil {
        return database.NullMaCreationArea{}
    }
    
    return database.NullMaCreationArea{
        MaCreationArea: database.MaCreationArea(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationArea(ne database.NullMaCreationArea) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationArea)
    return &val
}

func GetNullMaCreationArea(e *database.MaCreationArea) database.NullMaCreationArea {
    if e == nil {
        return database.NullMaCreationArea{}
    }

    return database.NullMaCreationArea{
        MaCreationArea: *e,
        Valid:  true,
    }
}


func NullMaCreationCategory(s *string) database.NullMaCreationCategory {
    if s == nil {
        return database.NullMaCreationCategory{}
    }
    
    return database.NullMaCreationCategory{
        MaCreationCategory: database.MaCreationCategory(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationCategory(ne database.NullMaCreationCategory) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationCategory)
    return &val
}

func GetNullMaCreationCategory(e *database.MaCreationCategory) database.NullMaCreationCategory {
    if e == nil {
        return database.NullMaCreationCategory{}
    }

    return database.NullMaCreationCategory{
        MaCreationCategory: *e,
        Valid:  true,
    }
}


func NullMaCreationSpecies(s *string) database.NullMaCreationSpecies {
    if s == nil {
        return database.NullMaCreationSpecies{}
    }
    
    return database.NullMaCreationSpecies{
        MaCreationSpecies: database.MaCreationSpecies(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationSpecies(ne database.NullMaCreationSpecies) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationSpecies)
    return &val
}

func GetNullMaCreationSpecies(e *database.MaCreationSpecies) database.NullMaCreationSpecies {
    if e == nil {
        return database.NullMaCreationSpecies{}
    }

    return database.NullMaCreationSpecies{
        MaCreationSpecies: *e,
        Valid:  true,
    }
}


func NullTargetType(s *string) database.NullTargetType {
    if s == nil {
        return database.NullTargetType{}
    }
    
    return database.NullTargetType{
        TargetType: database.TargetType(*s),
        Valid: true,
    }
}

func ConvertNullTargetType(ne database.NullTargetType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.TargetType)
    return &val
}

func GetNullTargetType(e *database.TargetType) database.NullTargetType {
    if e == nil {
        return database.NullTargetType{}
    }

    return database.NullTargetType{
        TargetType: *e,
        Valid:  true,
    }
}


func NullMusicUseCase(s *string) database.NullMusicUseCase {
    if s == nil {
        return database.NullMusicUseCase{}
    }
    
    return database.NullMusicUseCase{
        MusicUseCase: database.MusicUseCase(*s),
        Valid: true,
    }
}

func ConvertNullMusicUseCase(ne database.NullMusicUseCase) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MusicUseCase)
    return &val
}

func GetNullMusicUseCase(e *database.MusicUseCase) database.NullMusicUseCase {
    if e == nil {
        return database.NullMusicUseCase{}
    }

    return database.NullMusicUseCase{
        MusicUseCase: *e,
        Valid:  true,
    }
}


func NullBgReplacementType(s *string) database.NullBgReplacementType {
    if s == nil {
        return database.NullBgReplacementType{}
    }
    
    return database.NullBgReplacementType{
        BgReplacementType: database.BgReplacementType(*s),
        Valid: true,
    }
}

func ConvertNullBgReplacementType(ne database.NullBgReplacementType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.BgReplacementType)
    return &val
}

func GetNullBgReplacementType(e *database.BgReplacementType) database.NullBgReplacementType {
    if e == nil {
        return database.NullBgReplacementType{}
    }

    return database.NullBgReplacementType{
        BgReplacementType: *e,
        Valid:  true,
    }
}


func NullSpecialActionType(s *string) database.NullSpecialActionType {
    if s == nil {
        return database.NullSpecialActionType{}
    }
    
    return database.NullSpecialActionType{
        SpecialActionType: database.SpecialActionType(*s),
        Valid: true,
    }
}

func ConvertNullSpecialActionType(ne database.NullSpecialActionType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.SpecialActionType)
    return &val
}

func GetNullSpecialActionType(e *database.SpecialActionType) database.NullSpecialActionType {
    if e == nil {
        return database.NullSpecialActionType{}
    }

    return database.NullSpecialActionType{
        SpecialActionType: *e,
        Valid:  true,
    }
}


func NullCriticalType(s *string) database.NullCriticalType {
    if s == nil {
        return database.NullCriticalType{}
    }
    
    return database.NullCriticalType{
        CriticalType: database.CriticalType(*s),
        Valid: true,
    }
}

func ConvertNullCriticalType(ne database.NullCriticalType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CriticalType)
    return &val
}

func GetNullCriticalType(e *database.CriticalType) database.NullCriticalType {
    if e == nil {
        return database.NullCriticalType{}
    }

    return database.NullCriticalType{
        CriticalType: *e,
        Valid:  true,
    }
}


func NullBreakDmgLmtType(s *string) database.NullBreakDmgLmtType {
    if s == nil {
        return database.NullBreakDmgLmtType{}
    }
    
    return database.NullBreakDmgLmtType{
        BreakDmgLmtType: database.BreakDmgLmtType(*s),
        Valid: true,
    }
}

func ConvertNullBreakDmgLmtType(ne database.NullBreakDmgLmtType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.BreakDmgLmtType)
    return &val
}

func GetNullBreakDmgLmtType(e *database.BreakDmgLmtType) database.NullBreakDmgLmtType {
    if e == nil {
        return database.NullBreakDmgLmtType{}
    }

    return database.NullBreakDmgLmtType{
        BreakDmgLmtType: *e,
        Valid:  true,
    }
}


func NullComposer(s *string) database.NullComposer {
    if s == nil {
        return database.NullComposer{}
    }
    
    return database.NullComposer{
        Composer: database.Composer(*s),
        Valid: true,
    }
}

func ConvertNullComposer(ne database.NullComposer) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.Composer)
    return &val
}

func GetNullComposer(e *database.Composer) database.NullComposer {
    if e == nil {
        return database.NullComposer{}
    }

    return database.NullComposer{
        Composer: *e,
        Valid:  true,
    }
}


func NullArranger(s *string) database.NullArranger {
    if s == nil {
        return database.NullArranger{}
    }
    
    return database.NullArranger{
        Arranger: database.Arranger(*s),
        Valid: true,
    }
}

func ConvertNullArranger(ne database.NullArranger) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.Arranger)
    return &val
}

func GetNullArranger(e *database.Arranger) database.NullArranger {
    if e == nil {
        return database.NullArranger{}
    }

    return database.NullArranger{
        Arranger: *e,
        Valid:  true,
    }
}


func NullShopType(s *string) database.NullShopType {
    if s == nil {
        return database.NullShopType{}
    }
    
    return database.NullShopType{
        ShopType: database.ShopType(*s),
        Valid: true,
    }
}

func ConvertNullShopType(ne database.NullShopType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.ShopType)
    return &val
}

func GetNullShopType(e *database.ShopType) database.NullShopType {
    if e == nil {
        return database.NullShopType{}
    }

    return database.NullShopType{
        ShopType: *e,
        Valid:  true,
    }
}


func NullAvailabilityType(s *string) database.NullAvailabilityType {
    if s == nil {
        return database.NullAvailabilityType{}
    }
    
    return database.NullAvailabilityType{
        AvailabilityType: database.AvailabilityType(*s),
        Valid: true,
    }
}

func ConvertNullAvailabilityType(ne database.NullAvailabilityType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.AvailabilityType)
    return &val
}

func GetNullAvailabilityType(e *database.AvailabilityType) database.NullAvailabilityType {
    if e == nil {
        return database.NullAvailabilityType{}
    }

    return database.NullAvailabilityType{
        AvailabilityType: *e,
        Valid:  true,
    }
}


