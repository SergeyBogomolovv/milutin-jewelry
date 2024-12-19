import { fetcher } from '@/shared/api/fetcher'
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
    const res = await fetcher('/auth/send-code', { method: 'POST' })
    if (!res.ok) {
      toast.error('Ошибка отправки кода')
      return { success: false }
    }
    toast.success('Код отправлен на почту')
    return { success: true }
  } catch (error) {
    toast.error('Ошибка отправки кода')
    return { success: false }
  }
}
