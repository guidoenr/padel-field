import React from "react";
import ReservationCard from "./ReservationCard";

const Reservation = () => {
  return (
    <main className="reservation w-full flex flex-col items-center">
      <div className="introimg hero h-[30rem] w-full mb-10 relative">
        <div className="hero-overlay bg-opacity-60 bg-[#151515]"></div>
        <div className="hero-content text-center text-neutral-content">
          <div className="max-w-md text-center flex flex-col items-center">
            <h2 className="text-[#e5e0df] w-full mb-20 text-4xl lg:text-5xl">
              Reserva tu turno
            </h2>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="text-[#e5e0df] w-20 h-20 animate-bounce absolute bottom-10"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M19.5 5.25l-7.5 7.5-7.5-7.5m15 6l-7.5 7.5-7.5-7.5"
              />
            </svg>
          </div>
        </div>
      </div>

      <div className="container mx-auto flex flex-col items-center gap-4 py-10 lg:grid lg:grid-cols-2 lg:place-items-center xl:lg:grid-cols-3">
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
        <ReservationCard />
      </div>
    </main>
  );
};

export default Reservation;
