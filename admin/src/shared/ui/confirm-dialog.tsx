'use client'
import { PropsWithChildren, useState } from 'react'
import {
  Dialog,
  DialogContent,
  DialogTrigger,
  DialogTitle,
  DialogDescription,
  DialogFooter,
  DialogClose,
} from './dialog'
import { Button } from './button'

interface Props extends PropsWithChildren {
  handleConfirm(): void
  title: string
  description: string
  confirmLabel?: string
  cancelLabel?: string
}

export default function ConfirmDialog({
  children,
  title,
  description,
  confirmLabel,
  cancelLabel,
  handleConfirm,
}: Props) {
  const [open, setOpen] = useState(false)
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogTitle>{title}</DialogTitle>
        <DialogDescription>{description}</DialogDescription>
        <DialogFooter>
          <DialogClose asChild>
            <Button onClick={handleConfirm}>{confirmLabel ? confirmLabel : 'Подтвердить'} </Button>
          </DialogClose>
          <DialogClose asChild>
            <Button variant='secondary'> {cancelLabel ? cancelLabel : 'Отменить'}</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
