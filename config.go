package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	RigidTheets float64 `json:"rigid_theets"`
	DiameterH   float64 `json:"diameter_h"`
	DiameterV   float64 `json:"diameter_v"`

	Scale      float64 `json:"scale"`
	TipCenterX float64 `json:"tip_center_x"`
	TipCenterY float64 `json:"tip_center_y"`
	TipRadius  float64 `json:"tip_radius"`
	TipStop    float64 `json:"tip_stop"`

	BottomCenterX   float64 `json:"bottom_center_x"`
	BottomCenterY   float64 `json:"bottom_center_y"`
	BottomRadius    float64 `json:"bottom_radius"`
	BottomStopFlex  float64 `json:"bottom_stop_flex"`
	BottomStopRigid float64 `json:"bottom_stop_rigid"`

	FlexCircumference float64 `json:"flex_circumference"`
}

func getConfig(cf string) (*Config, error) {
	f, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if err = json.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
