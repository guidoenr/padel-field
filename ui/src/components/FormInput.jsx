import React from "react";
import { BiErrorAlt } from "react-icons/bi";
import { useState } from "react";

const FormInput = (props) => {
  const [focused, setFocused] = useState(false);

  const { label, errorMessage, onChange, id, name, type, ...inputProps } =
    props;

  const handleFocus = (e) => {
    setFocused(true);
  };

  return (
    <div className="row form-input flex flex-col gap-1">
      <label htmlFor={name} className="text-primary/60 font-semibold text-sm">
        {label}
      </label>
      <input
        {...inputProps}
        type={type}
        id={id}
        name={name}
        onChange={onChange}
        className="focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
        onBlur={handleFocus}
        focused={focused.toString()}
      />
      <p className="error-msj text-red-700 text-sm">
        <span className="">
          <BiErrorAlt className="inline text-lg relative bottom-[.1rem]" />{" "}
          {errorMessage}
        </span>
      </p>
    </div>
  );
};

export default FormInput;
