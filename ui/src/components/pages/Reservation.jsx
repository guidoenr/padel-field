import React from "react";
import ReservationCard from "../ReservationCard";
import { Link, animateScroll as scroll } from "react-scroll";
// import { useEffect } from "react";

const Reservation = () => {
  /*useEffect(() => {
    (async () => {
      // fetch to user url
      await fetch("http://localhost:8080/auth/user", {
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    })();
  }); */

  return (
    <main className="reservation w-full flex flex-col items-center">
      <div className="introimg hero h-[35rem] w-full mb-10 relative">
        <div className="hero-overlay bg-opacity-60 bg-[#151515]"></div>
        <div className="hero-content text-center text-neutral-content">
          <div className="max-w-md text-center flex flex-col items-center gap-6">
            <h2 className="text-[#e5e0df] w-full text-4xl lg:text-5xl">
              Reserva tu turno
            </h2>
            <p className="text-[#e5e0df] text-2xl w-full mb-20 lg:text-3xl">
              $300/hr
            </p>
            <Link
              to="reservation"
              smooth={true}
              offset={-50}
              duration={700}
              className=" text-[#e5e0df] w-20 h-20 animate-bounce absolute bottom-10 cursor-pointer hover:text-accent/70 hover:scale-105 transition ease-in-out"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className=""
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
        className="container mx-auto max-w-[95%] flex flex-col items-center py-10 sm:grid sm:grid-cols-2 sm:place-items-center lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 first:bg-red-500"
        name="reservation"
      >
        <ReservationCard day="LUNES - 22/08" />
        <ReservationCard day="MARTES - 23/08" />
        <ReservationCard day="MIERCOLES - 24/08" />
        <ReservationCard day="JUEVES - 25/08" />
        <ReservationCard day="VIERNES - 26/08" />
        <ReservationCard day="SABADO - 27/08" />
        <ReservationCard day="DOMINGO - 28/08" />
      </div>
    </main>
  );
};

export default Reservation;
