export interface Sensor {
  id: string
  name: string
  description: string
  sensorCategoryId: string
  sensorCategory: {
    name: string
  }
  equipmentId: string
  status: string
  observation: string
  sectorId: string
  createdAt: Date
}
