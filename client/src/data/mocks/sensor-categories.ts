import { SensorCategory } from '../types/sensor-category'
import { v4 } from 'uuid'

export function sensorCategories(amount: number): SensorCategory[] {
  const sensorCategories: SensorCategory[] = []

  for (let i = 0; i < amount; i++) {
    const sc: SensorCategory = {
      id: v4(),
      name: `Categoria ${i}`,
      description: `Descriçao ${i}`,
      createdAt: new Date(),
    }

    sensorCategories.push(sc)
  }

  return sensorCategories
}
