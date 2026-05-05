package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type LookupObject struct {
	Name    string
	Version *int32
}

func (l LookupObject) ToKeyFields() []any {
	return []any{
		l.Name,
		h.DerefOrNil(l.Version),
	}
}

func (l LookupObject) Error() string {
	return fmt.Sprintf("lookup object '%s'", h.NameToString(l.Name, l.Version, nil))
}

func GetResourceByID[T LookupableID](id int32, lookup map[int32]T) (T, error) {
	resource, found := lookup[id]
	if !found {
		var zeroType T
		return zeroType, fmt.Errorf("couldn't find %s with id '%d'.", getTypeName[T](), id)
	}

	return resource, nil
}

func GetResource[T LookupableKey, K any](key K, lookup map[string]T) (T, error) {
	switch k := any(key).(type) {
	case string:
		return getResourceByName(k, lookup)
	case LookupableKey:
		return getResourceByKey(k, lookup)
	default:
		var zeroType T
		return zeroType, fmt.Errorf("key must be either string or Lookupable, got %T", key)
	}
}

func getResourceByName[T LookupableKey](key string, lookup map[string]T) (T, error) {
	resource, found := lookup[key]
	if !found {
		var zeroType T
		return zeroType, h.NewErr(key, fmt.Errorf("couldn't find %s '%s'.", getTypeName[T](), key))
	}

	return resource, nil
}

func getResourceByKey[T, K LookupableKey](obj K, lookup map[string]T) (T, error) {
	key := Key(obj)

	resource, err := GetResource(key, lookup)
	if err != nil {
		var zeroType T
		return zeroType, h.NewErr(obj.Error(), fmt.Errorf("couldn't find %s.", getTypeName[T]()))
	}

	return resource, nil
}

func getResources[T LookupableKey, K any](keys []K, lookup map[string]T) ([]T, error) {
	objects := make([]T, len(keys))

	for i, key := range keys {
		obj, err := GetResource(key, lookup)
		if err != nil {
			return nil, err
		}

		objects[i] = obj
	}

	return objects, nil
}

func typedRefsToResources[T LookupableKey](refs []AbilityReference, lookup map[string]T) ([]T, error) {
	objects := make([]T, len(refs))

	for i, ref := range refs {
		obj, err := GetResource(ref.Untyped(), lookup)
		if err != nil {
			return nil, err
		}

		objects[i] = obj
	}

	return objects, nil
}
