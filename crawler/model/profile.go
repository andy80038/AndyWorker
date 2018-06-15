package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

func fromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	a, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(a, &profile)
	return profile, err
}
