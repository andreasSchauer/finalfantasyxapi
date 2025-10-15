package seeding

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


func nullNullifyArmored(s *string) database.NullNullifyArmored {
    if s == nil {
        return database.NullNullifyArmored{}
    }
    
    return database.NullNullifyArmored{
        NullifyArmored: database.NullifyArmored(*s),
        Valid: true,
    }
}


func nullAccuracySource(s *string) database.NullAccuracySource {
    if s == nil {
        return database.NullAccuracySource{}
    }
    
    return database.NullAccuracySource{
        AccuracySource: database.AccuracySource(*s),
        Valid: true,
    }
}


func nullCreationsUnlockedCategory(s *string) database.NullCreationsUnlockedCategory {
    if s == nil {
        return database.NullCreationsUnlockedCategory{}
    }
    
    return database.NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: database.CreationsUnlockedCategory(*s),
        Valid: true,
    }
}


func nullTopmenuType(s *string) database.NullTopmenuType {
    if s == nil {
        return database.NullTopmenuType{}
    }
    
    return database.NullTopmenuType{
        TopmenuType: database.TopmenuType(*s),
        Valid: true,
    }
}


func nullEquipType(s *string) database.NullEquipType {
    if s == nil {
        return database.NullEquipType{}
    }
    
    return database.NullEquipType{
        EquipType: database.EquipType(*s),
        Valid: true,
    }
}


func nullAaActivationCondition(s *string) database.NullAaActivationCondition {
    if s == nil {
        return database.NullAaActivationCondition{}
    }
    
    return database.NullAaActivationCondition{
        AaActivationCondition: database.AaActivationCondition(*s),
        Valid: true,
    }
}


func nullCounterType(s *string) database.NullCounterType {
    if s == nil {
        return database.NullCounterType{}
    }
    
    return database.NullCounterType{
        CounterType: database.CounterType(*s),
        Valid: true,
    }
}



func nullMaCreationArea(s *string) database.NullMaCreationArea {
    if s == nil {
        return database.NullMaCreationArea{}
    }
    
    return database.NullMaCreationArea{
        MaCreationArea: database.MaCreationArea(*s),
        Valid: true,
    }
}


func nullMaCreationSpecies(s *string) database.NullMaCreationSpecies {
    if s == nil {
        return database.NullMaCreationSpecies{}
    }
    
    return database.NullMaCreationSpecies{
        MaCreationSpecies: database.MaCreationSpecies(*s),
        Valid: true,
    }
}


func nullTargetType(s *string) database.NullTargetType {
    if s == nil {
        return database.NullTargetType{}
    }
    
    return database.NullTargetType{
        TargetType: database.TargetType(*s),
        Valid: true,
    }
}


func nullItemUsability(s *string) database.NullItemUsability {
    if s == nil {
        return database.NullItemUsability{}
    }
    
    return database.NullItemUsability{
        ItemUsability: database.ItemUsability(*s),
        Valid: true,
    }
}


func nullMusicUseCase(s *string) database.NullMusicUseCase {
    if s == nil {
        return database.NullMusicUseCase{}
    }
    
    return database.NullMusicUseCase{
        MusicUseCase: database.MusicUseCase(*s),
        Valid: true,
    }
}


