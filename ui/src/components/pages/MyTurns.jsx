import React from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import { TbMoodEmpty } from "react-icons/tb";
import { motion } from "framer-motion";

const MyTurns = () => {
  const [emptyTurns, setEmptyTurns] = useState(true);
  const [nextTurns, setNextTurns] = useState(false);

  const handleReservedTurns = () => {
    setEmptyTurns(false);
    setNextTurns(true);
  };

  const containerVariants = {
    hidden: {
      opacity: 0,
      x: "100vw",
    },
    visible: {
      opacity: 1,
      x: 0,
      transition: {
        type: "spring",
        bounce: 0.25,
      },
    },
    exit: {
      x: "-100vw",
      transition: {
        type: "spring",
        bounce: 0.25,
      },
    },
  };

  return (
    <section className="myturns h-[80vh] py-8 w-full">
      <motion.div
        className="myturns-container container mx-auto px-1 flex flex-col"
        variants={containerVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
      >
        <div className="myturns-header relative py-4">
          <h3 className="text-2xl lg:text-3xl font-medium">
            Mis próximos turnos:
          </h3>
          <div className="absolute bottom-1 w-[3rem] h-[.25rem] bg-accent"></div>
          <div className="absolute bottom-[-.25rem] w-[1.5rem] h-[.25rem] bg-accent"></div>
        </div>
        <button
          className=" bg-amber-500 font-medium border text-secondary w-[em] mx-auto rounded-md text-sm p-2"
          onClick={handleReservedTurns}
        >
          Cambiar estado
        </button>
        {emptyTurns && (
          <div className="text-lg empty-turns pt-8 flex flex-col items-center gap-2 my-6 ">
            <p className="flex flex-col gap-3 max-w-md w-[95%] text-center mx-auto py-6 text-lg ">
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
      </motion.div>
    </section>
  );
};

export default MyTurns;
