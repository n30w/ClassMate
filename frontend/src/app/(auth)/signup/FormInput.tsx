import React from "react";

interface FormInputProps {
  label: string;
  type: string;
  name: string;
  placeholder: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  errorMessage?: string;
}

const FormInput: React.FC<FormInputProps> = ({
  label,
  type,
  name,
  placeholder,
  value,
  onChange,
  errorMessage,
}) => {
  return (
    <div>
      <label className="text-white font-light py-2" htmlFor={name}>
        {label}
      </label>
      <input
        type={type}
        id={name}
        name={name}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
        className="w-80 h-10 px-4 mb-8"
        required
      />
      {errorMessage && (
        <p data-testid="errorMessage" className="text-red-500 pb-2">
          {errorMessage}
        </p>
      )}
    </div>
  );
};

export default FormInput;
