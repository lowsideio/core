package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Motorcycle struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id" string:"id" form:"id" query:"id"`

	Year                string `json:"year" string:"year" form:"year" query:"year"`
	Category            string `json:"category" string:"category" form:"category" query:"category"`
	Cc                  string `json:"cc" string:"cc" form:"cc" query:"cc"`
	EngineType          string `json:"engineType" string:"engineType" form:"engineType" query:"engineType"`
	Power               string `json:"power" string:"power" form:"power" query:"power"`
	Torque              string `json:"torque" string:"torque" form:"torque" query:"torque"`
	Compression         string `json:"compression" string:"compression" form:"compression" query:"compression"`
	Bore                string `json:"bore" string:"bore" form:"bore" query:"bore"`
	Stroke              string `json:"stroke" string:"stroke" form:"stroke" query:"stroke"`
	FuelSystem          string `json:"fuelSystem" string:"fuelSystem" form:"fuelSystem" query:"fuelSystem"`
	FuelControl         string `json:"fuelControl" string:"fuelControl" form:"fuelControl" query:"fuelControl"`
	Ignition            string `json:"ignition" string:"ignition" form:"ignition" query:"ignition"`
	Lubrication         string `json:"lubrication" string:"lubrication" form:"lubrication" query:"lubrication"`
	Cooling             string `json:"cooling" string:"cooling" form:"cooling" query:"cooling"`
	Gearbox             string `json:"gearbox" string:"gearbox" form:"gearbox" query:"gearbox"`
	Transmission        string `json:"transmission" string:"transmission" form:"transmission" query:"transmission"`
	Clutch              string `json:"clutch" string:"clutch" form:"clutch" query:"clutch"`
	FrameType           string `json:"frameType" string:"frameType" form:"frameType" query:"frameType"`
	FrontSuspension     string `json:"frontSuspension" string:"frontSuspension" form:"frontSuspension" query:"frontSuspension"`
	FrontWheelTravel    string `json:"frontWheelTravel" string:"frontWheelTravel" form:"frontWheelTravel" query:"frontWheelTravel"`
	RearSuspension      string `json:"rearSuspension" string:"rearSuspension" form:"rearSuspension" query:"rearSuspension"`
	RearWheelTravel     string `json:"rearWheelTravel" string:"rearWheelTravel" form:"rearWheelTravel" query:"rearWheelTravel"`
	FrontTyre           string `json:"frontTyre" string:"frontTyre" form:"frontTyre" query:"frontTyre"`
	RearTyre            string `json:"rearTyre" string:"rearTyre" form:"rearTyre" query:"rearTyre"`
	FrontBrakes         string `json:"frontBrakes" string:"frontBrakes" form:"frontBrakes" query:"frontBrakes"`
	FrontBrakesDiameter string `json:"frontBrakesDiameter" string:"frontBrakesDiameter" form:"frontBrakesDiameter" query:"frontBrakesDiameter"`
	RearBrakes          string `json:"rearBrakes" string:"rearBrakes" form:"rearBrakes" query:"rearBrakes"`
	RearBrakesDiameter  string `json:"rearBrakesDiameter" string:"rearBrakesDiameter" form:"rearBrakesDiameter" query:"rearBrakesDiameter"`
	DryWeight           string `json:"dryWeight" string:"dryWeight" form:"dryWeight" query:"dryWeight"`
	PowerWeight         string `json:"powerWeight" string:"powerWeight" form:"powerWeight" query:"powerWeight"`
	SeatHeight          string `json:"seatHeight" string:"seatHeight" form:"seatHeight" query:"seatHeight"`
	OverallHeight       string `json:"overallHeight" string:"overallHeight" form:"overallHeight" query:"overallHeight"`
	OverallLength       string `json:"overallLength" string:"overallLength" form:"overallLength" query:"overallLength"`
	OverallWidth        string `json:"overallWidth" string:"overallWidth" form:"overallWidth" query:"overallWidth"`
	GroundClearence     string `json:"groundClearence" string:"groundClearence" form:"groundClearence" query:"groundClearence"`
	Wheelbase           string `json:"wheelbase" string:"wheelbase" form:"wheelbase" query:"wheelbase"`
	FuelCapacity        string `json:"fuelCapacity" string:"fuelCapacity" form:"fuelCapacity" query:"fuelCapacity"`
	Starter             string `json:"starter" string:"starter" form:"starter" query:"starter"`

	MotorcycleModelId uuid.UUID `gorm:"foreign_key;type:uuid()" json:"motorcycles_models_id" string:"motorcycles_models_id" form:"motorcycles_models_id" query:"motorcycles_models_id"`

	CreatedAt time.Time `json:"createdAt" string:"createdAt" form:"createdAt" query:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" string:"updatedAt" form:"updatedAt" query:"updatedAt"`
	DeletedAt time.Time `gorm:"default:null" json:"deletedAt" string:"deletedAt" form:"deletedAt" query:"deletedAt"`

	Price uint `json:"price" string:"price" form:"price" query:"price"`
}

func (m Motorcycle) TableName() string {
	return "motorcycles_specs"
}
