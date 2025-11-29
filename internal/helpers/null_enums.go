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


func NullTopmenuType(s *string) database.NullTopmenuType {
    if s == nil {
        return database.NullTopmenuType{}
    }
    
    return database.NullTopmenuType{
        TopmenuType: database.TopmenuType(*s),
        Valid: true,
    }
}

func ConvertNullTopmenuType(ne database.NullTopmenuType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.TopmenuType)
    return &val
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


func NullItemUsability(s *string) database.NullItemUsability {
    if s == nil {
        return database.NullItemUsability{}
    }
    
    return database.NullItemUsability{
        ItemUsability: database.ItemUsability(*s),
        Valid: true,
    }
}

func ConvertNullItemUsability(ne database.NullItemUsability) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.ItemUsability)
    return &val
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


