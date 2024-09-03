'use client'

import { PaddingContainer } from '@/components/padding-container'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Textarea } from '@/components/ui/textarea'
import { ArrowRight } from 'lucide-react'
import { useSensorRegisterController } from '../../../../hooks/form/use-sensor-register-controller'
import { useEffect } from 'react'
import { RequiredFieldContainer } from '@/components/required-field-indicator'
import Link from 'next/link'

export function SensorRegisterForm() {
  const { form, onSubmit, bootstrap, categories } =
    useSensorRegisterController()

  useEffect(() => {
    const fetchData = async () => {
      bootstrap()
    }

    fetchData()
  }, [])

  return (
    <Form {...form}>
      <form
        className="flex flex-col flex-1"
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <PaddingContainer
          verticalPadding
          className="flex flex-col flex-1 pb-24 gap-4"
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <RequiredFieldContainer>
                  <FormLabel>Nome</FormLabel>
                </RequiredFieldContainer>
                <FormControl>
                  <Input placeholder="Nome do sensor" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <RequiredFieldContainer>
                  <FormLabel>Descrição</FormLabel>
                </RequiredFieldContainer>
                <FormControl>
                  <Input
                    placeholder="Descrição sobre o motivo da implantação"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="equipmentName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome do equipamento</FormLabel>
                <FormControl>
                  <Input placeholder="Nome do equipamento" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="sectorName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Nome do setor</FormLabel>
                <FormControl>
                  <Input placeholder="Nome do setor" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="categoryName"
            render={({ field }) => (
              <FormItem>
                <RequiredFieldContainer>
                  <FormLabel>Categoria</FormLabel>
                </RequiredFieldContainer>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Escolher categoria" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    {categories.map(({ id, name }) => (
                      <SelectItem key={id} value={name}>
                        {name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="observation"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Observação</FormLabel>
                <FormControl>
                  <Textarea
                    placeholder="Alguma informação a ser observada"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </PaddingContainer>

        <div className="mt-auto w-full fixed bottom-0 z-10">
          <PaddingContainer className="bg-background border-t h-20 flex items-center justify-end gap-4">
            <Button asChild variant="outline">
              <Link prefetch={false} href="/sensors">
                <span>Ver sensores</span>
              </Link>
            </Button>

            <Button className="gap-1" type="submit">
              <span>Registrar</span>
              <ArrowRight className="h-4 w-4" />
            </Button>
          </PaddingContainer>
        </div>
      </form>
    </Form>
  )
}
