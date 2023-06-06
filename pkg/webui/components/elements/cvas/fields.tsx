'use client'

import {VariantProps, cva} from 'class-variance-authority'
import { SelectHTMLAttributes } from 'react'
import {  AriaTextFieldOptions,AriaComboBoxOptions } from 'react-aria'


export const variants = cva(
    [

         'px-4 py-2', 'rounded-md', 'border border-green-200', 
        'focus:outline-none focus:ring-2 focus:ring-green-400/70',
        'shadow',
        'font-semibold',
        'text-green-400'
        
    ],
    {
        variants:{
            variant:{
                primary:[
                    'disabled:bg-gray-400/50',
                    'disabled:border-gray-300',
                    'disabled:cursor-not-allowed',
                    'disabled:shadow',
                    'data-[focus-visible]:ring-green-400/70',
                    'data-[pressed=true]:scale-[0.98]',
                    'data-[focus-visible]:ring-1',
                    'data-[focus-visible]:ring-offset-2'
                ],
                secondary:[ 'shadow'],
                limited:[ 'shadow'],
                nobackground:[ 'shadow'],
            },
            size:{
                small:['w-small'],
                default:['w-full'],
                large:['w-full']
            }
        },
        defaultVariants:{
            variant:'primary',
            size:'default'
        }
    }
)


export const loading = cva(
    ['absolute', 'inline-flex','item-center'],{
    variants:{
        variant:{
            primary: ['border-white'],
            secondary: ['border-gray-950'],
            limited: ['border-white'],
            nobackground: ['border-gray-950'],
            link: ['border-green-400'],
        }
    }}
)

export const Loading = ({variant}:VariantProps<typeof loading>)=>(
    <div className={loading({variant})}>
         <div className="w-4 h-4 rounded-full border-2 border-b-transparent animate-spin border-[inherit]"></div>
    </div>
)


export type InputProps =  React.InputHTMLAttributes<HTMLInputElement> & 
AriaTextFieldOptions<"input"> &
VariantProps<typeof variants> & {
    loading?:boolean,
    icon?:any,
    label?: string,
    description?: string,
    errorMessage?:string,
    onFocus?: (event: React.FocusEvent<HTMLInputElement>) => void,
    onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void,
}


export type ComboboxProps = React.SelectHTMLAttributes<HTMLSelectElement> & 
AriaComboBoxOptions<any> &
VariantProps<typeof variants> & {
    loading?:boolean,
    icon?:any,
    disabled?:boolean,
    className?:string,
    label?: string,
    description?: string,
    errorMessage?:string,
    onFocus?: (event: React.FocusEvent<HTMLInputElement>) => void,
    onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void,
}