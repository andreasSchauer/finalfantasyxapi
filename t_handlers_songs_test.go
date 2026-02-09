package main

import (
	//"net/http"
	"testing"

	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func TestGetSong(t *testing.T) {
	tests := []expSong{
		
	}

	testSingleResources(t, tests, "GetSongs", testCfg.HandleSongs, compareSongs)
}

func TestRetrieveSongs(t *testing.T) {
	tests := []expListIDs{
		
	}

	testIdList(t, tests, testCfg.e.songs.endpoint, "RetrieveSongs", testCfg.HandleSongs, compareAPIResourceLists[NamedApiResourceList])
}