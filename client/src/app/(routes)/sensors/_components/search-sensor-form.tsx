'use client'

import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Search } from 'lucide-react'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { useSearchSensorController } from './use-search-sensor-controller'
import { Form, FormControl, FormField, FormItem } from '@/components/ui/form'
import { SensorCard } from './sensor-card'
import { useEffect } from 'react'

export function SearchSensorForm() {
  const { bootstrap, sensorsResult, form, onSubmit, statusOptions } =
    useSearchSensorController()

  useEffect(() => bootstrap(), [])

  return (
    <>
      <Form {...form}>
        <form
          className="flex items-center gap-2 w-full"
          onSubmit={form.handleSubmit(onSubmit)}
        >
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormControl>
                  <Input
                    placeholder="Pesquisar por nome do sensor..."
                    {...field}
                  />
                </FormControl>
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="status"
            render={({ field }) => (
              <FormItem>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger className="w-[180px]">
                      <SelectValue placeholder="Escolher status" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectItem key={-1} value={'None'}>
                      Qualquer um
                    </SelectItem>
                    {statusOptions.map((sts, idx) => (
                      <SelectItem key={idx} value={sts}>
                        {sts}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </FormItem>
            )}
          />

          <div>
            <Button size="icon">
              <Search width={18} />
            </Button>
          </div>
        </form>
      </Form>

      <div className="flex flex-col gap-4">
        {sensorsResult.map(
          ({ id, name, description, sensorCategory, status }) => (
            <SensorCard
              key={id}
              name={name}
              description={description}
              sensorCategoryName={sensorCategory.name}
              status={status}
            />
          ),
        )}
      </div>
    </>
  )
}
