import { CircleAlert } from 'lucide-react'
import { StatusAnnotation } from './status-annotation'

interface SensorCardProps {
  name: string
  description: string
  observation: string
  status: string
}

export function SensorCard({
  name,
  description,
  observation,
  status,
}: SensorCardProps) {
  return (
    <div className="border px-4 py-6 rounded">
      <div className="flex justify-between items-start gap-8">
        <div className="space-y-2">
          <div className="flex flex-col gap-1">
            <strong>{name}</strong>
            {observation && (
              <div className="flex gap-1 items-center">
                <CircleAlert
                  strokeWidth={2.5}
                  width={12}
                  className="text-primary"
                />
                <span className="text-muted-foreground text-sm">
                  {observation}
                </span>
              </div>
            )}
          </div>

          <div>
            <span className="text-base">{description}</span>
          </div>
        </div>

        <StatusAnnotation status={status} />
      </div>
    </div>
  )
}
