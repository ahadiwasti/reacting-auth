'use client'
import React, { forwardRef, useRef } from 'react';
import {  mergeRefs } from '@/lib'
import {mergeProps,useFocusRing, useHover, useTextField} from 'react-aria';
import {InputProps, Loading, variants} from '../cvas/fields'
import { twMerge } from 'tailwind-merge';
import clsx from 'clsx';
const Input = forwardRef<HTMLInputElement, InputProps>((props, forwardedRef) => {
    const {
      className,
      variant,
      size,
      loading,
      disabled,
      icon,
      name,
      description,
      errorMessage,
      ...rest
    } = props;
  
    const ref = useRef<HTMLInputElement>(null);
    const { focusProps, isFocusVisible } = useFocusRing();
    const { hoverProps, isHovered } = useHover({
      ...props,
      isDisabled: disabled,
    });
  
  
    const { inputProps, labelProps, descriptionProps, errorMessageProps } = useTextField(
      {
    
      ...rest,
      isDisabled: disabled,
      },
      ref
    );
  
    return (
      <>
    <div>
    <div>
        {loading && <Loading variant={variant} />}
        </div>
    <div>
        
        <div>
            <label className="font-bold text-green-400 font-sans" {...labelProps}>{name} :</label>
        </div>
        <div className="grid grid-cols-2 gap-4 text-center self-center items-center content-center justify-center justify-items-center place-self-center place-content-center place-items-center ">
        <div>
        <input
          ref={mergeRefs([ref, forwardedRef])}
          className={twMerge(clsx(variants({ variant, size, className })))}
          {...mergeProps(inputProps, focusProps, hoverProps, labelProps, descriptionProps, errorMessageProps)}
      
          {...rest}
          data-hovered={isHovered || undefined}
          data-focus-visible={isFocusVisible || undefined}
        />
        </div>
        <div className="text-center h-full self-center items-center content-center justify-center justify-items-center place-self-center place-content-center place-items-center">
        {icon && (
              <span className="h-full text-green-400 text-center self-center items-center content-center justify-center justify-items-center place-self-center place-content-center place-items-center">
                {icon}
              </span>
            )}
        </div>
        </div>
       
        <div>
        {props.description && <div {...descriptionProps}>{props.description}</div>}
        </div>
        <div>
        {props.errorMessage && <div {...errorMessageProps}>{props.errorMessage}</div>}
        </div>
      </div>
    </div>
     
        
        
            
        
  
     
      
      </>
    );
  });
  




Input.displayName = 'Input'
export { Input }
export type { InputProps }