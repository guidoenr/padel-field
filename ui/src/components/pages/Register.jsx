import React, { SyntheticEvent, useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Navigate } from "react-router-dom";
import { motion } from "framer-motion";
import FormInput from "../FormInput";

const Register = () => {
  const [isPending, setIsPending] = useState(false);
  const [redirect, setRedirect] = useState(false);
  // form validation

  const [values, setValues] = useState({
    name: "",
    surname: "",
    username: "",
    password: "",
    email: "",
    phone: "",
  });

  const inputs = [
    {
      id: 1,
      name: "name",
      label: "Nombre:",
      type: "text",
      errorMessage: "Tu nombre debe contener entre 3 y 16 caracteres",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
    {
      id: 2,
      name: "surname",
      label: "Apellido:",
      type: "text",
      errorMessage: "Tu apellido debe contener entre 3 y 16 caracteres",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
    {
      id: 3,
      name: "username",
      label: "Nombre de Usuario:",
      type: "text",
      errorMessage:
        "Tu nombre de usuario debe contener entre 3 y 16 caracteres",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
    {
      id: 4,
      name: "password",
      label: "Contraseña:",
      type: "password",
      errorMessage: "Tu contraseña debe contener entre 3 y 16 caracteres",
      required: true,
      pattern: "^[A-Za-z0-9]{3,16}$",
    },
    {
      id: 5,
      name: "email",
      label: "Email:",
      type: "email",
      errorMessage: "Debe ser un email válido",
      required: true,
    },
    {
      id: 6,
      name: "phone",
      label: "Celular:",
      type: "tel",
      errorMessage: "Debe ser un número celular que exista.",
      required: true, // Check regex
    },
  ];

  const onChange = (e) => {
    setValues({ ...values, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: SyntheticEvent) => {
    const url = "http://localhost:8080/auth/register"; // ""
    e.preventDefault();
    setIsPending(true);
    const response = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(values),
    });

    if (response.ok) {
      setIsPending(false);
      console.log(response);
      console.log(values);
      setRedirect(true);
    } else {
      switch (response.status){
        case 409:
          // el username ya existe
          // @ marcos ,hace lo que quieras aca, mostrale lo que se te cante al usuario
        case 406:
          // el email ya existe
      }
    }
  };
  if (redirect) {
    // @marcos, se puede aca redireccionarlo al /home ya con la cuenta logeada?
    // osea se que se puede si le meto toda la funcion del pegado a la api que esta en login
    // pero hay alguna forma de redireccionarlo con los parametros que lleno aca? para que sea mas lindo obvio
    // entiendo que si hacemos que la funcion reciba parametros si, sino decime vos
    return <Navigate to="/"> </Navigate>;
  }

  // framer motion

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

  /*

  */
  return (
    <section className="register w-full h-[100vh] lg:h-[90vh] lg:flex lg:items-center">
      <motion.div
        variants={containerVariants}
        initial="hidden"
        animate="visible"
        exit="exit"
        className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container mt-8 w-[95%] lg:mt-0 lg:w-[90%] lg:h-[90%] mx-auto p-8 lg:flex lg:justify-between lg:items-center lg:max-w-7xl relative"
      >
        <div className="form-container flex flex-col gap-2 lg:w-[50%] lg:order-2">
          <div className="hidden lg:flex lg:absolute lg:top-8 lg:right-8">
            <a
              className="cursor-pointer font-secondary-font text-4xl text-primary"
              href="#home"
            >
              Pádel-Logo
            </a>
          </div>
          <h3 className="form-title text-2xl font-semibold">Crear cuenta</h3>
          <p className="form-text text-sm text-primary/60">
            Una vez que tengas tu cuenta, podrás reservar tu turno.
          </p>
          <form className="flex flex-col gap-3" onSubmit={handleSubmit}>
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
                  Crear cuenta
                </button>
              )}
              {isPending && (
                <button
                  disabled
                  className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
                >
                  Creando Cuenta...
                </button>
              )}
            </div>
            <div className="row">
              <button
                type="button"
                className="btn w-full bg-transparent flex gap-1 normal-case text-primary border-primary hover:border-primary hover:bg-primary/10 hover:scale-105 transition ease-in-out"
              >
                <FcGoogle className="text-2xl" />
                Crear cuenta con Google
              </button>
            </div>
          </form>
          <div className="log-in-container mt-8 flex flex-col gap-4">
            <p className="log-in-text text-xl font-semibold">
              ¿Ya tenes cuenta?
            </p>
            <button
              type="button"
              className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
            >
              Entrar
            </button>
          </div>
        </div>
        <div className="register-img-container hidden lg:flex w-[45%] h-[40rem] rounded-lg">
          <div className="hero-overlay bg-opacity-20 bg-[#151515] rounded-lg"></div>
        </div>
      </motion.div>
    </section>
  );
};

export default Register;
