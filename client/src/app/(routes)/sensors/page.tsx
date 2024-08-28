import { PaddingContainer } from '@/components/padding-container'
import { sensorsData } from '@/data/mocks/sensors'
import { SensorCard } from './_components/sensor-card'

export default function SensorsPage() {
  const sensors = sensorsData(20)

  return (
    <div className="flex flex-col flex-1">
      <PaddingContainer
        verticalPadding
        className="flex flex-col flex-1 pb-4 gap-4"
      >
        {sensors.map(({ id, name, description, sensorCategory, status }) => (
          <SensorCard
            key={id}
            name={name}
            description={description}
            sensorCategoryName={sensorCategory.name}
            status={status}
          />
        ))}
      </PaddingContainer>
    </div>
  )
}
