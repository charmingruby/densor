import { AudioWaveform } from 'lucide-react'
import { PaddingContainer } from './padding-container'

export function Header() {
  return (
    <div className="border-b fixed flex items-center max-w-4xl w-full">
      <PaddingContainer className="flex items-center gap-2 h-20 w-full">
        <div>
          <AudioWaveform className="text-primary h-7 w-7" />
        </div>

        <strong className="text-2xl">Densor</strong>
      </PaddingContainer>
    </div>
  )
}
