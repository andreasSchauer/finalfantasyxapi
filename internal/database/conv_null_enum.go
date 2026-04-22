package database

func ToNullNullifyArmored(s *string) NullNullifyArmored {
    if s == nil {
        return NullNullifyArmored{}
    }
    
    return NullNullifyArmored{
        NullifyArmored: NullifyArmored(*s),
        Valid: true,
    }
}

func ConvertNullNullifyArmored(ne NullNullifyArmored) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.NullifyArmored)
    return &val
}

func GetNullNullifyArmored(e *NullifyArmored) NullNullifyArmored {
    if e == nil {
        return NullNullifyArmored{}
    }

    return NullNullifyArmored{
        NullifyArmored: *e,
        Valid:  true,
    }
}


func ToNullCreationsUnlockedCategory(s *string) NullCreationsUnlockedCategory {
    if s == nil {
        return NullCreationsUnlockedCategory{}
    }
    
    return NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: CreationsUnlockedCategory(*s),
        Valid: true,
    }
}

func ConvertNullCreationsUnlockedCategory(ne NullCreationsUnlockedCategory) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CreationsUnlockedCategory)
    return &val
}

func GetNullCreationsUnlockedCategory(e *CreationsUnlockedCategory) NullCreationsUnlockedCategory {
    if e == nil {
        return NullCreationsUnlockedCategory{}
    }

    return NullCreationsUnlockedCategory{
        CreationsUnlockedCategory: *e,
        Valid:  true,
    }
}


func ToNullCounterType(s *string) NullCounterType {
    if s == nil {
        return NullCounterType{}
    }
    
    return NullCounterType{
        CounterType: CounterType(*s),
        Valid: true,
    }
}

func ConvertNullCounterType(ne NullCounterType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CounterType)
    return &val
}

func GetNullCounterType(e *CounterType) NullCounterType {
    if e == nil {
        return NullCounterType{}
    }

    return NullCounterType{
        CounterType: *e,
        Valid:  true,
    }
}


func ToNullMaCreationArea(s *string) NullMaCreationArea {
    if s == nil {
        return NullMaCreationArea{}
    }
    
    return NullMaCreationArea{
        MaCreationArea: MaCreationArea(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationArea(ne NullMaCreationArea) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationArea)
    return &val
}

func GetNullMaCreationArea(e *MaCreationArea) NullMaCreationArea {
    if e == nil {
        return NullMaCreationArea{}
    }

    return NullMaCreationArea{
        MaCreationArea: *e,
        Valid:  true,
    }
}


func ToNullMaCreationCategory(s *string) NullMaCreationCategory {
    if s == nil {
        return NullMaCreationCategory{}
    }
    
    return NullMaCreationCategory{
        MaCreationCategory: MaCreationCategory(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationCategory(ne NullMaCreationCategory) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationCategory)
    return &val
}

func GetNullMaCreationCategory(e *MaCreationCategory) NullMaCreationCategory {
    if e == nil {
        return NullMaCreationCategory{}
    }

    return NullMaCreationCategory{
        MaCreationCategory: *e,
        Valid:  true,
    }
}


func ToNullMaCreationSpecies(s *string) NullMaCreationSpecies {
    if s == nil {
        return NullMaCreationSpecies{}
    }
    
    return NullMaCreationSpecies{
        MaCreationSpecies: MaCreationSpecies(*s),
        Valid: true,
    }
}

func ConvertNullMaCreationSpecies(ne NullMaCreationSpecies) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MaCreationSpecies)
    return &val
}

func GetNullMaCreationSpecies(e *MaCreationSpecies) NullMaCreationSpecies {
    if e == nil {
        return NullMaCreationSpecies{}
    }

    return NullMaCreationSpecies{
        MaCreationSpecies: *e,
        Valid:  true,
    }
}


func ToNullTargetType(s *string) NullTargetType {
    if s == nil {
        return NullTargetType{}
    }
    
    return NullTargetType{
        TargetType: TargetType(*s),
        Valid: true,
    }
}

func ConvertNullTargetType(ne NullTargetType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.TargetType)
    return &val
}

func GetNullTargetType(e *TargetType) NullTargetType {
    if e == nil {
        return NullTargetType{}
    }

    return NullTargetType{
        TargetType: *e,
        Valid:  true,
    }
}


func ToNullMusicUseCase(s *string) NullMusicUseCase {
    if s == nil {
        return NullMusicUseCase{}
    }
    
    return NullMusicUseCase{
        MusicUseCase: MusicUseCase(*s),
        Valid: true,
    }
}

func ConvertNullMusicUseCase(ne NullMusicUseCase) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.MusicUseCase)
    return &val
}

func GetNullMusicUseCase(e *MusicUseCase) NullMusicUseCase {
    if e == nil {
        return NullMusicUseCase{}
    }

    return NullMusicUseCase{
        MusicUseCase: *e,
        Valid:  true,
    }
}


