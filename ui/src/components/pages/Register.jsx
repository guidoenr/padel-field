import React, {SyntheticEvent, useState} from "react";
import { FcGoogle } from "react-icons/fc";


const Register = () => {
  const [name, setName] = useState('')
  const [surname, setSurname] = useState('')
  const [email, setEmail] = useState('')
  const [phone, setPhone] = useState('')
  const [password, setPassword] = useState('')
  const [username, setUsername] = useState('')

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault()
    const response = await fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: {'Content-Type':'application/json'},
      body: JSON.stringify({
        name,
        surname,
        email,
        phone,
        password,
        username
      })
    });
    const bbb = response.body
    console.log(response)
    console.log(bbb)

  }

  return (

    <section className="register w-full h-[80vh] mt-12">
      <div className="border border-primary/30 shadow-lg bg-neutral/60 rounded-lg login-container container w-[95%] m-auto p-8 md:flex md:justify-between md:items-center md:max-w-6xl relative">
        <div className="form-container flex flex-col gap-2 md:w-[45%]">
          <div className="md:absolute md:top-8 md:left-8">
            <a
              className="cursor-pointer font-secondary-font text-4xl text-primary"
              href="#home"
            >
              Pádel-Logo
            </a>
          </div>
          <h3 className="form-title text-2xl font-semibold">
            Crear una cuenta
          </h3>
          <p className="form-text text-sm text-primary/60">
            Crea una cuenta para poder reservar tu turno.
          </p>
          <form action="#" className="flex flex-col gap-3" onSubmit={submit}>
            <div className="row border-b border-b-primary/30">
              <input
                type="text"
                placeholder="Nombre"
                onChange={e => setName(e.target.value)}
                className="focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                  type="text"
                  placeholder="Apellido"
                  onChange={e => setSurname(e.target.value)}
                  className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                  type="username"
                  placeholder="Nombre de usuario"
                  name="username"
                  id="username"
                  onChange={e => setUsername(e.target.value)}
                  className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                type="email"
                placeholder="Email"
                name="email"
                id="email"
                onChange={e => setEmail(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                type="password"
                placeholder="Contraseña"
                name="password"
                id="password"
                onChange={e => setPassword(e.target.value)}
                className="form-control focus:outline-none input input-ghost w-full max-w-xs p-2 bg-transparent active:border-none border-none placeholder:text-primary/60 placeholder:font-semibold"
              />
            </div>
            <div className="row border-b border-b-primary/30">
              <input
                  type="text"
                  placeholder="Celular"
                  name="phone"
                  id="phone"
                  onChange={e => setPhone(e.target.value)}
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
              ¿Ya tenes una cuenta?
            </p>
            <button
              type="button"
              className="btn w-full normal-case bg-accent text-primary border-none hover:bg-accent/70 hover:scale-105 transition ease-in-out"
            >
              Entrar
            </button>
          </div>
        </div>
        <div className="login-img-container hidden md:flex w-[50%] h-[40rem] rounded-lg">
          <div className="hero-overlay bg-opacity-30 bg-[#151515] rounded-lg"></div>
        </div>
      </div>
    </section>
  );
};

export default Register;
