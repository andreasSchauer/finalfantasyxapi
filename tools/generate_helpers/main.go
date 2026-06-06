package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	nullEnumTypes := []string{
		"Arranger",
		"AvailabilityType",
		"BgReplacementType",
		"BreakDmgLmtType",
		"Composer",
		"CounterType",
		"CreationsUnlockedCategory",
		"CriticalType",
		"MaCreationArea",
		"MaCreationCategory",
		"MaCreationSpecies",
		"MusicUseCase",
		"NodeState",
		"NullifyArmored",
		"ShopType",
		"SpecialActionType",
		"TargetType",
	}

	filePath := "./internal/database/conv_null_enum.go"

	var output strings.Builder
	output.WriteString("package database\n\n")

	for _, enumType := range nullEnumTypes {
		nullEnumFuncName := "ToNull" + enumType
		convertFuncName := "ConvertNull" + enumType
		getNullEnumFuncName := "GetNull" + enumType
		fmt.Fprintf(&output, `func %s(s *string) Null%s {
    if s == nil {
        return Null%s{}
    }
    
    return Null%s{
        %s: %s(*s),
        Valid: true,
    }
}

func %s(ne Null%s) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.%s)
    return &val
}

func %s(e *%s) Null%s {
    if e == nil {
        return Null%s{}
    }

    return Null%s{
        %s: *e,
        Valid:  true,
    }
}


`, nullEnumFuncName, enumType, enumType, enumType, enumType, enumType, convertFuncName, enumType, enumType, getNullEnumFuncName, enumType, enumType, enumType, enumType, enumType)
	}

	os.WriteFile(filePath, []byte(output.String()), 0644)
}
