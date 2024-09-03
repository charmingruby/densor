import { Sensor } from '../types/sensor'

export interface APISensor {
  id: string
  name: string
  description: string
  sensor_category_id: string
  equipment_id: string
  status: string
  observation: string
  sector_id: string
  created_at: string
}

export interface APIData {
  message: string
  data: APISensor[]
}

export function parseFetchSensorResult({ data }: APIData): Sensor[] {
  const parsedResults = data.map((sensor) => {
    const parsedSensor: Sensor = {
      id: sensor.id,
      name: sensor.name,
      description: sensor.description,
      sensorCategoryId: sensor.sensor_category_id,
      equipmentId: sensor.equipment_id,
      status: sensor.status,
      observation: sensor.observation,
      sectorId: sensor.sector_id,
      createdAt: new Date(sensor.created_at),
    }
    return parsedSensor
  })

  return parsedResults
}
