import React from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import { TbMoodEmpty } from "react-icons/tb";

const MyTurns = () => {
  const [emptyTurns, setEmptyTurns] = useState(true);
  const [nextTurns, setNextTurns] = useState(false);

  return (
    <section className="myturns h-[80vh] py-8 w-full">
      <div className="myturns-container container mx-auto px-1 flex flex-col">
        <div className="myturns-header relative py-4">
          <h3 className="text-2xl lg:text-3xl font-medium">
            Tus proximos turnos:
          </h3>
          <div className="absolute bottom-1 w-[3rem] h-[.25rem] bg-accent"></div>
          <div className="absolute bottom-[-.25rem] w-[1.5rem] h-[.25rem] bg-accent"></div>
        </div>
        {emptyTurns && (
          <div className="text-lg empty-turns pt-8 flex flex-col items-center gap-2 my-6">
            <p className="flex flex-col gap-3 max-w-md w-[95%] text-center mx-auto py-6 text-lg rounded-md border border-primary/30 bg-primary/10">
              <span className="w-full flex justify-center">
                <TbMoodEmpty className="text-[8rem]" />
              </span>
              No tenes ningún turno reservado
            </p>
            <Link
              to="/"
              className="font-medium my-2 btn w-[95%] max-w-md mx-auto normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out sm:w-[7rem]"
            >
              Reservar
            </Link>
          </div>
        )}
        {nextTurns && (
          <div className="text-lg next-turns pt-8 text-center text-secondary">
            <p className="flex flex-col max-w-md w-[95%] mx-auto py-4 my-2 text-base rounded-md border border-primary/30 bg-[#d19c63]">
              Tu próximo turno será el día{" "}
              <span className="">Martes 30/08 a las 19:00</span>
            </p>
          </div>
        )}
      </div>
    </section>
  );
};

export default MyTurns;
