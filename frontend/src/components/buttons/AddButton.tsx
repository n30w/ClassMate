/*
 * AddButton defines a button that adds content to either the screen or the the backend API.
 * The button takes an optional text prop.
 * */
import React from "react";
import { AddIcon } from "@sanity/icons";

interface ButtonProps {
  text?: string;
  fullWidth?: boolean;
  onClick: () => void;
}

const AddButton: React.FC<ButtonProps> = ({
  text,
  fullWidth,
  onClick,
}: ButtonProps) => {
  return (
    <button
      className={`flex space-x-1 items-center border-2 border-dashed shadow bg-gray-800 border-gray-700 transition-all ease-in-out duration-75 hover:bg-gray-500 text-white px-4 py-2 h-12 ${fullWidth ? "w-full justify-center" : ""}`}
      onClick={onClick}
    >
      <AddIcon className="text-xl" />
      {text && (
        <p className="text-sm font-semibold tracking-wide">
          {text.toUpperCase()}
        </p>
      )}
    </button>
  );
};

export default AddButton;
