import React from "react";
import ReservationItem from "./ReservationItem";
import Weather from "./Weather";

const ReservationCard = ({ day }) => {
  return (
    <article className="cardd w-[90%] max-w-sm md:max-w-xs relative rounded-lg mb-20 border border-primary/30 bg-neutral/80 shadow-lg">
      <div className="container p-4">
        <Weather />
        <div className="card-header text-center">
          <p className="card__title font-medium text-xl xl:text-[1.35rem] py-2 pb-4">
            {day}
          </p>
        </div>
        <div className="card-content">
          <ul className="reservation-list">
            
            <ReservationItem hr="9:00" />
            <ReservationItem hr="10:00" />
            <ReservationItem hr="11:00" />
            <ReservationItem hr="12:00" />
            <ReservationItem hr="13:00" />
            <ReservationItem hr="14:00" />
            <ReservationItem hr="15:00" />
            <ReservationItem hr="16:00" />
            <ReservationItem hr="17:00" />
            <ReservationItem hr="18:00" />
            <ReservationItem hr="19:00" />
            <ReservationItem hr="20:00" />
            <ReservationItem hr="21:00" />
            <ReservationItem hr="22:00" />
            <ReservationItem hr="23:00" />
          </ul>
        </div>
      </div>
    </article>
  );
};

export default ReservationCard;
