import { PaddingContainer } from '@/components/padding-container'

import { SearchSensorForm } from './_components/search-sensor-form'

export default function SensorsPage() {
  return (
    <div className="flex flex-col flex-1">
      <PaddingContainer
        verticalPadding
        className="flex flex-col flex-1 pb-8 gap-6"
      >
        <SearchSensorForm />
      </PaddingContainer>
    </div>
  )
}
