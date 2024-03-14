/*
* CloseButton is a reusable close button.
* */
import React from "react";
import {CloseIcon} from "@sanity/icons";

interface ButtonProps {
    text?: string
    onClick: () => void
}

const CloseButton: React.FC<ButtonProps>  = (props: ButtonProps) => {
    return (
        <button
            className="absolute top-0 right-0 m-2 text-black text-lg font-bold cursor-pointer"
            onClick={props.onClick}
        >
            <CloseIcon style={{ color: `white`, scale: `1.5`}} />
            {props.text && <p>{props.text}</p>}
        </button>
    )
}

export default CloseButton;