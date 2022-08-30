import React from "react";
import { useState } from "react";
import { NavLink, Link } from "react-router-dom";
import ThemeSwitcher from "./ThemeSwitcher";
import { HiMenuAlt3 } from "react-icons/hi";
import { MdOutlineClose } from "react-icons/md";

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
                <NavLink
                  to="/"
                  className={({ isActive }) =>
                    isActive
                      ? "font-bold text-accent uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                      : "font-medium uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                  }
                >
                  Reservas
                </NavLink>
              </li>
              <li className="text-sm">
                <NavLink
                  to="/misturnos"
                  className={({ isActive }) =>
                    isActive
                      ? "font-bold text-accent uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                      : "font-medium uppercase tracking-widest hover:scale-105 hover:bg-transparent focus:bg-transparent active:bg-transparent active:text-primary transition ease-in-out p-0"
                  }
                >
                  Mis Turnos
                </NavLink>
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

          <div className="navbar-end relative">
            <div className="flex items-center justify-end gap-4">
              <div className="theme-container z-[90] md:hidden">
                <ThemeSwitcher />
              </div>

              <div className="md:hidden mr-4 z-[90]" onClick={handleMenu}>
                {!nav ? (
                  <HiMenuAlt3 className="text-4xl cursor-pointer" />
                ) : (
                  <MdOutlineClose className="text-4xl cursor-pointer" />
                )}
              </div>
            </div>
            <div
              className={
                !nav
                  ? "hidden"
                  : "md:hidden absolute right-5 top-[4.2rem] h-[17rem] bg-base-100 w-[12rem] z-[60] rounded-box border border-primary/30"
              }
            >
              <div className="flex flex-col w-full h-full justify-start relative">
                <nav className="flex items-center w-full h-full">
                  <ul className="menu menu-compact w-full bg-base-100 rounded-box">
                    <li className="border-b border-b-primary/30 py-3 flex">
                      <NavLink
                        onClick={handleClose}
                        to="/"
                        className={({ isActive }) =>
                          isActive
                            ? "text-accent font-bold uppercase self-center text-base tracking-widest hover:bg-base-100"
                            : "uppercase font-semibold self-center text-base tracking-widest hover:bg-base-100"
                        }
                      >
                        Reservas
                      </NavLink>
                    </li>
                    <li className="border-b border-b-primary/30 py-3 flex">
                      <NavLink
                        onClick={handleClose}
                        to="/misturnos"
                        className={({ isActive }) =>
                          isActive
                            ? "text-accent font-bold uppercase self-center text-base tracking-widest hover:bg-base-100"
                            : "uppercase font-semibold self-center text-base tracking-widest hover:bg-base-100"
                        }
                      >
                        Mis Turnos
                      </NavLink>
                    </li>
                    <li className="border-b border-b-primary/30 py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="/contacto"
                        className="uppercase font-semibold self-center text-base tracking-widest hover:bg-base-100"
                      >
                        Contacto
                      </Link>
                    </li>
                    <li className="py-3">
                      <Link
                        onClick={handleClose}
                        to="/profile"
                        className="uppercase font-semibold self-center text-base tracking-widest hover:bg-base-100"
                      >
                        Profile
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
