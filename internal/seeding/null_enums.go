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


