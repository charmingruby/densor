'use client'

import { sensorStatus } from '@/data/mocks/sensor-status'
import { sensorsData } from '@/data/mocks/sensors'
import { Sensor } from '@/data/types/sensor'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'

import { useForm } from 'react-hook-form'
import { z } from 'zod'

const formSchema = z.object({
  name: z.string(),
  status: z.string(),
})

type FormData = z.infer<typeof formSchema>

export function useSearchSensorController() {
  const [sensors, setSensors] = useState<Sensor[]>([])
  const [sensorsResult, setSensorsResult] = useState<Sensor[]>([])
  const statusOptions = sensorStatus

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      status: '',
    },
  })

  const bootstrap = () => {
    const sensors = sensorsData(20)
    setSensors(sensors)
    setSensorsResult(sensors)
  }

  function onSubmit(values: FormData) {
    let filteredSensors: Sensor[] = sensors

    if (values.status !== 'None' && values.status) {
      filteredSensors = filteredSensors.filter(
        (s) => s.status === values.status,
      )
      console.log('stat pos', filteredSensors)
    }

    if (values.name.trim() !== '') {
      const searchName = values.name.toLowerCase()
      filteredSensors = filteredSensors.filter((s) =>
        s.name.toLowerCase().includes(searchName),
      )
    }

    setSensorsResult(filteredSensors)
  }

  return { bootstrap, sensorsResult, form, onSubmit, statusOptions }
}
