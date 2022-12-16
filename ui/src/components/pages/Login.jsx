import React, { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Link, Navigate } from "react-router-dom";
import { motion } from "framer-motion";
import FormInput from "../FormInput";

const Login = () => {
  const [redirect, setRedirect] = useState(false);
  const [isPending, setIsPending] = useState(false);

  const [values, setValues] = useState({
    username: "",
    password: "",
  });

  const inputs = [
    {
      id: 1,
      name: "username",
      label: "Nombre de Usuario:",
      type: "text",
      errorMessage: "Usuario no existe",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
    {
      id: 2,
      name: "password",
      label: "Contraseña:",
      type: "password",
      errorMessage: "Contraseña incorrecta",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
  ];

  const onChange = (e) => {
    setValues({ ...values, [e.target.name]: e.target.value });
  };

  const submit = async (e) => {
    e.preventDefault();
    setIsPending(true);
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify(values),
    });
    if (response.ok) {
      setIsPending(false);
      console.log(response);
      console.log(values);
      setRedirect(true);
    } else {
      switch (response.status) {
        case 400:
        // no existe el usuario
        case 401:
        // mala password
        case 406:
        // mal email
      }
    }
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
    <section className="login w-full h-[90vh] flex items-center">
      {redirect && <Navigate replace to="/" />}
      <motion.div
        variants={containerVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
        className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container pt-12 lg:pt-8 w-[95%] lg:h-[90%] m-auto p-8 lg:flex lg:justify-between lg:items-center lg:max-w-7xl relative"
      >
        <div className="form-container flex flex-col gap-2 lg:w-[50%]">
          <div className="relative bottom-5 lg:absolute lg:top-8 lg:left-8 pb-2">
            <span className="select-none font-secondary-font text-4xl text-primary">
              Pádel-Logo
            </span>
          </div>
          <h3 className="form-title text-2xl font-semibold">Iniciar sesion</h3>
          <p className="form-text text-sm text-primary/60">
            Entra a tu cuenta para poder reservar tu turno.
          </p>
          <form className="flex flex-col gap-3" onSubmit={submit}>
            {inputs.map((input) => (
              <FormInput
                key={input.id}
                {...input}
                value={values[input.name]}
                onChange={onChange}
              />
            ))}
            <div className="row pt-4">
              {!isPending && (
                <button
                  type="submit"
                  className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
                >
                  Entrar
                </button>
              )}
              {isPending && (
                <button
                  type="submit"
                  className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
                >
                  Entrando...
                </button>
              )}
            </div>
            <div className="row">
              <button
                type="button"
                className="btn w-full bg-transparent flex gap-2 normal-case text-primary border-primary hover:border-primary hover:bg-primary/10 hover:scale-105 transition ease-in-out"
              >
                <FcGoogle className="text-2xl" />
                Entrar con Google
              </button>
            </div>
          </form>
          <div className="log-in-container mt-8 flex flex-col gap-4">
            <p className="log-in-text text-xl font-semibold">
              ¿No tenes cuenta?
            </p>
            <Link
              to="/register"
              className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
            >
              Crear cuenta
            </Link>
          </div>
        </div>
        <div className="login-img-container hidden lg:flex w-[45%] h-[40rem] rounded-lg">
          <div className="hero-overlay bg-opacity-30 bg-[#151515] rounded-lg"></div>
        </div>
      </motion.div>
    </section>
  );
};

export default Login;
