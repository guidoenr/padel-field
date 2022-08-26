import React from "react";
import { useState } from "react";
import { Link, animateScroll as scroll } from "react-scroll";
import ThemeSwitcher from "./ThemeSwitcher";

// Tomorrow start with Palette Color configurations in DaisyUI...
// Fix some errors, and finish this project. :)

const Header = () => {
  const [nav, setNav] = useState(false);

  const handleMenu = () => setNav(!nav);

  const handleClose = () => setNav(!nav);

  return (
    <header className="w-full h-[10vh] flex shadow-2xl z-50 px-[20px]">
      <div className="container m-auto navbar flex justify-between">
        <div className="navbar-start">
          <a
            className="cursor-pointer font-secondary-font text-4xl text-primary"
            href="#home"
          >
            Pádel-Logo
          </a>
        </div>

        <div className="flex justify-end">
          <nav className="navbar-center hidden md:flex">
            <ul className="menu menu-horizontal p-0 gap-8">
              <li className="text-sm">
                <Link
                  onClick={handleClose}
                  to="features"
                  smooth={true}
                  offset={-70}
                  duration={500}
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent transition ease-in-out p-0"
                >
                  Home
                </Link>
              </li>
              <li className="text-sm">
                <Link
                  onClick={handleClose}
                  to="download"
                  smooth={true}
                  offset={-70}
                  duration={500}
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent transition ease-in-out p-0 text-accent"
                >
                  Reservas
                </Link>
              </li>
              <li className="text-sm">
                <Link
                  onClick={handleClose}
                  to="contact"
                  smooth={true}
                  offset={400}
                  duration={500}
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent transition ease-in-out p-0"
                >
                  Nosotros
                </Link>
              </li>
              <li className="text-sm">
                <a
                  href="#home"
                  className="uppercase font-medium tracking-widest hover:scale-105 hover:bg-transparent transition ease-in-out p-0"
                >
                  Contacto
                </a>
              </li>
              <ThemeSwitcher className="" />
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
                    className="swap-off fill-current w-8 h-8 relative top-[.1rem]"
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 512 512"
                  >
                    <path d="M64,384H448V341.33H64Zm0-106.67H448V234.67H64ZM64,128v42.67H448V128Z" />
                  </svg>

                  <svg
                    className="swap-on fill-current w-8 h-8 relative top-[.1rem]"
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
                  : "md:hidden absolute left-0 top-0 h-screen w-screen bg-very-dark-blue/90"
              }
            >
              <div className="flex flex-col w-[92%] h-[90vh] mx-auto">
                <nav className="hidden">
                  <ul className="menu menu-compact w-[100%] p-2 rounded-box">
                    <li className="text-base border-b border-t rounded-none py-3 mt-8 flex">
                      <Link
                        onClick={handleClose}
                        to="features"
                        smooth={true}
                        offset={-70}
                        duration={500}
                        className="uppercase font-normal self-center text-lg tracking-widest"
                      >
                        Features
                      </Link>
                    </li>
                    <li className="border-b py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="download"
                        smooth={true}
                        offset={-70}
                        duration={500}
                        className="uppercase font-normal self-center text-lg tracking-widest"
                      >
                        Download
                      </Link>
                    </li>
                    <li className="border-b py-3 flex">
                      <Link
                        onClick={handleClose}
                        to="contact"
                        smooth={true}
                        offset={400}
                        duration={500}
                        className="uppercase font-normal self-center text-lg tracking-widest"
                      >
                        Contact
                      </Link>
                    </li>
                    <li className="py-3">
                      <a
                        onClick={handleClose}
                        href="#home"
                        className="btn font-normal px-11 bg-accent border-2 border-bg-color text-bg-color text-lg tracking-widest mt-4"
                      >
                        Login
                      </a>
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
