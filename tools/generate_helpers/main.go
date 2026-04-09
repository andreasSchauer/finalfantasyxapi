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
        "EquipType",
        "AaActivationCondition",
        "CounterType",
        "MaCreationArea",
        "MaCreationCategory",
        "MaCreationSpecies",
        "TargetType",
        "MusicUseCase",
        "BgReplacementType",
        "SpecialActionType",
        "CriticalType",
        "BreakDmgLmtType",
        "Composer",
        "Arranger",
        "ShopType",
        "AvailabilityType",
        "NodeState",
    }

    filePath := "./internal/helpers/conv_null_enum.go"
    
    var output strings.Builder
    output.WriteString("package helpers\n\nimport \"github.com/andreasSchauer/finalfantasyxapi/internal/database\"\n\n\n")

    for _, enumType := range nullEnumTypes {
        nullEnumFuncName := "Null" + enumType
        convertFuncName := "Convert" + nullEnumFuncName
        getNullEnumFuncName := "Get" + nullEnumFuncName
        fmt.Fprintf(&output, `func %s(s *string) database.Null%s {
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

func %s(e *database.%s) database.Null%s {
    if e == nil {
        return database.Null%s{}
    }

    return database.Null%s{
        %s: *e,
        Valid:  true,
    }
}


`, nullEnumFuncName, enumType, enumType, enumType, enumType, enumType, convertFuncName, enumType, enumType, getNullEnumFuncName, enumType, enumType, enumType, enumType, enumType)
    }

    os.WriteFile(filePath, []byte(output.String()), 0644)
}