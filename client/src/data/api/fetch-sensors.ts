import { Sensor } from '../types/sensor'

export interface APISensor {
  ID: string
  Name: string
  Description: string
  SensorCategoryID: string
  EquipmentID: string
  Status: string
  Observation: string
  SectorID: string
  CreatedAt: string
}

export interface APIData {
  message: string
  data: APISensor[]
}

export function parseFetchSensorResult({ data }: APIData): Sensor[] {
  const parsedResults = data.map((sensor) => {
    const parsedSensor: Sensor = {
      id: sensor.ID,
      name: sensor.Name,
      description: sensor.Description,
      sensorCategoryId: sensor.SensorCategoryID,
      equipmentId: sensor.EquipmentID,
      status: sensor.Status,
      observation: sensor.Observation,
      sectorId: sensor.SectorID,
      createdAt: new Date(sensor.CreatedAt),
    }
    return parsedSensor
  })

  return parsedResults
}
