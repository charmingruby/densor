'use client'

import { parseFetchSensorResult } from '@/data/api/fetch-sensors'
import { sensorStatus } from '@/data/types/sensor-status'
import { Sensor } from '@/data/types/sensor'
import { api } from '@/lib/axios'
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

  const bootstrap = async () => {
    const { data } = await api.get('/sensors')

    const result = parseFetchSensorResult(data)

    setSensors(result)
    setSensorsResult(result)
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
