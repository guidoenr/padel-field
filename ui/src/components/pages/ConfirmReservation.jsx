import React from "react";
import { useState } from "react";
import { BiTimeFive, BiDollar } from "react-icons/bi";
import { motion } from "framer-motion";
import { NavLink, Link } from "react-router-dom";

const ConfirmReservation = ({
  day = "23/10",
  hr = "20:00",
  price = 300,
  handleReservedTurns,
}) => {
  const [hour, setHour] = useState("1 hora");

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
    <section className="login w-full h-[80vh] flex items-center">
      <motion.div
        variants={containerVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
        className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container w-[95%] max-w-lg mx-auto p-8 relative"
      >
        <div className="form-container flex flex-col gap-2">
          <h3 className="form-title text-2xl font-semibold text-center">
            Confirmar Turno
          </h3>
          <h2 className="flex flex-col max-w-md py-4 text-lg border-b border-b-accent font-medium text-center">
            Día: {day} - Horario: {hr}
          </h2>
          <form className="flex flex-col gap-7 mt-6">
            <div className="row flex gap-2 items-center">
              <BiTimeFive className="text-2xl text-accent" />
              <label htmlFor="" className="font-medium">
                Duración:
              </label>

              <select
                value={hour}
                onChange={(e) => setHour(e.target.value)}
                className="select select-accent w-auto select-sm"
              >
                <option value="1 hora">1 hora</option>
                <option value="2 horas">2 horas</option>
                <option value="3 horas">3 horas</option>
              </select>
            </div>
            <div className="row">
              <p className="flex items-center gap-2 font-medium">
                <BiDollar className="text-2xl text-accent" />
                Precio por persona:
                <span className="text-lg">
                  $
                  {hour === "2 horas"
                    ? price * 2
                    : hour === "3 horas"
                    ? price * 3
                    : price}
                </span>
              </p>
            </div>
            <div className="row pt-4">
              <NavLink
                to="/misturnos"
                className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
                onClick={handleReservedTurns}
              >
                Confirmar
              </NavLink>
            </div>
          </form>
        </div>
      </motion.div>
    </section>
  );
};

export default ConfirmReservation;
