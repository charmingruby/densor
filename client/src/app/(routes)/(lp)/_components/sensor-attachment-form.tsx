import { PaddingContainer } from '@/components/padding-container'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

export function SensorAttachmentForm() {
  return (
    <form>
      <PaddingContainer verticalPadding className="flex flex-col flex-1 pb-24">
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>{' '}
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>
        <div>
          <Label>Email do Responsavel</Label>
          <Input />
        </div>
      </PaddingContainer>

      <div className="mt-auto w-full fixed bottom-0 z-10">
        <PaddingContainer className="bg-background border-t h-20 flex items-center justify-end">
          <Button>
            <span>Anexar</span>
          </Button>
        </PaddingContainer>
      </div>
    </form>
  )
}
