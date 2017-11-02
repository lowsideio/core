package models

import (
  "time"

  _ "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/satori/go.uuid"
)

type Motorcycle struct {
  ID                    uuid.UUID   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id" string:"id" form:"id" query:"id"`

  Year                  string      `json:"year,omitempty" string:"year" form:"year" query:"year"`
  Category              string      `json:"category,omitempty" string:"category" form:"category" query:"category"`
  Cc                    string      `json:"cc,omitempty" string:"cc" form:"cc" query:"cc"`
  EngineType            string      `json:"engineType,omitempty" string:"engineType" form:"engineType" query:"engineType"`
  Power                 string      `json:"power,omitempty" string:"power" form:"power" query:"power"`
  Torque                string      `json:"torque,omitempty" string:"torque" form:"torque" query:"torque"`
  Compression           string      `json:"compression,omitempty" string:"compression" form:"compression" query:"compression"`
  Bore                  string      `json:"bore,omitempty" string:"bore" form:"bore" query:"bore"`
  Stroke                string      `json:"stroke,omitempty" string:"stroke" form:"stroke" query:"stroke"`
  FuelSystem            string      `json:"fuelSystem,omitempty" string:"fuelSystem" form:"fuelSystem" query:"fuelSystem"`
  FuelControl           string      `json:"fuelControl,omitempty" string:"fuelControl" form:"fuelControl" query:"fuelControl"`
  Ignition              string      `json:"ignition,omitempty" string:"ignition" form:"ignition" query:"ignition"`
  Lubrication           string      `json:"lubrication,omitempty" string:"lubrication" form:"lubrication" query:"lubrication"`
  Cooling               string      `json:"cooling,omitempty" string:"cooling" form:"cooling" query:"cooling"`
  Gearbox               string      `json:"gearbox,omitempty" string:"gearbox" form:"gearbox" query:"gearbox"`
  Transmission          string      `json:"transmission,omitempty" string:"transmission" form:"transmission" query:"transmission"`
  Clutch                string      `json:"clutch,omitempty" string:"clutch" form:"clutch" query:"clutch"`
  FrameType             string      `json:"frameType,omitempty" string:"frameType" form:"frameType" query:"frameType"`
  FrontSuspension       string      `json:"frontSuspension,omitempty" string:"frontSuspension" form:"frontSuspension" query:"frontSuspension"`
  FrontWheelTravel      string      `json:"frontWheelTravel,omitempty" string:"frontWheelTravel" form:"frontWheelTravel" query:"frontWheelTravel"`
  RearSuspension        string      `json:"fearSuspension,omitempty" string:"fearSuspension" form:"fearSuspension" query:"fearSuspension"`
  RearWheelTravel       string      `json:"fearWheelTravel,omitempty" string:"fearWheelTravel" form:"fearWheelTravel" query:"fearWheelTravel"`
  FrontTyre             string      `json:"frontTyre,omitempty" string:"frontTyre" form:"frontTyre" query:"frontTyre"`
  RearTyre              string      `json:"rearTyre,omitempty" string:"rearTyre" form:"rearTyre" query:"rearTyre"`
  FrontBrakes           string      `json:"frontBrakes,omitempty" string:"frontBrakes" form:"frontBrakes" query:"frontBrakes"`
  FrontBrakesDiameter   string      `json:"frontBrakesDiameter,omitempty" string:"frontBrakesDiameter" form:"frontBrakesDiameter" query:"frontBrakesDiameter"`
  RearBrakes            string      `json:"rearBrakes,omitempty" string:"rearBrakes" form:"rearBrakes" query:"rearBrakes"`
  RearBrakesDiameter    string      `json:"rearBrakesDiameter,omitempty" string:"rearBrakesDiameter" form:"rearBrakesDiameter" query:"rearBrakesDiameter"`
  DryWeight             string      `json:"dryWeight,omitempty" string:"dryWeight" form:"dryWeight" query:"dryWeight"`
  PowerWeight           string      `json:"powerWeight,omitempty" string:"powerWeight" form:"powerWeight" query:"powerWeight"`
  SeatHeight            string      `json:"seatHeight,omitempty" string:"seatHeight" form:"seatHeight" query:"seatHeight"`
  OverallHeight         string      `json:"overallHeight,omitempty" string:"overallHeight" form:"overallHeight" query:"overallHeight"`
  OverallLength         string      `json:"overallLength,omitempty" string:"overallLength" form:"overallLength" query:"overallLength"`
  OverallWidth          string      `json:"overallWidth,omitempty" string:"overallWidth" form:"overallWidth" query:"overallWidth"`
  GroundClearence       string      `json:"groundClearence,omitempty" string:"groundClearence" form:"groundClearence" query:"groundClearence"`
  Wheelbase             string      `json:"wheelbase,omitempty" string:"wheelbase" form:"wheelbase" query:"wheelbase"`
  FuelCapacity          string      `json:"fuelCapacity,omitempty" string:"fuelCapacity" form:"fuelCapacity" query:"fuelCapacity"`
  Starter               string      `json:"starter,omitempty" string:"starter" form:"starter" query:"starter"`

  MotorcycleModelId       uuid.UUID   `gorm:"foreign_key;type:uuid()" json:"motorcycles_models_id" string:"motorcycles_models_id" form:"motorcycles_models_id" query:"motorcycles_models_id"`


  CreatedAt             time.Time   `json:"createdAt,omitempty" string:"createdAt" form:"createdAt" query:"createdAt"`
  UpdatedAt             time.Time   `json:"updatedAt,omitempty" string:"updatedAt" form:"updatedAt" query:"updatedAt"`
  DeletedAt             time.Time   `gorm:"default:null" json:"deletedAt" string:"deletedAt" form:"deletedAt" query:"deletedAt"`

  Price uint `json:"price,omitempty" string:"price" form:"price" query:"price"`
}

func (m Motorcycle) TableName() string {
  return "motorcycles_specs";
}
