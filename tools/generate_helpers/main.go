package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    nullEnumTypes := []string{
        "NullifyArmored",
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
        "SpecialActionType",
        "CriticalType",
        "BreakDmgLmtType",
        "Composer",
        "Arranger",
    }

    filePath := "./internal/helpers/null_enums.go"
    
    var output strings.Builder
    output.WriteString("package helpers\n\nimport \"github.com/andreasSchauer/finalfantasyxapi/internal/database\"\n\n\n")

    for _, enumType := range nullEnumTypes {
        nullEnumFuncName := "Null" + enumType
        convertFuncName := "Convert" + nullEnumFuncName
        output.WriteString(fmt.Sprintf(`func %s(s *string) database.Null%s {
    if s == nil {
        return database.Null%s{}
    }
    
    return database.Null%s{
        %s: database.%s(*s),
        Valid: true,
    }
}

func %s(ne database.Null%s) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.%s)
    return &val
}


`, nullEnumFuncName, enumType, enumType, enumType, enumType, enumType, convertFuncName, enumType, enumType))
    }

    os.WriteFile(filePath, []byte(output.String()), 0644)
}