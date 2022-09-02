import React, { SyntheticEvent, useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Navigate } from "react-router-dom";

const Register = () => {
  const [name, setName] = useState("");
  const [surname, setSurname] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("");
  const [redirect, setRedirect] = useState(false);

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/auth/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        name,
        surname,
        email,
        phone,
        password,
        username,
      }),
    });
    if (response.ok) {
      setRedirect(true);
    }
  };
  if (redirect) {
    return <Navigate to="/login"> </Navigate>;
  }

  /*
@marcos: asi no me olvido v2

me gustaria que la foto del login/register cambie, osea no que sea la misma porque me paso a mi por ejemplo
( q soy bastante crack ) de confundirme las dos pages, pensando que estaba por logearme y me estaba registrando
me gustaria que
- cambies la foto de c/u
- y que cambies de lugar las dos cosas, es decir, del lado derecho la foto para el login y el izquierdo para el register

ej: (ponelo como quieras me chupa un huevo a mi)

|-----------LOGIN-----------|
| Username    |      foto   |
| Password    |             |
----------------------------

|-----------REGISTER---------|
|          |      Username   |
|          |       Password   |
|    foto  |       Email      |
|          |       Nombre     |
|          |       Phone      |
----------------------------

happy coding


@MARCOS: Listorti, habia puesto la misma img solo de ejemplo :)


*/
  return (
    <section className="register w-full h-[90vh] flex items-center">
      <div className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container w-[95%] mt-5 lg:mt-0 lg:w-[90%] lg:h-[90%] mx-auto p-8 lg:flex lg:justify-between lg:items-center lg:max-w-7xl relative">
        <div className="form-container flex flex-col gap-2 lg:w-[50%] lg:order-2">
          <div className="lg:absolute lg:top-8 lg:right-8">
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
          <form className="flex flex-col gap-3 " onSubmit={submit}>
            <div className="lg:flex lg:w-full lg:gap-4">
              <div className="row border-b border-b-primary/30">
                <input
                  type="text"
                  placeholder="Nombre"
                  onChange={(e) => setName(e.target.value)}
                  className="focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
                />
              </div>
              <div className="row border-b border-b-primary/30">
                <input
                  type="text"
                  placeholder="Apellido"
                  onChange={(e) => setSurname(e.target.value)}
                  className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
                />
              </div>
            </div>
            <div className="lg:flex lg:w-full lg:gap-4">
              <div className="row border-b border-b-primary/30">
                <input
                  type="username"
                  placeholder="Nombre de usuario"
                  name="username"
                  id="username"
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
                  onChange={(e) => setPassword(e.target.value)}
                  className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
                />
              </div>
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                type="email"
                placeholder="Email"
                name="email"
                id="email"
                onChange={(e) => setEmail(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                type="text"
                placeholder="Celular"
                name="phone"
                id="phone"
                onChange={(e) => setPhone(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row pt-4">
              <button
                type="submit"
                className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
              >
                Crear cuenta
              </button>
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
      </div>
    </section>
  );
};

export default Register;
