import React from "react";
import { useEffect } from "react";
import { themeChange } from "theme-change";
import { BiSun } from "react-icons/bi";
import { BsMoonStars } from "react-icons/bs";

const ThemeSwitcher = () => {
  useEffect(() => {
    themeChange(false);
    // 👆 false parameter is required for react project
  }, []);
  return (
    <div className="switcher-container flex gap-4 rounded-lg bg-primary/5 px-3 py-1">
      <button
        type="button"
        data-set-theme="light"
        data-act-class="ACTIVECLASS"
        className="flex items-center justify-center"
      >
        <BiSun className="w-6 h-7 hover:scale-105 transition ease-in-out" />
      </button>
      <button
        type="button"
        data-set-theme="dark"
        data-act-class="ACTIVECLASS"
        className="flex items-center justify-center"
      >
        <BsMoonStars className="w-5 h-6 hover:scale-105 transition ease-in-out" />
      </button>
    </div>
  );
};

export default ThemeSwitcher;
