import { ChangeEventHandler, RefObject } from 'react'

interface Props {
  ref: RefObject<HTMLInputElement | null>
  handleImageChange: ChangeEventHandler<HTMLInputElement>
}

export default function HiddenInput({ ref, handleImageChange }: Props) {
  return (
    <input
      type='file'
      accept='image/*'
      ref={ref}
      hidden
      onChange={handleImageChange}
      aria-hidden='true'
      tabIndex={-1}
    />
  )
}
