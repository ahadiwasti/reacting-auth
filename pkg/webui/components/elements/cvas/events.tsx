'use client'
import {VariantProps, cva } from 'class-variance-authority'
import { AriaButtonProps, AriaSelectProps} from 'react-aria'

export const variants = cva(
    [
      'inline-flex',
      'items-center',
      'justify-center',
      'relative',
      'cursor-pointer',
      'disabled:cursor-not-allowed',
      'tracking-wide',
      'transition-[background-color,box-shadow,text-color,transform]',
      'duration-200',
      'rounded-full',
      'outline-none',
    ],
    {
        variants:{
            variant:{
                primary:[
                    'font-semibold',
                    'bg-green-400',
                    'data-[hovered=true]:bg-green-500',
                    'text-white',
                    'shadow',
                    'data-[hovered=true]:shadow-md',
                    'disabled:bg-green-400/50',
                    'disabled:shadow',
                    'data-[focus-visible]:ring-green-400/70',
                    'data-[pressed=true]:scale-[0.98]',
                    'data-[focus-visible]:ring-2',
                    'data-[focus-visible]:ring-offset-2',
                ],
                secondary:[
                    'font-normal',
                    'bg-gray-50',
                    'hover:bg-gray-100',
                    'disabled:bg-gray-50',
                    'text-gray-950',
                    'shadow',
                    'border',
                    'border-neutral-200/50',
                    'data-[focus-visible]:ring-gray-200',
                    'data-[pressed=true]:scale-[0.98]',
                    'data-[focus-visible]:ring-2',
                    'data-[focus-visible]:ring-offset-2',
                ],
                limited:[
                    'font-semibold',
                    'bg-red-400',
                    'hover:bg-red-500',
                    'text-white',
                    'rounded-full',
                    'shadow',
                    'hover:shadow-md',
                    'disabled:bg-red-400/50',
                    'disabled:shadow',
                    'data-[focus-visible]:ring-red-400',
                    'data-[pressed=true]:scale-[0.98]',
                    'data-[focus-visible]:ring-2',
                    'data-[focus-visible]:ring-offset-2',
                ],
                nobackground:[
                    'font-light',
                    'text-gray-950',
                    'hover:text-gray-500',
                    'disabled:text-gray-950',
                    'data-[focus-visible]:ring-gray-400/30',
                    'data-[focus-visible]:ring-1',
                ],
                link:[
                    'font-light',
                    'text-green-400',
                    'hover:text-green-500',
                    'disabled:text-green-400/50',
                    'disabled:no-underline',
                    'hover:underline',
                    'data-[focus-visible]:ring-green-400/30',
                    'data-[focus-visible]:ring-1',
                ],
            },
            size: {
                small: ['text-sm', 'py-1', 'px-4'],
                default: ['text-base', 'py-2', 'px-8'],
                large: ['text-lg', 'py-3', 'px-12'],
              }
        },
        defaultVariants: {
            variant: 'primary',
            size: 'default'
        },
    }
)

export const loading = cva(['absolute','inline-flex','items-center'],
{
    variants:{
        variant:{
            primary: ['border-white'],
            secondary: ['border-gray-950'],
            limited: ['border-white'],
            nobackground: ['border-gray-950'],
            link: ['border-green-400'],
        }
    }
})





export const Loading = ({variant}: VariantProps<typeof loading>)=>(
    <div className={loading({variant})}>
        <div className="w-4 h-4 rounded-full border-2 border-b-transparent animate-spin border-[inherit]"></div>
    </div>)


export type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & 
                    AriaButtonProps<'button'> & 
                    VariantProps<typeof variants> &  {
                      loading?:boolean,
                      icon?:any
                    }

export type DropdownProps = React.SelectHTMLAttributes<HTMLSelectElement> & 
AriaSelectProps<'select'> & 
VariantProps<typeof variants> &  {
    loading?:boolean,
    icon?:any
}