/*
* AddButton defines a button that adds content to either the screen or the the backend API.
* The button takes an optional text prop.
* */
import React from "react";
import {AddIcon} from "@sanity/icons";

interface ButtonProps {
    text: string
}

const AddButton: React.FC<ButtonProps>  = (props: ButtonProps) => {
    return (
        <button
            className="flex space-x-1 items-center bg-white transition-all ease-out duration-75 hover:bg-green-300 text-black font-light px-4 py-2 h-12"
        >
            <AddIcon className={""}/>
            <p>{props.text}</p>
        </button>
    )
}

export default AddButton;