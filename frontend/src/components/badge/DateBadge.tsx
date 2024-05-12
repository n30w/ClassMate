import formattedDate from "@/lib/helpers/formattedDate";
import React from "react";

/**
 * @prop date string in RFC3339 Format
 * @prop colorClass string in the format bg-[color]-[value]
 */
interface DateBadgeProps {
  date: string;
  colorClass?: string;
}

const DateBadge: React.FC<DateBadgeProps> = ({
  date,
  colorClass,
}: DateBadgeProps) => {
  return (
    <div
      className={`mb-2 rounded-xl w-fit ${colorClass ? colorClass : "bg-yellow-700"}`}
    >
      <h6 className="text-xs tracking-wide text-white px-2 py-1">
        {formattedDate(date).toLocaleUpperCase()}
      </h6>
    </div>
  );
};

export default DateBadge;
