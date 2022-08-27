import React from "react";
import { useState } from "react";
import { Link } from "react-router-dom";
import ThemeSwitcher from "./ThemeSwitcher";

const Header = () => {
  const [nav, setNav] = useState(false);

  const handleMenu = () => setNav(!nav);

  const handleClose = () => setNav(!nav);

  return (
    <header className="w-full h-[10vh] flex shadow-2xl z-50">
      <div className="container m-auto navbar flex justify-between">
        <div className="navbar-start">
          <Link
            to="/"
            className="cursor-pointer font-secondary-font text-4xl text-primary"
          >
            Pádel-Logo
          </Link>
        </div>

        <div className="flex justify-end">
          <nav className="navbar-center hidden md:flex">
            <ul className="menu menu-horizontal p-0 gap-8">
              <li className="text-sm">
                <Link
                  to="/"
                  className="uppercase font-bold tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0 text-accent"
                >
                  Reservas
                </Link>
              </li>
              <li className="text-sm">
                <Link
                  to="/nosotros"
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                >
                  Nosotros
                </Link>
              </li>
              <li className="text-sm">
                <Link
                  to="/contacto"
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                >
                  Contacto
                </Link>
              </li>
              <li className="flex items-center">
                <ThemeSwitcher className="" />
              </li>
              <li className="">
                <Link
                  to="/login"
                  className="border-accent uppercase btn bg-primary/5 px-6 font-medium tracking-widest hover:border-accent/80 hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                >
                  Entrar
                </Link>
              </li>
            </ul>
          </nav>

          <div className="navbar-end">
            <div className="flex items-center justify-end gap-4">
              <div className="theme-container z-[90] md:hidden">
                <ThemeSwitcher />
              </div>

              <div className="hamburger-container md:hidden z-[90] flex items-center">
                <label className="bg-transparent border-none swap swap-rotate">
                  <input type="checkbox" onClick={handleMenu} />
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth={1.5}
                    stroke="currentColor"
                    className="swap-off fill-current w-9 h-9 relative"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M3.75 6.75h16.5M3.75 12h16.5M12 17.25h8.25"
                    />
                  </svg>

                  <svg
                    className="swap-on fill-current w-9 h-9 relative"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 512 512"
                  >
                    <polygon points="400 145.49 366.51 112 256 222.51 145.49 112 112 145.49 222.51 256 112 366.51 145.49 400 256 289.49 366.51 400 400 366.51 289.49 256 400 145.49" />
                  </svg>
                </label>
              </div>
            </div>
            <div
              className={
                !nav
                  ? "hidden"
                  : "md:hidden absolute left-16 top-0 h-screen w-screen bg-base-100 z-[60]"
              }
            >
              <div className="flex flex-col w-[80%] h-[90vh] mx-auto">
                <nav className="flex items-center h-[80vh]">
                  <ul className="menu menu-compact w-[80%] p-2 rounded-box">
                    <li className="text-base border-b border-b-primary/30 py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="/"
                        className="uppercase font-normal self-center text-lg tracking-widest hover:bg-base-100"
                      >
                        Reservas
                      </Link>
                    </li>
                    <li className="border-b border-b-primary/30 py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="/nosotros"
                        className="uppercase font-normal self-center text-lg tracking-widest hover:bg-base-100"
                      >
                        Nosotros
                      </Link>
                    </li>
                    <li className="border-b border-b-primary/30 py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="/contacto"
                        className="uppercase font-normal self-center text-lg tracking-widest hover:bg-base-100"
                      >
                        Contacto
                      </Link>
                    </li>
                    <li className="py-6">
                      <Link
                        onClick={handleClose}
                        to="/login"
                        className="border-accent uppercase btn bg-primary/5 px-6 font-medium tracking-widest hover:border-accent/80 hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                      >
                        Entrar
                      </Link>
                    </li>
                  </ul>
                </nav>
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
