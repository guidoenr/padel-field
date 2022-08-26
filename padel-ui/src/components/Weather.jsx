import React from "react";
import { TiWeatherPartlySunny } from "react-icons/ti";

const Weather = () => {
  return (
    <div className="weather-container absolute top-[-2.5rem] left-0 flex items-center justify-center w-full z-40 select-none">
      <div className="px-3 py-1 flex gap-2 items-center rounded-lg border border-primary/30 hover:border-primary hover:scale-105 transition ease-in-out">
        <TiWeatherPartlySunny className="text-2xl" />
        <p className="weather-text font-bold relative top-[.1rem]">21 °C</p>
      </div>
    </div>
  );
};

export default Weather;
