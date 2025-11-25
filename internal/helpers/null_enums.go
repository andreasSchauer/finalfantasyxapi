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


func NullCreationsUnlockedCategory(s *string) database.NullCreationsUnlockedCategory {
    if s == nil {
        return database.NullCreationsUnlockedCategory{}
    }
    
    return database.NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: database.CreationsUnlockedCategory(*s),
        Valid: true,
    }
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


func NullEquipType(s *string) database.NullEquipType {
    if s == nil {
        return database.NullEquipType{}
    }
    
    return database.NullEquipType{
        EquipType: database.EquipType(*s),
        Valid: true,
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


func NullCounterType(s *string) database.NullCounterType {
    if s == nil {
        return database.NullCounterType{}
    }
    
    return database.NullCounterType{
        CounterType: database.CounterType(*s),
        Valid: true,
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


func NullMaCreationSpecies(s *string) database.NullMaCreationSpecies {
    if s == nil {
        return database.NullMaCreationSpecies{}
    }
    
    return database.NullMaCreationSpecies{
        MaCreationSpecies: database.MaCreationSpecies(*s),
        Valid: true,
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


func NullItemUsability(s *string) database.NullItemUsability {
    if s == nil {
        return database.NullItemUsability{}
    }
    
    return database.NullItemUsability{
        ItemUsability: database.ItemUsability(*s),
        Valid: true,
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


func NullBgReplacementType(s *string) database.NullBgReplacementType {
    if s == nil {
        return database.NullBgReplacementType{}
    }
    
    return database.NullBgReplacementType{
        BgReplacementType: database.BgReplacementType(*s),
        Valid: true,
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


func NullCriticalType(s *string) database.NullCriticalType {
    if s == nil {
        return database.NullCriticalType{}
    }
    
    return database.NullCriticalType{
        CriticalType: database.CriticalType(*s),
        Valid: true,
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


