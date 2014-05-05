package json

import (
	"dota/testutils"
	"testing"
)

func Test012(t *testing.T) {
	CreateJSON(nil)
}

func Test000(t *testing.T) {
	answer := "null"
	result := CreateJSON(nil)
	testutils.Assert(result, answer, t)
}

// func Test002(t *testing.T) {
// 	data := map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"heroes": []structs.Hero{structs.DefaultHeroStruct},
// 		},
// 	}

// 	answer := "{\"data\":{\"heroes\":[" + structs.DefaultHeroJSON + "]}}"
// 	result := CreateJSON(data)
// 	testutils.Assert(result, answer, t)
// }

// func Test003(t *testing.T) {
// 	data := map[string]interface{}{
// 		"data": map[string]interface{}{
// 			"heroes": []structs.Hero{structs.DefaultHeroStruct},
// 			"roles":  []structs.Role{structs.DefaultRoleStruct},
// 		},
// 	}

// 	answer := "{\"data\":{\"heroes\":[" + structs.DefaultHeroJSON + "],\"roles\":[" + structs.DefaultRoleJSON + "]}}"
// 	result := CreateJSON(data)
// 	testutils.Assert(result, answer, t)
// }
