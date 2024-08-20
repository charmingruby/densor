import { PaddingContainer } from '@/components/padding-container'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Textarea } from '@/components/ui/textarea'
import { ArrowRight } from 'lucide-react'

export function SensorRegisterForm() {
  return (
    <form className="flex flex-col flex-1">
      <PaddingContainer
        verticalPadding
        className="flex flex-col flex-1 pb-24 gap-4"
      >
        <div className="space-y-2">
          <Label>Nome</Label>
          <Input placeholder="Nome do sensor" />
        </div>

        <div className="space-y-2">
          <Label>Descrição</Label>
          <Textarea placeholder="Descrição sobre o motivo da implantação" />
        </div>

        <div className="space-y-2">
          <Label>Código do equipamento</Label>
          <Input placeholder="Identificador do equipamento" />
        </div>

        <div className="space-y-2">
          <Label>Código do setor</Label>
          <Input placeholder="Identificador do setor" />
        </div>

        <div className="space-y-2">
          <Label>Categoria do sensor</Label>
          <Select>
            <SelectTrigger className="w-full">
              <SelectValue placeholder="Escolher categoria" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="light">Movimento</SelectItem>
              <SelectItem value="dark">Temperatura</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <div className="space-y-2">
          <Label>Observação</Label>
          <Textarea placeholder="Alguma informação a ser observada" />
        </div>
      </PaddingContainer>

      <div className="mt-auto w-full fixed bottom-0 z-10">
        <PaddingContainer className="bg-background border-t h-20 flex items-center justify-end">
          <Button className="gap-1">
            <span>Registrar</span>
            <ArrowRight className="h-4 w-4" />
          </Button>
        </PaddingContainer>
      </div>
    </form>
  )
}
