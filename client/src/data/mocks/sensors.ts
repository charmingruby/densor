import { v4 } from 'uuid'
import { Sensor } from '../types/sensor'

export function sensorsData(amount: number): Sensor[] {
  const sensor: Sensor[] = []

  for (let i = 0; i < amount; i++) {
    const sc: Sensor = {
      id: v4(),
      name: `Sensor ${i}`,
      description: `Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao Descriçao ${i}`,
      equipmentId: i.toString(),
      observation: `Observaçao ${i}`,
      sectorId: i.toString(),
      sensorCategory: {
        name: `Categoria ${i}`,
      },
      sensorCategoryId: i.toString(),
      status: 'On',
      createdAt: new Date(),
    }

    sensor.push(sc)
  }

  return sensor
}
