'use client'

import {  mergeRefs } from '@/lib'
import clsx from 'clsx'
import { forwardRef, useRef } from 'react'
import { useButton, useFocusRing, useHover,mergeProps } from 'react-aria'
import { twMerge } from 'tailwind-merge'
import {variants,Loading, ButtonProps} from '../cvas/events'




const Button = forwardRef<HTMLButtonElement,ButtonProps> (
    (props, forwardedRef) => {
        const { className, children, variant, size, loading, disabled,icon, ...rest } =
      props
      const ref = useRef<HTMLButtonElement>(null)
      const { focusProps, isFocusVisible } = useFocusRing()
      const { hoverProps, isHovered } = useHover({
        ...props,
        isDisabled: disabled,
      })
      const { buttonProps, isPressed } = useButton(
        {
          ...rest,
          isDisabled: disabled,
        },
        ref
      )
  
      return (
        <button
          ref={mergeRefs([ref, forwardedRef])}
          className={twMerge(clsx(variants({ variant, size,className })))}
          {...mergeProps(buttonProps, focusProps, hoverProps)}
          data-pressed={isPressed || undefined}
          data-hovered={isHovered || undefined}
          data-focus-visible={isFocusVisible || undefined}
        >
          {loading && <Loading variant={variant} />}
          <span
            className={clsx('transition', {
              'opacity-0': loading,
              'opacity-100': !loading,
            },'inline-flex')}
          >
             {icon && (
                <span className="mr-2">
                  {icon}
                </span>
              )}
            {children}
          </span>
        </button>
      )
    }
  )

Button.displayName = 'Button'
export { Button }
export type { ButtonProps }