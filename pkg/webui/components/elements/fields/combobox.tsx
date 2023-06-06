import * as React from "react";
import type { ComboBoxProps } from "@react-types/combobox";
import { useComboBoxState } from "react-stately";
import { useComboBox, useFilter, useButton } from "react-aria";
import { ChevronDownIcon } from "@heroicons/react/20/solid";
import { VariantProps } from 'class-variance-authority';

import { ListBox } from "./ListBox";
import { Popover } from "./Popover";

export { Item, Section } from "react-stately";

type ComboBoxHTMLProps = React.HTMLAttributes<HTMLDivElement>;

export function ComboBox<T extends object>(props: ComboBoxProps<T> &  ComboBoxHTMLProps) {
  let { contains } = useFilter({ sensitivity: "base" });
  let state = useComboBoxState({ ...props, defaultFilter: contains });

  let buttonRef = React.useRef(null);
  let inputRef = React.useRef(null);
  let listBoxRef = React.useRef(null);
  let popoverRef = React.useRef(null);

  let {
    buttonProps: triggerProps,
    inputProps,
    listBoxProps,
    labelProps,
  } = useComboBox(
    {
      ...props,
      inputRef,
      buttonRef,
      listBoxRef,
      popoverRef,
    },
    state
  );

  let { buttonProps } = useButton(triggerProps, buttonRef);

  const variantClassNames = {
    default: "border-gray-300",
    pink: "border-pink-500",
    // Add more variant class names as needed
  };

  const variant = props.variant || "default";
  const variantClassName = variantClassNames[variant] || variantClassNames.default;

  return (
    <div className={`inline-flex flex-col relative w-52 ${props.className}`}>
      <label
        {...labelProps}
        className="block text-sm font-medium text-gray-700 text-left"
      >
        {props.label}
      </label>
      <div
        className={`relative flex inline-flex flex-row rounded-md overflow-hidden shadow-sm border-2 ${variantClassName} ${
          state.isFocused ? "border-pink-500" : "border-gray-300"
        }`}
      >
        <input
          {...inputProps}
          ref={inputRef}
          className="outline-none px-3 py-1 w-full"
        />
        <button
          {...buttonProps}
          ref={buttonRef}
          className={`px-1 bg-gray-100 cursor-default border-l-2 ${variantClassName} ${
            state.isFocused
              ? "border-pink-500 text-pink-600"
              : "border-gray-300 text-gray-500"
          }`}
        >
          <ChevronDownIcon className="w-5 h-5" aria-hidden="true" />
        </button>
      </div>
      {state.isOpen && (
        <Popover
          popoverRef={popoverRef}
          triggerRef={inputRef}
          state={state}
          isNonModal
          placement="bottom start"
          className="w-52"
        >
          <ListBox {...listBoxProps} listBoxRef={listBoxRef} state={state} />
        </Popover>
      )}
    </div>
  );
}
