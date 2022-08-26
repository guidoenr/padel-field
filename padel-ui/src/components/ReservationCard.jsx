import React from "react";
import ReservationItem from "./ReservationItem";
import Weather from "./Weather";

const ReservationCard = () => {
  return (
    <article className="cardd w-[90%] max-w-md relative rounded-lg mb-12 border border-primary/30 shadow-lg">
      <div className="container p-4">
        <Weather />
        <div className="card-header text-center">
          <p className="card__title font-medium text-xl py-2">LUNES - 22/08</p>
        </div>
        <div className="card-content">
          <ul className="reservation-list">
            <ReservationItem />
            <ReservationItem />
            <ReservationItem />
            <ReservationItem />
            <ReservationItem />
            <ReservationItem />
          </ul>
        </div>
      </div>
    </article>
  );
};

export default ReservationCard;
