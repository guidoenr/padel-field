import React, { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import {Link, Navigate} from "react-router-dom";
import {SyntheticEvent} from "react";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false)

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault()
    const response = await fetch("http://localhost:8080/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        username,
        password,
      }),
    });
    if (response.ok){
      setRedirect(true)
    }

  }
  if (redirect) {
    return <Navigate to="/"> </Navigate>
  }

  return (
    <section className="login w-full h-[90vh] flex items-center">
      <div className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container pt-12 lg:pt-0 w-[95%] h-[74%] lg:h-[85%] m-auto p-8 md:flex md:justify-between md:items-center md:max-w-6xl relative">
        <div className="form-container flex flex-col gap-2 md:w-[45%]">
          <div className="md:absolute md:top-8 md:left-8 pb-2">
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
                onChange={e => setUsername(e.target.value)}
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
                onChange={e => setPassword(e.target.value)}
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
        <div className="login-img-container hidden md:flex w-[50%] h-[40rem] rounded-lg">
          <div className="hero-overlay bg-opacity-30 bg-[#151515] rounded-lg"></div>
        </div>
      </div>
    </section>
  );
};

export default Login;
