import React, { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Link, Navigate } from "react-router-dom";
import { SyntheticEvent } from "react";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        username,
        password,
      }),
    });
    if (response.ok) {
      setRedirect(true);
    }
  };
  if (redirect) {
    return <Navigate to="/"> </Navigate>;
  }
  /*
  @marcos: asi no me olvido

  el estado de esta page va a depender del http status code de la response,
  (eso olvidate que lo manejo yo que ya me di maña con react)
  lo que te voy a pedir es, una serie de cartelitos/aviso (vos sabras como hacerlo lindo)
  que salten en el login, y digan algo asi:

  [para el login, osea en esta page]
   - usuario no existe
   - email no existe
   - contraseña incorrecta
   - error interno

   [para el register]
    - usuario ya registrado
    - email ya registrado
    - error interno

  entiendo que esto es totalmente dynamic, porque cuando jugue con los <Navigate to> me di cuenta
  que puedo hacer ifs para chequear el estado de algunas variables que van cambiando a medida
  que el user va usando la pag (+1 puntito para react)
  hace que se muestre cada uno en cada field, es decir:
  (e.g: "usuario no existe" que se muestre en el field 'Nombre de usuario', and go on...)

  lo de las cookies, register, login.. toodo anda flamaaaaa (duran 3 horas las cookies)

  happy coding


  @MARCOS: 
  
  Joyaaa. 

  "puedo hacer ifs para chequear el estado de algunas variables que van cambiando a medida
  que el user va usando la pag" ---> Si.

  Puedo mostrar distintos mensajes dependiendo
  de lo que escribe el usuario en el input
      Ejemplo: 
          Si el nombre de usuario ya existe,
          que salte un cartel "nombre de usuario ya registrado".

  Tmb puedo mostrar un layout u otro dependiendo
  de que se cumpla o no alguna condicion.
      Ejemplo:
          If (userLoggedIn === false) {

              Fijar un cartel o un button en la pantalla
              que diga que se tiene que loguear 
              primero para poder sacar turno.

              O tmb que cuando haga click en Reservar, redireccionarlo
              a Login y no a confirmar reserva.
          }
          
  */

  return (
    <section className="login w-full h-[90vh] flex items-center">
      <div className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container pt-12 lg:pt-8 w-[95%] h-[74%] lg:h-[90%] m-auto p-8 lg:flex lg:justify-between lg:items-center lg:max-w-7xl relative">
        <div className="form-container flex flex-col gap-2 lg:w-[50%]">
          <div className="lg:absolute lg:top-8 lg:left-8 pb-2">
            <a
              className="cursor-pointer font-secondary-font text-4xl text-primary"
              href="#home"
            >
              Pádel-Logo
            </a>
          </div>
          <h3 className="form-title text-2xl font-semibold">Iniciar sesion</h3>
          <p className="form-text text-sm text-primary/60">
            Entra a tu cuenta para poder reservar tu turno.
          </p>
          <form className="flex flex-col gap-3" onSubmit={submit}>
            <div className="row border-b border-b-primary/30">
              <input
                type="text"
                placeholder="Nombre de usuario"
                name="username"
                id="username"
                required
                onChange={(e) => setUsername(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                type="password"
                placeholder="Contraseña"
                name="password"
                id="password"
                required
                onChange={(e) => setPassword(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row pt-4">
              <button
                type="submit"
                className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
              >
                Entrar
              </button>
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
      </div>
    </section>
  );
};

export default Login;
