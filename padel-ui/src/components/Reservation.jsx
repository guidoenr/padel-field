import React from "react";
import ReservationCard from "./ReservationCard";
import { BsArrowDownShort } from "react-icons/bs";

const Reservation = () => {
  return (
    <section className="reservation w-full flex flex-col items-center">
      <div className="introimg hero h-[20rem] w-full mb-10 relative">
        <div className="hero-overlay bg-opacity-50 bg-[#151515]"></div>
        <div className="hero-content text-center text-neutral-content">
          <div className="max-w-md text-center flex flex-col items-center">
            <h2 className="text-3xl text-[#e5e0df] mb-12 xl:text-4xl">
              Reserva tu turno
            </h2>
            <BsArrowDownShort className="text-[#e5e0df] text-8xl animate-bounce absolute bottom-5" />
          </div>
        </div>
      </div>

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
