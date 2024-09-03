import { SensorCategory } from '../types/sensor-category'

interface APISensorCategory {
  id: string
  name: string
  description: string
  created_at: string
}

export interface APIData {
  message: string
  data: APISensorCategory[]
}

export function parseFetchSensorCategoriesResult({
  data,
}: APIData): SensorCategory[] {
  const parsedResults = data.map((category) => {
    const parsedCategory: SensorCategory = {
      id: category.id,
      name: category.name,
      description: category.description,
      createdAt: new Date(category.created_at),
    }
    return parsedCategory
  })

  return parsedResults
}
