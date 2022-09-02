import React, { SyntheticEvent, useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Navigate } from "react-router-dom";
import { BiErrorAlt } from "react-icons/bi";

const Register = () => {
  const [name, setName] = useState("");
  const [surname, setSurname] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("");
  const [redirect, setRedirect] = useState(false);
  const [focused, setFocused] = useState(false);

  const handleSubmit = async (e: SyntheticEvent) => {
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

  const handleFocus = (e) => {
    setFocused(true);
  };

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

----------------------------------------------------- ready --------------------------------------------------------------


*/
  return (
    <section className="register w-full h-[100vh] lg:h-[90vh] lg:flex lg:items-center">
      <div className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container mt-8 w-[95%] lg:mt-0 lg:w-[90%] lg:h-[90%] mx-auto p-8 lg:flex lg:justify-between lg:items-center lg:max-w-7xl relative">
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
            <div className="flex flex-col gap-2 lg:flex-row lg:w-full lg:gap-4">
              <div className="row form-input flex flex-col gap-1 lg:w-[50%]">
                <label
                  htmlFor="name"
                  className="text-primary/60 font-semibold text-sm"
                >
                  Nombre:
                </label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  onChange={(e) => setName(e.target.value)}
                  className="focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
                  pattern="^[A-Za-z0-9]{3,16}$"
                  onBlur={handleFocus}
                  focused={focused.toString()}
                  required
                />
                <p className="error-msj text-red-700 text-sm">
                  <span>
                    <BiErrorAlt className="inline text-lg" /> Tu nombre debe
                    contener entre 3 y 16 caracteres y no debe contener ningun
                    caracter especial!
                  </span>
                </p>
              </div>
              <div className="row form-input flex flex-col gap-1 lg:w-[50%]">
                <label
                  htmlFor="surname"
                  className="text-primary/60 font-semibold text-sm"
                >
                  Apellido:
                </label>
                <input
                  id="surname"
                  name="surname"
                  type="text"
                  onChange={(e) => setSurname(e.target.value)}
                  className="form-control focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
                  pattern="^[A-Za-z0-9]{3,16}$"
                />
                <p className="error-msj text-red-700 text-sm">
                  <span>
                    <BiErrorAlt className="inline text-lg" /> Tu nombre debe
                    contener entre 3 y 16 caracteres y no debe contener ningun
                    caracter especial!
                  </span>
                </p>
              </div>
            </div>
            <div className="flex flex-col gap-2 lg:flex-row lg:w-full lg:gap-4">
              <div className="row form-input flex flex-col gap-1 lg:w-[50%]">
                <label
                  htmlFor="username"
                  className="text-primary/60 font-semibold text-sm"
                >
                  Nombre de usuario:
                </label>
                <input
                  type="username"
                  id="username"
                  name="username"
                  onChange={(e) => setUsername(e.target.value)}
                  className="form-control focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
                />
              </div>
              <div className="row form-input flex flex-col gap-1 lg:w-[50%]">
                <label
                  htmlFor="password"
                  className="text-primary/60 font-semibold text-sm"
                >
                  Contraseña:
                </label>
                <input
                  type="password"
                  name="password"
                  id="password"
                  onChange={(e) => setPassword(e.target.value)}
                  className="form-control focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
                />
              </div>
            </div>
            <div className="row form-input flex flex-col gap-1">
              <label
                htmlFor="email"
                className="text-primary/60 font-semibold text-sm"
              >
                Email:
              </label>
              <input
                type="email"
                name="email"
                id="email"
                onChange={(e) => setEmail(e.target.value)}
                className="form-control focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
              />
            </div>
            <div className="row form-input flex flex-col gap-1">
              <label
                htmlFor="phone"
                className="text-primary/60 font-semibold text-sm"
              >
                Celular:
              </label>
              <input
                type="text"
                name="phone"
                id="phone"
                onChange={(e) => setPhone(e.target.value)}
                className="form-control focus:outline-accent input h-10 w-full p-1 border-r border-l bg-transparent border-primary/30"
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
