import { $api } from '@/shared/api/api'
import { startTransition, useActionState } from 'react'
import { toast } from 'sonner'

export const useSendCode = () => {
  const [state, action, isPending] = useActionState(sendCode, { success: false })
  const handleSendCode = () => {
    startTransition(action)
  }
  return { handleSendCode, isPending, state }
}

async function sendCode() {
  try {
    await $api.post('/auth/send-code')
    toast.success('Код отправлен на почту')
    return { success: true }
  } catch (error) {
    toast.error('Ошибка отправки кода')
    return { success: false }
  }
}
