package marshal

import (
	"../powscl-alg"
	"bytes"
	"encoding/json"
)

func MarshalTeams(Teams map[int]powalg.Team) string {
	mapstr := `[`
	i := 0
	for _, v := range Teams {
		val, _ := json.Marshal(v)
		mapstr += bytes.NewBuffer(val).String()
		if i != len(Teams)-1 {
			mapstr += ","
		}
		i++
	}
	mapstr += `]`
	return mapstr
}

func MarshalTeam(Team powalg.Team) string {
	val, _ := json.Marshal(Team)
	return bytes.NewBuffer(val).String()
}

func MarshalMatch(Match []powalg.Match) string {
	val, _ := json.Marshal(Match)
	return bytes.NewBuffer(val).String()
}

func UnmarshalTeam(j []byte) powalg.Team {
	var t powalg.Team
	json.Unmarshal(j, &t)
	return t
}

func UnmarshalMatch(j []byte) powalg.Match {
	var t powalg.Match
	json.Unmarshal(j, &t)
	return t
}
