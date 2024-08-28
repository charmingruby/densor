import { CategoryTag } from './category-tag'
import { StatusAnnotation } from './status-annotation'

interface SensorCardProps {
  name: string
  description: string
  sensorCategoryName: string
  status: string
}

export function SensorCard({
  name,
  description,
  sensorCategoryName,
  status,
}: SensorCardProps) {
  return (
    <div className="border px-4 py-6 rounded">
      <div className="flex justify-between items-start gap-8">
        <div className="space-y-2">
          <CategoryTag category={sensorCategoryName} />

          <div className="flex flex-col gap-1">
            <strong>{name}</strong>
            <span className="text-sm text-muted-foreground">{description}</span>
          </div>
        </div>

        <StatusAnnotation status={status} />
      </div>
    </div>
  )
}
