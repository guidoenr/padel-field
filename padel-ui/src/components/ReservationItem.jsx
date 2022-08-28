import React from "react";
import { BiTennisBall } from "react-icons/bi";
import { useState } from "react";

const ReservationItem = () => {
  const [available, setAvailable] = useState(true);

  const [recerved, setRecerved] = useState(false);

  const showReserved = () => {
    setAvailable(false);
    setRecerved(true);
  };

  return (
    <li className="reservation-item flex justify-between items-center border-b border-b-primary/30 font-medium py-1 last-of-type:border-none">
      <p className="reservation-item-schedule text-primary text-lg">13:00</p>
      {available && (
        <button
          onClick={showReserved}
          className="reservation-btn-container reservation-btn px-3 border-none text-sm text-primary transition pl-10 relative flex items-center btn hover:bg-[#4bbb90] hover:scale-105 ease-in-out bg-[#3faa81]"
        >
          <BiTennisBall className="absolute left-0 w-8 p-1 rounded-l-md text-2xl cursor-pointer h-full bg-[#17382b] text-[#46bb8e]" />
          Reservar
        </button>
      )}
      {recerved && (
        <button className="reservation-btn-container reservation-btn px-3 border-none text-sm text-primary hover:bg-[#911b2afd] pl-10 relative flex items-center btn cursor-not-allowed bg-[#911b2afd]">
          <BiTennisBall className="cursor-not-allowed absolute left-0 w-8 p-1 rounded-l-lg text-2xl h-full bg-[#440713] text-[#c4032ac2]" />
          Reservado
        </button>
      )}
    </li>
  );
};

export default ReservationItem;
