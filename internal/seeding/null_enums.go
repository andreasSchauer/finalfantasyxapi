package seeding

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


func nullAccuracySource(s *string) database.NullAccuracySource {
    if s == nil {
        return database.NullAccuracySource{}
    }
    
    return database.NullAccuracySource{
        AccuracySource: database.AccuracySource(*s),
        Valid: true,
    }
}


func nullAeonCategory(s *string) database.NullAeonCategory {
    if s == nil {
        return database.NullAeonCategory{}
    }
    
    return database.NullAeonCategory{
        AeonCategory: database.AeonCategory(*s),
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


func nullNullifyArmored(s *string) database.NullNullifyArmored {
    if s == nil {
        return database.NullNullifyArmored{}
    }
    
    return database.NullNullifyArmored{
        NullifyArmored: database.NullifyArmored(*s),
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


func nullRecoveryType(s *string) database.NullRecoveryType {
    if s == nil {
        return database.NullRecoveryType{}
    }
    
    return database.NullRecoveryType{
        RecoveryType: database.RecoveryType(*s),
        Valid: true,
    }
}


func nullElementType(s *string) database.NullElementType {
    if s == nil {
        return database.NullElementType{}
    }
    
    return database.NullElementType{
        ElementType: database.ElementType(*s),
        Valid: true,
    }
}


func nullParameter(s *string) database.NullParameter {
    if s == nil {
        return database.NullParameter{}
    }
    
    return database.NullParameter{
        Parameter: database.Parameter(*s),
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


func nullSubmenuType(s *string) database.NullSubmenuType {
    if s == nil {
        return database.NullSubmenuType{}
    }
    
    return database.NullSubmenuType{
        SubmenuType: database.SubmenuType(*s),
        Valid: true,
    }
}


