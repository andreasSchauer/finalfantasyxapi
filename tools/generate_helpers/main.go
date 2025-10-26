package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    nullEnumTypes := []string{
        "NullifyArmored",
        "AccuracySource",
        "CreationsUnlockedCategory",
        "TopmenuType",
        "EquipType",
        "AaActivationCondition",
        "CounterType",
        "MaCreationArea",
        "MaCreationSpecies",
        "TargetType",
        "ItemUsability",
        "MusicUseCase",
        "BgReplacementType",
    }

    filePath := "./internal/seeding/null_enums.go"
    
    var output strings.Builder
    output.WriteString("package seeding\n\nimport \"github.com/andreasSchauer/finalfantasyxapi/internal/database\"\n\n\n")

    for _, enumType := range nullEnumTypes {
        funcName := "null" + enumType
        output.WriteString(fmt.Sprintf(`func %s(s *string) database.Null%s {
    if s == nil {
        return database.Null%s{}
    }
    
    return database.Null%s{
        %s: database.%s(*s),
        Valid: true,
    }
}


`, funcName, enumType, enumType, enumType, enumType, enumType))
    }

    os.WriteFile(filePath, []byte(output.String()), 0644)
}