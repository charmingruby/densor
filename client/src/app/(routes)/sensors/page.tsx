import { PaddingContainer } from '@/components/padding-container'
import { SearchSensorForm } from './_components/search-sensor-form'
import { RegisterLink } from './_components/register-link'

export default function SensorsPage() {
  return (
    <div className="flex flex-col flex-1">
      <PaddingContainer
        verticalPadding
        className="flex flex-col flex-1 pb-28 gap-6"
      >
        <SearchSensorForm />
      </PaddingContainer>

      <RegisterLink />
    </div>
  )
}
