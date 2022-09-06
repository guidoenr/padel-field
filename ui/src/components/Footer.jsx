import React from "react";
import { AiOutlineTwitter, AiFillInstagram } from "react-icons/ai";
import { GrFacebook } from "react-icons/gr";
import { Link } from "react-router-dom";

const Footer = () => {
  return (
    <footer className="footer h-[13vh] w-full flex justify-center items-center">
      <div className="footer-container container w-full px-1 flex items-center justify-between">
        <div className="linkedin-container lg:w-[10rem] flex gap-3 items-center text-base">
          <Link
            to="/"
            className="cursor-pointer font-secondary-font text-4xl text-[#151515]"
          >
            Pádel-Logo
          </Link>
        </div>
        <ul className="social-links-list flex items-center gap-3 p-3">
          <li className="social-link pr-2 border-r border-r-secondary">
            <a
              href="https://www.instagram.com"
              target="_blank"
              rel="noreferrer"
            >
              <AiFillInstagram className="text-2xl hover:text-secondary/60 text-secondary/80 cursor-pointer text-center hover:translate-y-[-3px] transition ease-out" />
            </a>
          </li>
          <li className="social-link pr-2 border-r border-r-secondary">
            <a href="https://www.facebook.com" target="_blank" rel="noreferrer">
              <GrFacebook className="text-xl hover:text-secondary/60 text-secondary/80 cursor-pointer text-center hover:translate-y-[-3px] transition ease-out" />
            </a>
          </li>
          <li className="social-link pr-2 border-r border-r-secondary">
            <a href="https://www.twitter.com" target="_blank" rel="noreferrer">
              <AiOutlineTwitter className="text-2xl hover:text-secondary/60 text-secondary/80 cursor-pointer text-center hover:translate-y-[-3px] transition ease-out" />
            </a>
          </li>
        </ul>
      </div>
    </footer>
  );
};

export default Footer;
