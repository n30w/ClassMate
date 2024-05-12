import formattedDate from "@/lib/helpers/formattedDate";
import React from "react";

/**
 * @prop date string in RFC3339 Format
 * @prop colorClass string in the format bg-[color]-[value]
 */
interface DateBadgeProps {
  text: string;
  colorClass?: string;
}

const InfoBadge: React.FC<DateBadgeProps> = ({
  text,
  colorClass,
}: DateBadgeProps) => {
  return (
    <div
      className={`mb-2 rounded-2xl w-fit ${colorClass ? colorClass : "bg-yellow-700"}`}
    >
      <p className="text-xs tracking-wide text-white px-2 py-1">{text}</p>
    </div>
  );
};

export default InfoBadge;
