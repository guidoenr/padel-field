import React from "react";
import { BiTennisBall } from "react-icons/bi";
import { useState } from "react";
import { Link } from "react-router-dom";

const ReservationItem = ({ hr }) => {
  const [available, setAvailable] = useState(true);
  const [recerved, setRecerved] = useState(false);

  const showReserved = () => {
    setAvailable(false);
    setRecerved(true);
  };

  return (
    <li className="reservation-item flex justify-between items-center border-b border-b-primary/30 font-medium py-1 last-of-type:border-none">
      <p className="reservation-item-schedule text-primary text-lg">{hr}</p>
      {available && (
        <Link
          to="confirmreservation"
          onClick={showReserved}
          className="h-[2rem] rounded-lg reservation-btn-container reservation-btn px-3 text-[.9rem] border-none text-primary transition pl-10 relative flex items-center hover:bg-[#4bbb90] hover:scale-105 ease-in-out bg-[#3faa81]"
        >
          <BiTennisBall className="absolute left-0 w-8 p-1 rounded-l-md text-2xl cursor-pointer h-full bg-[#17382b] text-[#46bb8e]" />
          Reservar
        </Link>
      )}
      {recerved && (
        <button className="h-[2rem] rounded-lg reservation-btn-container reservation-btn px-3 border-none text-sm text-primary hover:bg-[#911b2afd] pl-10 relative flex items-center cursor-not-allowed bg-[#911b2afd]">
          <BiTennisBall className="cursor-not-allowed absolute left-0 w-8 p-1 rounded-l-lg text-2xl h-full bg-[#440713] text-[#c4032ac2]" />
          Reservado
        </button>
      )}
    </li>
  );
};

export default ReservationItem;
