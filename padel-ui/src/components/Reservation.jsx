import React from "react";
import ReservationCard from "./ReservationCard";

const Reservation = () => {
  return (
    <section className="reservation w-full flex flex-col items-center">
      <div className="introimg h-[20rem] w-full mb-10"></div>
      <h2 className="mb-14 text-3xl">Reserva tu turno:</h2>
      <div className="container mx-auto flex flex-col items-center gap-4 py-4 lg:grid lg:grid-cols-2 lg:place-items-center xl:lg:grid-cols-3">
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
      </div>
    </section>
  );
};

export default Reservation;
