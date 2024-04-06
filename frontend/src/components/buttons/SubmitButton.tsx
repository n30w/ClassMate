"use client";

import { useFormStatus } from "react-dom";

interface props {
  text: string;
  className: string;
}

const SubmitButton: React.FC<props> = (props: props) => {
  const { pending } = useFormStatus();

  return (
    <button type="submit" className={props.className} aria-disabled={pending}>
      {props.text}
    </button>
  );
};

export default SubmitButton;