func ToNullBgReplacementType(s *string) NullBgReplacementType {
    if s == nil {
        return NullBgReplacementType{}
    }
    
    return NullBgReplacementType{
        BgReplacementType: BgReplacementType(*s),
        Valid: true,
    }
}

func ConvertNullBgReplacementType(ne NullBgReplacementType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.BgReplacementType)
    return &val
}

func GetNullBgReplacementType(e *BgReplacementType) NullBgReplacementType {
    if e == nil {
        return NullBgReplacementType{}
    }

    return NullBgReplacementType{
        BgReplacementType: *e,
        Valid:  true,
    }
}


func ToNullSpecialActionType(s *string) NullSpecialActionType {
    if s == nil {
        return NullSpecialActionType{}
    }
    
    return NullSpecialActionType{
        SpecialActionType: SpecialActionType(*s),
        Valid: true,
    }
}

func ConvertNullSpecialActionType(ne NullSpecialActionType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.SpecialActionType)
    return &val
}

func GetNullSpecialActionType(e *SpecialActionType) NullSpecialActionType {
    if e == nil {
        return NullSpecialActionType{}
    }

    return NullSpecialActionType{
        SpecialActionType: *e,
        Valid:  true,
    }
}


func ToNullCriticalType(s *string) NullCriticalType {
    if s == nil {
        return NullCriticalType{}
    }
    
    return NullCriticalType{
        CriticalType: CriticalType(*s),
        Valid: true,
    }
}

func ConvertNullCriticalType(ne NullCriticalType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.CriticalType)
    return &val
}

func GetNullCriticalType(e *CriticalType) NullCriticalType {
    if e == nil {
        return NullCriticalType{}
    }

    return NullCriticalType{
        CriticalType: *e,
        Valid:  true,
    }
}


func ToNullBreakDmgLmtType(s *string) NullBreakDmgLmtType {
    if s == nil {
        return NullBreakDmgLmtType{}
    }
    
    return NullBreakDmgLmtType{
        BreakDmgLmtType: BreakDmgLmtType(*s),
        Valid: true,
    }
}

func ConvertNullBreakDmgLmtType(ne NullBreakDmgLmtType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.BreakDmgLmtType)
    return &val
}

func GetNullBreakDmgLmtType(e *BreakDmgLmtType) NullBreakDmgLmtType {
    if e == nil {
        return NullBreakDmgLmtType{}
    }

    return NullBreakDmgLmtType{
        BreakDmgLmtType: *e,
        Valid:  true,
    }
}


func ToNullComposer(s *string) NullComposer {
    if s == nil {
        return NullComposer{}
    }
    
    return NullComposer{
        Composer: Composer(*s),
        Valid: true,
    }
}

func ConvertNullComposer(ne NullComposer) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.Composer)
    return &val
}

func GetNullComposer(e *Composer) NullComposer {
    if e == nil {
        return NullComposer{}
    }

    return NullComposer{
        Composer: *e,
        Valid:  true,
    }
}


func ToNullArranger(s *string) NullArranger {
    if s == nil {
        return NullArranger{}
    }
    
    return NullArranger{
        Arranger: Arranger(*s),
        Valid: true,
    }
}

func ConvertNullArranger(ne NullArranger) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.Arranger)
    return &val
}

func GetNullArranger(e *Arranger) NullArranger {
    if e == nil {
        return NullArranger{}
    }

    return NullArranger{
        Arranger: *e,
        Valid:  true,
    }
}


func ToNullShopType(s *string) NullShopType {
    if s == nil {
        return NullShopType{}
    }
    
    return NullShopType{
        ShopType: ShopType(*s),
        Valid: true,
    }
}

func ConvertNullShopType(ne NullShopType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.ShopType)
    return &val
}

func GetNullShopType(e *ShopType) NullShopType {
    if e == nil {
        return NullShopType{}
    }

    return NullShopType{
        ShopType: *e,
        Valid:  true,
    }
}


func ToNullAvailabilityType(s *string) NullAvailabilityType {
    if s == nil {
        return NullAvailabilityType{}
    }
    
    return NullAvailabilityType{
        AvailabilityType: AvailabilityType(*s),
        Valid: true,
    }
}

func ConvertNullAvailabilityType(ne NullAvailabilityType) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.AvailabilityType)
    return &val
}

func GetNullAvailabilityType(e *AvailabilityType) NullAvailabilityType {
    if e == nil {
        return NullAvailabilityType{}
    }

    return NullAvailabilityType{
        AvailabilityType: *e,
        Valid:  true,
    }
}


func ToNullNodeState(s *string) NullNodeState {
    if s == nil {
        return NullNodeState{}
    }
    
    return NullNodeState{
        NodeState: NodeState(*s),
        Valid: true,
    }
}

func ConvertNullNodeState(ne NullNodeState) *string {
    if !ne.Valid {
        return nil
    }

    val := string(ne.NodeState)
    return &val
}

func GetNullNodeState(e *NodeState) NullNodeState {
    if e == nil {
        return NullNodeState{}
    }

    return NullNodeState{
        NodeState: *e,
        Valid:  true,
    }
}


