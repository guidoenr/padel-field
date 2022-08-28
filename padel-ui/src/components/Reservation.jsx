import React from "react";
import ReservationCard from "./ReservationCard";
import { Link, animateScroll as scroll } from "react-scroll";

const Reservation = () => {
  return (
    <main className="reservation w-full flex flex-col items-center">
      <div className="introimg hero h-[35rem] w-full mb-10 relative">
        <div className="hero-overlay bg-opacity-60 bg-[#151515]"></div>
        <div className="hero-content text-center text-neutral-content">
          <div className="max-w-md text-center flex flex-col items-center">
            <h2 className="text-[#e5e0df] w-full mb-20 text-4xl lg:text-5xl">
              Reserva tu turno
            </h2>
            <Link
              to="reservation"
              smooth={true}
              offset={-50}
              duration={700}
              className="text-[#e5e0df] w-20 h-20 animate-bounce absolute bottom-10 cursor-pointer hover:text-accent hover:scale-105 transition ease-in-out"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M19.5 5.25l-7.5 7.5-7.5-7.5m15 6l-7.5 7.5-7.5-7.5"
                />
              </svg>
            </Link>
          </div>
        </div>
      </div>

      <div
        className="container mx-auto flex flex-col items-center gap-4 py-10 xl:last:col-start-2 xl:last:col-end-3  md:grid md:grid-cols-2 md:place-items-center xl:grid-cols-3"
        name="reservation"
      >
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
